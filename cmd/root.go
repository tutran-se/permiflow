package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dryRun     bool
	kubeconfig string
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

// add global flags to the root command
func init() {
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Enable dry run mode (no changes made)")
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "Path to kubeconfig file (default uses KUBECONFIG env or ~/.kube/config)")
}
