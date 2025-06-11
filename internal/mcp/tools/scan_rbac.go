package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
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

// ScanRBACTool implements the RBAC scanning tool
var ScanRBACTool = mcp.NewTool("scan_rbac",
	mcp.WithDescription("Scan Kubernetes RBAC rules and generate a report"),
	mcp.WithString("output_format",
		mcp.Default("json"),
		mcp.Enum("json", "markdown", "csv"),
		mcp.Description("Output format for the report"),
	),
	mcp.WithArray("namespaces",
		mcp.Description("List of namespaces to scan (empty for all)"),
		mcp.Items("string"),
	),
)

func scanRBAC(ctx context.Context, data json.RawMessage) (interface{}, error) {
	var req ScanRBACRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// TODO: Implement actual RBAC scanning using the existing Permiflow logic
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
