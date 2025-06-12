package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/cors"
	mcpServer "github.com/mark3labs/mcp-go/server"
	"github.com/tutran-se/permiflow/internal/mcp/config"
	"github.com/tutran-se/permiflow/internal/mcp/tools"
)

// Server represents the MCP server
type Server struct {
	config    *config.Config
	server    *mcpServer.MCPServer
	stopChan  chan struct{}
	wg        sync.WaitGroup
	httpSrv   *http.Server
	done      chan struct{}
	closeOnce sync.Once
}

// GetMCPServer returns the underlying MCP server instance
func (s *Server) GetMCPServer() *mcpServer.MCPServer {
	return s.server
}

// NewServer creates a new Server instance
func NewServer(cfg *config.Config) (*Server, error) {
	srv := &Server{
		config:   cfg,
		stopChan: make(chan struct{}),
		done:     make(chan struct{}),
		server:   mcpServer.NewMCPServer("permiflow", "0.1.0"),
	}

	// Register tools
	if err := tools.RegisterTools(srv.server); err != nil {
		return nil, fmt.Errorf("failed to register tools: %w", err)
	}

	return srv, nil
}

// Start starts the MCP server with the configured transport
func (s *Server) Start() error {
	log.Printf("Starting MCP server with %s transport...", s.config.Transport)
	
	// Register all tools
	if err := tools.RegisterTools(s.server); err != nil {
		return fmt.Errorf("failed to register tools: %w", err)
	}
	log.Println("Successfully registered all tools")

	// Start the appropriate server based on transport
	var err error

	switch s.config.Transport {
	case "http":
		log.Printf("Initializing HTTP server on port %d", s.config.HTTPPort)
		_, err = s.startHTTPServer()
	case "stdio":
		log.Println("Initializing stdio server")
		_, err = s.startStdioServer()
	default:
		err = fmt.Errorf("unsupported transport: %s", s.config.Transport)
	}

	if err != nil {
		return fmt.Errorf("failed to start %s server: %w", s.config.Transport, err)
	}

	log.Printf("MCP server started successfully with %s transport", s.config.Transport)
	return nil
}

// Stop gracefully shuts down the server
func (s *Server) Stop() error {
	log.Println("Initiating graceful shutdown of MCP server...")

	s.closeOnce.Do(func() {
		// Close the stop channel to signal shutdown
		close(s.stopChan)

		// Stop HTTP server if running
		if s.httpSrv != nil {
			log.Println("Shutting down HTTP server...")
			if err := s.httpSrv.Shutdown(context.Background()); err != nil {
				log.Printf("HTTP server shutdown error: %v", err)
			} else {
				log.Println("HTTP server stopped")
			}
		}

		// Close the done channel
		close(s.done)
	})

	// Wait for all goroutines to finish
	log.Println("Waiting for all server operations to complete...")
	
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	// Wait with timeout
	select {
	case <-done:
		log.Println("All server operations completed")
	case <-time.After(30 * time.Second):
		log.Println("Warning: Server shutdown timed out, some operations may not have completed")
	}

	log.Println("MCP server shutdown complete")
	return nil
}

// healthHandler handles health check requests
type healthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}

// startHTTPServer starts the HTTP server
func (s *Server) startHTTPServer() (chan struct{}, error) {
	addr := fmt.Sprintf(":%d", s.config.HTTPPort)
	
	// Create a new router
	mux := http.NewServeMux()
	
	// Health check endpoints
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(healthResponse{
			Status:  "ok",
			Version: "0.1.0",
		})
	})
	
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(healthResponse{
			Status:  "ready",
			Version: "0.1.0",
		})
	})
	
	// MCP server handler with CORS
	mcpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Create a new HTTP server with the MCP server
		httpSrv := mcpServer.NewStreamableHTTPServer(s.server)
		httpSrv.ServeHTTP(w, r)
	})
	
	// Register MCP handler for the root path
	mux.Handle("/", mcpHandler)
	
	// Create a CORS handler
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(mux)

	s.httpSrv = &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	done := make(chan struct{})

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		log.Printf("Starting HTTP server on %s", addr)
		if err := s.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	return done, nil
}

// startStdioServer starts the MCP server in stdio mode
func (s *Server) startStdioServer() (chan struct{}, error) {
	done := make(chan struct{})
	
	// Create a new stdio server
	stdioSrv := mcpServer.NewStdioServer(s.server)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		log.Println("Starting MCP server in stdio mode")
		if err := stdioSrv.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
			log.Printf("Error in stdio server: %v", err)
		}
		close(done)
	}()

	return done, nil
}
