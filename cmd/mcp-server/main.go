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

func main() {
	// Create root command
	rootCmd := &cobra.Command{
		Use:   "mcp-server",
		Short: "MCP server for Permiflow",
		Long:  `MCP server that exposes Permiflow's RBAC scanning capabilities through the Model Context Protocol`,
	}

	// Initialize config with default values
	cfg := config.DefaultConfig()

	// Add flags
	rootCmd.Flags().StringVarP(&cfg.Transport, "transport", "t", cfg.Transport, "Transport type (stdio or http)")
	rootCmd.Flags().IntVarP(&cfg.HTTPPort, "http-port", "p", cfg.HTTPPort, "HTTP port (only used with http transport)")
	rootCmd.Flags().BoolVar(&cfg.Debug, "debug", cfg.Debug, "Enable debug logging")
	rootCmd.Flags().StringVar(&cfg.Kubeconfig, "kubeconfig", cfg.Kubeconfig, "Path to kubeconfig file")
	rootCmd.Flags().StringVar(&cfg.Context, "context", cfg.Context, "Kubernetes context to use")

	// Set the RunE function
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runServer(cmd, cfg)
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func runServer(cmd *cobra.Command, cfg *config.Config) error {
	// Load config from environment variables (overrides default values)
	cfg.LoadFromEnv()

	// Override with command line flags if provided
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

	// Create server
	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
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
		log.Printf("Server error: %v", err)
		return err
	case sig := <-sigCh:
		log.Printf("Received signal %v, shutting down...", sig)
	case <-ctx.Done():
		log.Println("Context cancelled, shutting down...")
	}

	// Perform graceful shutdown
	log.Println("Initiating graceful shutdown...")
	if err := srv.Stop(); err != nil {
		log.Printf("Error during server shutdown: %v", err)
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}
