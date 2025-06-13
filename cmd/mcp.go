package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tutran-se/permiflow/internal/mcp"
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
	Short: "Start the MCP server",
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
		// Create configuration
		cfg := mcp.DefaultConfig()

		// Set values from flags
		cfg.Transport = mcpTransport
		cfg.HTTPPort = mcpHTTPPort
		cfg.Debug = mcpDebug
		cfg.Kubeconfig = mcpKubeconfig
		cfg.Context = mcpContext

		// Load from environment variables (can override flags)
		cfg.LoadFromEnv()

		if cfg.Debug {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Printf("Starting MCP server in debug mode")
			log.Printf("Config: %+v", cfg)
		}

		// Create and start the MCP server
		server, err := mcp.NewServer(cfg)
		if err != nil {
			return err
		}

		// Start the server (this will block)
		return server.Start()
	},
}

func init() {
	rootCmd.AddCommand(mcpCmd)

	// Transport configuration
	mcpCmd.Flags().StringVarP(&mcpTransport, "transport", "t", "stdio", "Transport type (stdio or http)")
	mcpCmd.Flags().IntVarP(&mcpHTTPPort, "http-port", "p", 8080, "HTTP port (only used with http transport)")

	// Debug configuration
	mcpCmd.Flags().BoolVar(&mcpDebug, "debug", false, "Enable debug logging")

	// Kubernetes configuration
	defaultKubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	if kubeconfigEnv := os.Getenv("KUBECONFIG"); kubeconfigEnv != "" {
		defaultKubeconfig = kubeconfigEnv
	}
	mcpCmd.Flags().StringVar(&mcpKubeconfig, "kubeconfig", defaultKubeconfig, "Path to kubeconfig file (default: $HOME/.kube/config or KUBECONFIG env var)")
	mcpCmd.Flags().StringVar(&mcpContext, "context", "", "Kubernetes context to use")
}
