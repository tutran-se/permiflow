package main

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
	cfg *config.Config
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mcp-server",
		Short: "MCP server for Permiflow",
		Long:  `MCP server that exposes Permiflow's RBAC scanning capabilities through the Model Context Protocol`,
		RunE:  runServer,
	}

	// Add flags
	rootCmd.PersistentFlags().StringVarP(&cfg.Transport, "transport", "t", "stdio", "Transport type (stdio or http)")
	rootCmd.PersistentFlags().IntVarP(&cfg.HTTPPort, "http-port", "p", 8080, "HTTP port (only used with http transport)")
	rootCmd.PersistentFlags().BoolVar(&cfg.Debug, "debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().StringVar(&cfg.Kubeconfig, "kubeconfig", "", "Path to kubeconfig file")
	rootCmd.PersistentFlags().StringVar(&cfg.Context, "context", "", "Kubernetes context to use")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func runServer(cmd *cobra.Command, args []string) error {
	// Load config from environment variables
	cfg = config.DefaultConfig()
	cfg.LoadFromEnv()

	// Override with command line flags
	if cmd.Flags().Changed("transport") {
		transport, _ := cmd.Flags().GetString("transport")
		cfg.Transport = transport
	}
	if cmd.Flags().Changed("http-port") {
		port, _ := cmd.Flags().GetInt("http-port")
		cfg.HTTPPort = port
	}
	if cmd.Flags().Changed("debug") {
		debug, _ := cmd.Flags().GetBool("debug")
		cfg.Debug = debug
	}
	if cmd.Flags().Changed("kubeconfig") {
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		cfg.Kubeconfig = kubeconfig
	}
	if cmd.Flags().Changed("context") {
		ctx, _ := cmd.Flags().GetString("context")
		cfg.Context = ctx
	}

	// Create and start server
	srv := server.NewServer(cfg)

	// Handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start server in a goroutine
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	
	select {
	case <-sigCh:
		log.Println("Shutting down server...")
		srv.Stop()
	case <-ctx.Done():
	}

	return nil
}
