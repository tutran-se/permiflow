package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/tutran-se/permiflow/internal/mcp/config"
	"github.com/tutran-se/permiflow/internal/mcp/server"
)

var (
	mcpTransport  string
	mcpHTTPPort   int
	mcpDebug      bool
	mcpKubeconfig string
	mcpContext    string
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the MCP (Model Context Protocol) server",
	Long: `Start the MCP server that exposes Permiflow's RBAC scanning capabilities 
through the Model Context Protocol. Supports both HTTP and STDIO transports.`,
	Example: `
	# Start MCP server with STDIO transport (default)
	permiflow mcp

	# Start MCP server with HTTP transport
	permiflow mcp --transport http --http-port 8080

	# Start with debug logging
	permiflow mcp --debug

	# Use specific kubeconfig
	permiflow mcp --kubeconfig ~/.kube/config --context my-cluster
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMCPServer(cmd)
	},
}

func runMCPServer(cmd *cobra.Command) error {
	// Initialize config with default values
	cfg := config.DefaultConfig()

	// Load config from environment variables (overrides default values)
	cfg.LoadFromEnv()

	// Override with command line flags if provided
	if cmd.Flags().Changed("transport") {
		cfg.Transport = mcpTransport
	}
	if cmd.Flags().Changed("http-port") {
		cfg.HTTPPort = mcpHTTPPort
	}
	if cmd.Flags().Changed("debug") {
		cfg.Debug = mcpDebug
	}
	if cmd.Flags().Changed("kubeconfig") {
		cfg.Kubeconfig = mcpKubeconfig
	}
	if cmd.Flags().Changed("context") {
		cfg.Context = mcpContext
	}

	// Create server
	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}

	// Setup signal handling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel to handle server errors
	errChan := make(chan error, 1)

	// Start the server
	go func() {
		log.Printf("Starting MCP server with %s transport...", cfg.Transport)
		if err := srv.Start(); err != nil {
			errChan <- err
		}
	}()

	// Wait for interrupt signal or server error
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Printf("MCP server error: %v", err)
		return err
	case sig := <-sigCh:
		log.Printf("Received signal %v, shutting down MCP server...", sig)
	case <-ctx.Done():
		log.Println("Context cancelled, shutting down MCP server...")
	}

	// Perform graceful shutdown
	log.Println("Initiating graceful shutdown...")
	if err := srv.Stop(); err != nil {
		log.Printf("Error during MCP server shutdown: %v", err)
		return err
	}

	log.Println("MCP server stopped gracefully")
	return nil
}

func init() {
	rootCmd.AddCommand(mcpCmd)

	mcpCmd.Flags().StringVarP(&mcpTransport, "transport", "t", "stdio", "Transport type (stdio or http)")
	mcpCmd.Flags().IntVarP(&mcpHTTPPort, "http-port", "p", 8080, "HTTP port (only used with http transport)")
	mcpCmd.Flags().BoolVar(&mcpDebug, "debug", false, "Enable debug logging")
	mcpCmd.Flags().StringVar(&mcpKubeconfig, "kubeconfig", "", "Path to kubeconfig file (default: $HOME/.kube/config or KUBECONFIG env var)")
	mcpCmd.Flags().StringVar(&mcpContext, "context", "", "Kubernetes context to use")
}
