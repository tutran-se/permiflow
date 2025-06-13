package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

// Config holds server configuration
type Config struct {
	Transport  string
	HTTPPort   int
	Debug      bool
	Kubeconfig string
	Context    string
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		Transport:  "stdio",
		HTTPPort:   8080,
		Debug:      false,
		Kubeconfig: "",
		Context:    "",
	}
}

// LoadFromEnv loads configuration from environment variables
func (c *Config) LoadFromEnv() {
	if transport := os.Getenv("MCP_TRANSPORT"); transport != "" {
		c.Transport = transport
	}
	if debug := os.Getenv("MCP_DEBUG"); debug == "true" {
		c.Debug = true
	}
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		c.Kubeconfig = kubeconfig
	}
	if context := os.Getenv("MCP_KUBE_CONTEXT"); context != "" {
		c.Context = context
	}
}

// Server represents the MCP server instance
type Server struct {
	config *Config
	mcp    *server.MCPServer
}

// ScanRBACResult represents the result of RBAC scanning
type ScanRBACResult struct {
	Findings []permiflow.AccessBinding `json:"findings"`
	Summary  permiflow.Summary         `json:"summary"`
}

// NewServer creates a new MCP server instance
func NewServer(cfg *Config) (*Server, error) {
	switch cfg.Transport {
	case "http":
		return NewStreamableHTTPServer(cfg)
	case "stdio":
		return NewSTDIOServer(cfg)
	default:
		return nil, fmt.Errorf("unsupported transport: %s", cfg.Transport)
	}
}

// NewStreamableHTTPServer creates a new server using the StreamableHTTP approach
func NewStreamableHTTPServer(cfg *Config) (*Server, error) {
	// Create the MCP server with capabilities
	mcpServer := server.NewMCPServer(
		"Permiflow MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Add the scan_rbac tool
	mcpServer.AddTool(
		mcp.NewTool("scan_rbac",
			mcp.WithDescription("Scan Kubernetes RBAC configurations and identify potential security risks"),
			mcp.WithString("kubeconfig", mcp.Description("Path to kubeconfig file (optional, defaults to ~/.kube/config)")),
			mcp.WithString("context", mcp.Description("Kubernetes context to use (optional)")),
			mcp.WithString("format", mcp.Description("Output format: 'json' for detailed findings, 'summary' for overview only")),
		),
		handleScanRBAC,
	)

	return &Server{
		config: cfg,
		mcp:    mcpServer,
	}, nil
}

// NewSTDIOServer creates a new server using STDIO transport
func NewSTDIOServer(cfg *Config) (*Server, error) {
	// Create the MCP server with capabilities
	mcpServer := server.NewMCPServer(
		"Permiflow MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Add the scan_rbac tool
	mcpServer.AddTool(
		mcp.NewTool("scan_rbac",
			mcp.WithDescription("Scan Kubernetes RBAC configurations and identify potential security risks"),
			mcp.WithString("kubeconfig", mcp.Description("Path to kubeconfig file (optional, defaults to ~/.kube/config)")),
			mcp.WithString("context", mcp.Description("Kubernetes context to use (optional)")),
			mcp.WithString("format", mcp.Description("Output format: 'json' for detailed findings, 'summary' for overview only")),
		),
		handleScanRBAC,
	)

	return &Server{
		config: cfg,
		mcp:    mcpServer,
	}, nil
}

// handleScanRBAC is the tool handler for both transports
func handleScanRBAC(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract parameters safely
	var kubeconfigPath, format string
	var namespacesFilter []string

	// Parse arguments - handle the interface{} type properly
	if args, ok := req.Params.Arguments.(map[string]interface{}); ok {
		if val, exists := args["kubeconfig"]; exists && val != nil {
			if str, ok := val.(string); ok {
				kubeconfigPath = str
			}
		}

		if val, exists := args["format"]; exists && val != nil {
			if str, ok := val.(string); ok {
				format = str
			}
		}

		if val, exists := args["namespaces"]; exists && val != nil {
			if arr, ok := val.([]interface{}); ok {
				for _, v := range arr {
					if str, ok := v.(string); ok {
						namespacesFilter = append(namespacesFilter, str)
					}
				}
			}
		}
	}

	// Set default format
	if format == "" {
		format = "json"
	}

	// Use the existing permiflow package to get the Kubernetes client
	var client = permiflow.GetKubeClient("")
	if kubeconfigPath != "" {
		client = permiflow.GetKubeClient(kubeconfigPath)
	}

	if client == nil {
		return nil, fmt.Errorf("failed to create Kubernetes client")
	}

	// Call the existing permiflow.ScanRBAC function
	bindings, summary := permiflow.ScanRBAC(client)

	// Filter by namespaces if specified
	if len(namespacesFilter) > 0 {
		bindings = filterBindingsByNamespaces(bindings, namespacesFilter)
	}

	result := ScanRBACResult{
		Findings: bindings,
		Summary:  summary,
	}

	if format == "summary" {
		summaryText := formatSummaryText(summary)
		return mcp.NewToolResultText(summaryText), nil
	}

	// Return JSON result as text (default)
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result to JSON: %w", err)
	}
	return mcp.NewToolResultText(string(jsonData)), nil
}

// Start starts the appropriate server based on transport
func (s *Server) Start() error {
	switch s.config.Transport {
	case "http":
		return s.startHTTP()
	case "stdio":
		return s.startSTDIO()
	default:
		return fmt.Errorf("unsupported transport: %s", s.config.Transport)
	}
}

// startHTTP starts the StreamableHTTP server
func (s *Server) startHTTP() error {
	log.Printf("Starting MCP server with StreamableHTTP transport on port %d", s.config.HTTPPort)

	// Create StreamableHTTP server
	httpServer := server.NewStreamableHTTPServer(s.mcp)
	addr := fmt.Sprintf(":%d", s.config.HTTPPort)

	log.Printf("MCP server listening on %s", addr)
	log.Printf("Available endpoints:")
	log.Printf("  POST http://localhost%s/mcp - MCP JSON-RPC endpoint", addr)
	log.Printf("  GET  http://localhost%s/mcp - MCP streaming endpoint", addr)

	// Start the server (this will block)
	if err := httpServer.Start(addr); err != nil {
		return fmt.Errorf("failed to start StreamableHTTP server: %w", err)
	}

	return nil
}

// startSTDIO starts the STDIO server
func (s *Server) startSTDIO() error {
	log.Printf("Starting MCP server with STDIO transport")
	log.Printf("Server is ready to receive JSON-RPC messages on stdin")

	// Start the server (this will block)
	if err := server.ServeStdio(s.mcp); err != nil {
		return fmt.Errorf("failed to start STDIO server: %w", err)
	}

	return nil
}

// Helper function to filter bindings by namespaces
func filterBindingsByNamespaces(bindings []permiflow.AccessBinding, namespaces []string) []permiflow.AccessBinding {
	if len(namespaces) == 0 {
		return bindings
	}

	nsSet := make(map[string]bool)
	for _, ns := range namespaces {
		nsSet[ns] = true
	}

	var filtered []permiflow.AccessBinding
	for _, binding := range bindings {
		if nsSet[binding.Namespace] {
			filtered = append(filtered, binding)
		}
	}

	return filtered
}

// Helper function to format summary as text
func formatSummaryText(summary permiflow.Summary) string {
	var sb strings.Builder
	sb.WriteString("=== RBAC Security Scan Summary ===\n\n")
	sb.WriteString(fmt.Sprintf("Cluster admin bindings: %d\n", summary.ClusterAdminBindings))
	sb.WriteString(fmt.Sprintf("Wildcard verbs: %d\n", summary.WildcardVerbs))
	sb.WriteString(fmt.Sprintf("Secrets access: %d\n", summary.SecretsAccess))
	sb.WriteString(fmt.Sprintf("Privilege escalation risks: %d\n", summary.PrivilegeEscalation))
	sb.WriteString(fmt.Sprintf("Exec access: %d\n", summary.ExecAccess))
	sb.WriteString(fmt.Sprintf("ConfigMap/secrets read access: %d\n", summary.ConfigReadSecrets))

	return sb.String()
}
