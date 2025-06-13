package tools

import (
	"context"
	"encoding/json"
	"fmt"

	mcp "github.com/mark3labs/mcp-go/mcp"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

// ScanRBACRequest defines the request format for the scan_rbac tool
type ScanRBACRequest struct {
	Namespaces   []string `json:"namespaces,omitempty"`
	OutputFormat string   `json:"output_format,omitempty"`
	Kubeconfig   string   `json:"kubeconfig,omitempty"`
}

// ScanRBACResponse defines the response format for the scan_rbac tool
type ScanRBACResponse struct {
	Report   string        `json:"report,omitempty"`
	Findings []RBACFinding `json:"findings,omitempty"`
	Summary  ScanSummary   `json:"summary,omitempty"`
}

// RBACFinding represents a single RBAC finding
type RBACFinding struct {
	Subject     string   `json:"subject,omitempty"`
	SubjectKind string   `json:"subject_kind,omitempty"`
	Role        string   `json:"role,omitempty"`
	Namespace   string   `json:"namespace,omitempty"`
	Verbs       []string `json:"verbs,omitempty"`
	Resources   []string `json:"resources,omitempty"`
	Scope       string   `json:"scope,omitempty"`
	RiskLevel   string   `json:"risk_level,omitempty"`
	Reason      string   `json:"reason,omitempty"`
}

// ScanSummary represents the summary statistics from the scan
type ScanSummary struct {
	TotalBindings        int `json:"total_bindings"`
	ClusterAdminBindings int `json:"cluster_admin_bindings"`
	WildcardVerbs        int `json:"wildcard_verbs"`
	SecretsAccess        int `json:"secrets_access"`
	PrivilegeEscalation  int `json:"privilege_escalation"`
	ExecAccess           int `json:"exec_access"`
	ConfigReadSecrets    int `json:"config_read_secrets"`
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

	// Use the actual permiflow scanning logic
	client := permiflow.GetKubeClient(req.Kubeconfig)

	// Perform the actual RBAC scan
	bindings, summary := permiflow.ScanRBAC(client)

	// Convert permiflow.AccessBinding to RBACFinding
	findings := make([]RBACFinding, len(bindings))
	for i, binding := range bindings {
		findings[i] = RBACFinding{
			Subject:     binding.Subject,
			SubjectKind: binding.SubjectKind,
			Role:        binding.Role,
			Namespace:   binding.Namespace,
			Verbs:       binding.Verbs,
			Resources:   binding.Resources,
			Scope:       binding.Scope,
			RiskLevel:   binding.RiskLevel,
			Reason:      binding.Reason,
		}
	}

	// Convert permiflow.Summary to ScanSummary
	scanSummary := ScanSummary{
		TotalBindings:        len(bindings),
		ClusterAdminBindings: summary.ClusterAdminBindings,
		WildcardVerbs:        summary.WildcardVerbs,
		SecretsAccess:        summary.SecretsAccess,
		PrivilegeEscalation:  summary.PrivilegeEscalation,
		ExecAccess:           summary.ExecAccess,
		ConfigReadSecrets:    summary.ConfigReadSecrets,
	}

	return &ScanRBACResponse{
		Report:   fmt.Sprintf("rbac-report.%s", req.OutputFormat),
		Findings: findings,
		Summary:  scanSummary,
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
	mcp.WithString("kubeconfig",
		mcp.Description("Path to kubeconfig file (optional, defaults to $HOME/.kube/config)"),
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
