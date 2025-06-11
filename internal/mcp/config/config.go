package config

import (
	"os"
	"strconv"
)

// Config holds the MCP server configuration
type Config struct {
	// Transport configuration
	Transport string `json:"transport"` // "stdio" or "http"
	HTTPPort  int    `json:"http_port"`

	// Logging
	Debug bool `json:"debug"`

	// Kubernetes configuration
	Kubeconfig string `json:"kubeconfig"`
	Context    string `json:"context"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Transport:  "stdio",
		HTTPPort:   8080,
		Debug:      false,
		Kubeconfig: "",
		Context:    "",
	}
}

// LoadFromEnv loads configuration from environment variables
func (c *Config) LoadFromEnv() {
	if v := os.Getenv("MCP_TRANSPORT"); v != "" {
		c.Transport = v
	}

	if v := os.Getenv("MCP_HTTP_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			c.HTTPPort = port
		}
	}

	if v := os.Getenv("MCP_DEBUG"); v != "" {
		c.Debug = v == "true" || v == "1"
	}

	if v := os.Getenv("KUBECONFIG"); v != "" {
		c.Kubeconfig = v
	}

	if v := os.Getenv("KUBE_CONTEXT"); v != "" {
		c.Context = v
	}
}
