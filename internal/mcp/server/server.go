package server

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/tutran-se/permiflow/internal/mcp/config"
	"github.com/tutran-se/permiflow/internal/mcp/tools"
)

// MCPServer represents the MCP server
type MCPServer struct {
	config *config.Config
	mcp    *server.MCPServer
}

// NewServer creates a new MCP server instance
func NewServer(cfg *config.Config) (*MCPServer, error) {
	// Create the MCP server with server information
	mcpServer := server.NewMCPServer(
		"Permiflow MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Register the scan_rbac tool
	scanTool := tools.ScanRBACTool()
	mcpServer.AddTool(scanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.ProcessScanRBACRequest(request)
	})

	s := &MCPServer{
		config: cfg,
		mcp:    mcpServer,
	}

	return s, nil
}

// Start starts the MCP server
func (s *MCPServer) Start() error {
	if s.config.Debug {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Printf("Starting MCP server in debug mode")
		log.Printf("Config: %+v", s.config)
	}

	switch s.config.Transport {
	case "stdio":
		return s.startStdio()
	case "http":
		return s.startHTTP()
	default:
		return fmt.Errorf("unsupported transport: %s", s.config.Transport)
	}
}

// Stop stops the MCP server
func (s *MCPServer) Stop() error {
	log.Printf("Stopping MCP server...")
	// MCP server doesn't have explicit stop method, just return nil
	return nil
}

// startStdio starts the server with stdio transport
func (s *MCPServer) startStdio() error {
	log.Printf("Starting MCP server with stdio transport")
	return server.ServeStdio(s.mcp)
}

// startHTTP starts the server with HTTP transport
func (s *MCPServer) startHTTP() error {
	log.Printf("Starting MCP server with HTTP transport on port %d", s.config.HTTPPort)

	// For now, HTTP transport is not fully supported in the current mcp-go version
	// The main transport mechanism is stdio
	log.Printf("HTTP transport not fully implemented in current mcp-go version")
	log.Printf("Falling back to stdio transport for now")
	return s.startStdio()
}
