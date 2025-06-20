package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dryRun bool
)

// rootCmd is the base command called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "permiflow",
	Short: "Permiflow scans Kubernetes RBAC bindings and generates audit reports",
	Long: `Permiflow is a focused, read-only CLI tool that scans Kubernetes RBAC bindings
and generates Markdown and CSV reports — perfect for security reviews, audits, and compliance.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
