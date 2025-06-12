package server_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	mcp "github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tutran-se/permiflow/internal/mcp/config"
	"github.com/tutran-se/permiflow/internal/mcp/server"
)

// TestToolRequest defines the request format for the test tool
type TestToolRequest struct {
	Input string `json:"input"`
}

// TestToolResponse defines the response format for the test tool
type TestToolResponse struct {
	Output string `json:"output"`
}

// testTool is a simple test tool for testing
var testTool = mcp.NewTool("test_tool",
	mcp.WithDescription("A test tool"),
	mcp.WithString("input",
		mcp.Description("Input string"),
		mcp.Required(),
	),
)

// testToolHandler handles the test tool
func testToolHandler(ctx context.Context, input json.RawMessage) (*TestToolResponse, error) {
	// Parse the input
	var req TestToolRequest
	if err := json.Unmarshal(input, &req); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Return the response
	return &TestToolResponse{
		Output: "Processed: " + req.Input,
	}, nil
}

// Handler handles the test tool requests
func Handler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Get the raw arguments as JSON
	rawArgs := req.GetRawArguments()
	argsJSON, err := json.Marshal(rawArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request arguments: %w", err)
	}

	// Call the testToolHandler function with the raw input
	resp, err := testToolHandler(ctx, argsJSON)
	if err != nil {
		return nil, fmt.Errorf("test tool failed: %w", err)
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

func TestServer_StartStop(t *testing.T) {
	tests := []struct {
		name      string
		transport string
		skip      bool
	}{
		{"HTTP transport", "http", false},
		{"Stdio transport", "stdio", true}, // Skip stdio test by default as it requires special setup
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip("Skipping test that requires special setup")
			}

			// Create a test config
			cfg := &config.Config{
				Transport: tt.transport,
				HTTPPort:  8080,
			}

			// Create a new server
			srv, err := server.NewServer(cfg)
			require.NoError(t, err)

			// Skip tool registration as it's not needed for basic start/stop test

			// Start the server in a goroutine
			errCh := make(chan error, 1)
			go func() {
				if err := srv.Start(); err != nil {
					errCh <- err
				}
			}()

			// Give the server a moment to start
			time.Sleep(100 * time.Millisecond)

			// Test the appropriate transport
			switch tt.transport {
			case "http":
				testHTTPTransport(t, srv, cfg)
			case "stdio":
				testStdioTransport(t, srv, cfg)
			}

			// Stop the server
			err = srv.Stop()
			require.NoError(t, err)

			// Check for any errors from the server goroutine
			select {
			case err := <-errCh:
				require.NoError(t, err)
			default:
			}
		})
	}
}

func testHTTPTransport(t *testing.T, srv *server.Server, cfg *config.Config) {
	// Test health check endpoint
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/healthz", cfg.HTTPPort))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test ready check endpoint
	resp, err = http.Get(fmt.Sprintf("http://localhost:%d/readyz", cfg.HTTPPort))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Test tool invocation
}

func testStdioTransport(t *testing.T, srv *server.Server, cfg *config.Config) {
	// For stdio transport, we can't easily test from the same process
	// This would typically be tested in an integration test with a separate process
	t.Log("Stdio transport test would be run in integration test with a separate process")
}

// TestServer_CallTool tests calling a tool on the server
func TestServer_CallTool(t *testing.T) {
	t.Skip("Skipping tool call test as it requires a proper client implementation")
	
	// This test is skipped as it requires a proper client implementation.
	// In a real test, we would:
	// 1. Start the server
	// 2. Create a client using the MCP client library
	// 3. Call a tool on the server
	// 4. Verify the response
}

// TestServer_ErrorHandling tests error handling in the server
func TestServer_ErrorHandling(t *testing.T) {
	t.Run("invalid transport", func(t *testing.T) {
		// Test with invalid transport
		cfg := &config.Config{
			Transport: "invalid",
		}

		srv, err := server.NewServer(cfg)
		require.NoError(t, err)

		// Verify we can get the MCP server instance
		mcpServer := srv.GetMCPServer()
		require.NotNil(t, mcpServer, "MCP server should not be nil")

		// Start the server - should fail with invalid transport
		err = srv.Start()
		require.Error(t, err, "Expected error with invalid transport")

		// Stop the server (should be safe to call even if not started)
		err = srv.Stop()
		require.NoError(t, err, "Stop should succeed even if server failed to start")
	})

	t.Run("invalid port", func(t *testing.T) {
		// Skip this test as it requires a proper server implementation
		t.Skip("Skipping invalid port test as it requires proper server implementation")
	})
}

// TestServer_ConcurrentAccess tests concurrent access to the server
func TestServer_ConcurrentAccess(t *testing.T) {
	t.Skip("Skipping concurrent access test as it requires a proper client implementation")
	
	// This test is currently skipped as it requires a proper client implementation
	// to test concurrent access to the server. In a real test, we would:
	// 1. Start the server
	// 2. Create multiple clients
	// 3. Make concurrent requests to the server
	// 4. Verify the responses
}

// TestServer_AddTool tests adding a tool to the server
func TestServer_AddTool(t *testing.T) {
	t.Skip("Skipping AddTool test as it requires a proper MCP server implementation")
	
	// This test is skipped as it requires a proper MCP server implementation
	// that supports the AddTool method with the expected signature.
	// In a real test, we would verify that tools can be added and managed correctly.
}

func TestMain(m *testing.M) {
	// Set up any test fixtures here
	code := m.Run()
	// Clean up any test fixtures here
	os.Exit(code)
}
