package tools

import (
	mcp "github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all available tools with the MCP server
func RegisterTools(server *mcpServer.MCPServer) error {
	// Register RBAC scanning tool with its handler
	server.AddTool(ScanRBACTool, Handler)
	return nil
}

// GetTool returns a tool by name (currently not used but kept for future use)
func GetTool(name string) *mcp.Tool {
	switch name {
	case "scan_rbac":
		return &ScanRBACTool
	default:
		return nil
	}
}
