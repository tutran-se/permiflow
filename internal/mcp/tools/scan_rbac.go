package tools

import (
	"context"
	"encoding/json"
	"fmt"

	mcp "github.com/mark3labs/mcp-go/mcp"
)

// ScanRBACRequest defines the request format for the scan_rbac tool
type ScanRBACRequest struct {
	Namespaces  []string `json:"namespaces,omitempty"`
	OutputFormat string   `json:"output_format,omitempty"`
}

// ScanRBACResponse defines the response format for the scan_rbac tool
type ScanRBACResponse struct {
	Report  string        `json:"report,omitempty"`
	Findings []RBACFinding `json:"findings,omitempty"`
}

// RBACFinding represents a single RBAC finding
type RBACFinding struct {
	Namespace string   `json:"namespace,omitempty"`
	Resource  string   `json:"resource,omitempty"`
	Name      string   `json:"name,omitempty"`
	Rules     []Rule   `json:"rules,omitempty"`
	RiskLevel string   `json:"risk_level,omitempty"`
}

// Rule represents an RBAC rule
type Rule struct {
	Verbs     []string `json:"verbs,omitempty"`
	Resources []string `json:"resources,omitempty"`
	APIGroups []string `json:"api_groups,omitempty"`
}

// scanRBAC handles the RBAC scanning logic
func scanRBAC(ctx context.Context, input json.RawMessage) (*ScanRBACResponse, error) {
	var req ScanRBACRequest
	if err := json.Unmarshal(input, &req); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// TODO: Replace with actual RBAC scanning logic from Permiflow
	// For now, return a mock response
	return &ScanRBACResponse{
		Report: fmt.Sprintf("rbac-report.%s", req.OutputFormat),
		Findings: []RBACFinding{
			{
				Namespace: "kube-system",
				Resource:  "clusterrole",
				Name:      "admin",
				Rules: []Rule{
					{
						Verbs:     []string{"*"},
						Resources: []string{"*"},
						APIGroups: []string{"*"},
					},
				},
				RiskLevel: "high",
			},
		},
	}, nil
}

// ScanRBACTool is the MCP tool for scanning Kubernetes RBAC rules
var ScanRBACTool = mcp.NewTool("scan_rbac",
	mcp.WithDescription("Scan Kubernetes RBAC rules and generate a report"),
	mcp.WithString("output_format",
		mcp.Description("Output format for the report"),
		mcp.DefaultString("json"),
		mcp.Enum("json", "markdown", "csv"),
	),
	mcp.WithArray("namespaces",
		mcp.Description("List of namespaces to scan (empty for all)"),
		mcp.Items("string"),
	),
)

// Handler handles the scan_rbac tool requests
func Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Get the raw arguments as JSON
	rawArgs := req.GetRawArguments()
	argsJSON, err := json.Marshal(rawArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request arguments: %w", err)
	}

	// Call the scanRBAC function with the raw input
	resp, err := scanRBAC(ctx, argsJSON)
	if err != nil {
		return nil, fmt.Errorf("RBAC scan failed: %w", err)
	}

	// Convert response to JSON
	resultJSON, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}

	// Return the result in the expected format
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent(string(resultJSON)),
		},
	}, nil
}


