package tools

import "github.com/mark3labs/mcp-go/mcp"

// Registry holds all available MCP tools
var Registry = []mcp.Tool{
	ScanRBACTool,
	// Add other tools here
}

// RegisterTools registers all tools with the MCP server
func RegisterTools(server *mcp.Server) error {
	for _, tool := range Registry {
		server.RegisterTool(tool)
	}
	return nil
}
