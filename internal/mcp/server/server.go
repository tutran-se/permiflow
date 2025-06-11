package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	mcp "github.com/mark3labs/mcp-go"
	mcpServer "github.com/mark3labs/mcp-go/server"
	"github.com/tuannvm/permiflow/internal/mcp/config"
	"github.com/tuannvm/permiflow/internal/mcp/tools"
)

// Server represents the MCP server
type Server struct {
	config   *config.Config
	server   *mcp.Server
	stopChan chan struct{}
	wg       sync.WaitGroup
}

// NewServer creates a new Server instance
func NewServer(cfg *config.Config) (*Server, error) {
	srv := &Server{
		config:   cfg,
		stopChan: make(chan struct{}),
	}

	// Create MCP server
	srv.server = mcp.NewServer()

	// Register tools
	if err := tools.RegisterTools(srv.server); err != nil {
		return nil, fmt.Errorf("failed to register tools: %w", err)
	}

	return srv, nil
}

// Start starts the MCP server
func (s *Server) Start() error {
	switch s.config.Transport {
	case "http":
		return s.startHTTPServer()
	case "stdio":
		fallthrough
	default:
		return s.startStdioServer()
	}
}

// Stop gracefully stops the MCP server
func (s *Server) Stop() {
	close(s.stopChan)
	s.wg.Wait()
}

// startStdioServer starts the MCP server with STDIO transport
func (s *Server) startStdioServer() error {
	log.Println("Starting MCP server with STDIO transport")
	return s.mcpSrv.Serve()
}

// startHTTPServer starts the MCP server with HTTP transport
func (s *Server) startHTTPServer() error {
	log.Printf("Starting MCP server with HTTP transport on :%d\n", s.config.HTTPPort)
	
	handler := s.mcpSrv.HTTPHandler()
	
	// Add CORS headers
	corsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.HTTPPort),
		Handler: corsHandler,
	}

	// Handle graceful shutdown
	go func() {
		<-s.stopChan
		server.Shutdown(context.Background())
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v\n", err)
		}
	}()

	return nil
}
