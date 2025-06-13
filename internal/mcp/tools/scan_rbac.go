package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

// ScanRBACResult represents the complete scan result
type ScanRBACResult struct {
	Findings []permiflow.AccessBinding `json:"findings"`
	Summary  permiflow.Summary         `json:"summary"`
}

// ScanRBACTool creates and returns the scan_rbac tool
func ScanRBACTool() mcp.Tool {
	return mcp.NewTool("scan_rbac",
		mcp.WithDescription("Scan Kubernetes RBAC configurations and identify potential security risks"),
		mcp.WithString("kubeconfig",
			mcp.Description("Path to kubeconfig file (optional, defaults to ~/.kube/config)"),
		),
		mcp.WithString("context",
			mcp.Description("Kubernetes context to use (optional)"),
		),
		mcp.WithArray("namespaces",
			mcp.Description("Specific namespaces to scan (optional, scans all if not specified)"),
		),
		mcp.WithString("format",
			mcp.Description("Output format: 'json' for detailed findings, 'summary' for overview only"),
			mcp.DefaultString("json"),
			mcp.Enum("json", "summary"),
		),
	)
}

// ProcessScanRBACRequest processes the scan_rbac tool request by calling the existing permiflow package
func ProcessScanRBACRequest(request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Processing scan_rbac request with arguments: %+v", request.GetArguments())

	// Extract parameters from request using the proper API
	kubeconfigPath := request.GetString("kubeconfig", "")
	contextName := request.GetString("context", "")
	format := request.GetString("format", "json")
	namespacesFilter := request.GetStringSlice("namespaces", []string{})

	log.Printf("Scan parameters - Kubeconfig: %s, Context: %s, Format: %s, Namespaces: %v",
		kubeconfigPath, contextName, format, namespacesFilter)

	// Use the existing permiflow package to get the Kubernetes client
	var client = permiflow.GetKubeClient("")
	if kubeconfigPath != "" {
		client = permiflow.GetKubeClient(kubeconfigPath)
	}

	if client == nil {
		return mcp.NewToolResultError("Failed to create Kubernetes client"), nil
	}

	log.Printf("Successfully created Kubernetes client")

	// Call the existing permiflow.ScanRBAC function
	bindings, summary := permiflow.ScanRBAC(client)
	log.Printf("Scan completed - Found %d bindings", len(bindings))

	// Filter by namespaces if specified
	if len(namespacesFilter) > 0 {
		bindings = filterBindingsByNamespaces(bindings, namespacesFilter)
		log.Printf("After namespace filtering: %d bindings", len(bindings))
	}

	result := ScanRBACResult{
		Findings: bindings,
		Summary:  summary,
	}

	// Format output based on requested format
	if format == "summary" {
		return mcp.NewToolResultText(formatSummaryText(summary)), nil
	} else {
		// JSON format (default)
		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			return mcp.NewToolResultError(fmt.Sprintf("Error formatting results: %v", err)), nil
		}
		return mcp.NewToolResultText(string(jsonData)), nil
	}
}

// filterBindingsByNamespaces filters the bindings to only include those in the specified namespaces
func filterBindingsByNamespaces(bindings []permiflow.AccessBinding, namespaces []string) []permiflow.AccessBinding {
	if len(namespaces) == 0 {
		return bindings
	}

	// Create a map for faster lookup
	nsMap := make(map[string]bool)
	for _, ns := range namespaces {
		nsMap[ns] = true
	}

	var filtered []permiflow.AccessBinding
	for _, binding := range bindings {
		// Include cluster-scoped bindings or bindings in the specified namespaces
		if binding.Scope == "cluster" || nsMap[binding.Namespace] {
			filtered = append(filtered, binding)
		}
	}

	return filtered
}

// formatSummaryText formats the scan summary as a human-readable text
func formatSummaryText(summary permiflow.Summary) string {
	var sb strings.Builder

	sb.WriteString("RBAC Security Scan Summary\n")
	sb.WriteString("==========================\n\n")

	sb.WriteString(fmt.Sprintf("Cluster Admin Bindings: %d\n", summary.ClusterAdminBindings))
	sb.WriteString(fmt.Sprintf("Wildcard Verbs: %d\n", summary.WildcardVerbs))
	sb.WriteString(fmt.Sprintf("Secrets Access: %d\n", summary.SecretsAccess))
	sb.WriteString(fmt.Sprintf("Privilege Escalation: %d\n", summary.PrivilegeEscalation))
	sb.WriteString(fmt.Sprintf("Exec Access: %d\n", summary.ExecAccess))
	sb.WriteString(fmt.Sprintf("Config Read Secrets: %d\n", summary.ConfigReadSecrets))

	return sb.String()
}
