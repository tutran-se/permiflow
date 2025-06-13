package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the MCP server
type Config struct {
	Transport  string
	HTTPPort   int
	Debug      bool
	Kubeconfig string
	Context    string
}

// DefaultConfig returns a new config with default values
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
	if transport := os.Getenv("MCP_TRANSPORT"); transport != "" {
		c.Transport = transport
	}
	if port := os.Getenv("MCP_HTTP_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			c.HTTPPort = p
		}
	}
	if debug := os.Getenv("MCP_DEBUG"); debug == "true" {
		c.Debug = true
	}
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		c.Kubeconfig = kubeconfig
	}
	if context := os.Getenv("KUBE_CONTEXT"); context != "" {
		c.Context = context
	}

	// Default kubeconfig location if not set
	if c.Kubeconfig == "" {
		if home := os.Getenv("HOME"); home != "" {
			c.Kubeconfig = home + "/.kube/config"
		}
	}
}
