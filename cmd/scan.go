package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

var (
	kubeconfig    string
	markdownOut   bool
	csvOut        bool
	JSONOut       bool
	dryRun        bool
	outputDir     string
	outputPrefix  string
	logTimestamps bool
)

func riskRank(level string) int {
	switch level {
	case "HIGH":
		return 0
	case "MEDIUM":
		return 1
	default:
		return 2 // LOW or unknown
	}
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan Kubernetes RBAC roles and bindings",
	Example: `
	# Basic usage
	permiflow scan

	# Custom output
	permiflow scan --out-dir ./audit --prefix audit

	# Use specific kubeconfig
	permiflow scan --kubeconfig ~/.kube/config

	# Output is sorted by risk level: HIGH > MEDIUM > LOW
  	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !logTimestamps {
			log.SetFlags(0) // Clean output for users
		} else {
			log.SetFlags(log.LstdFlags) // Includes timestamps
		}

		if os.Getenv("PERMIFLOW_NO_EMOJI") == "true" {
			noEmoji = true
		}
		permiflow.NoEmoji = noEmoji

		log.Printf("%s Permiflow: Scanning RBAC...", permiflow.Emoji("🔍"))

		client := permiflow.GetKubeClient(kubeconfig)

		start := time.Now()
		bindings, summary := permiflow.ScanRBAC(client)

		// Sort bindings by risk: HIGH > MEDIUM > LOW
		sort.Slice(bindings, func(i, j int) bool {
			return riskRank(bindings[i].RiskLevel) < riskRank(bindings[j].RiskLevel)
		})
		elapsed := time.Since(start)
		log.Printf("%s Scan completed in %.2fms", permiflow.Emoji("⏱"), elapsed.Seconds()*1000)

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		if dryRun {
			log.Printf("%s Dry run enabled — no files will be written.", permiflow.Emoji("🧪"))
		} else {
			if markdownOut {
				mdPath := filepath.Join(outputDir, outputPrefix+".md")
				permiflow.WriteMarkdown(bindings, mdPath, summary)
				log.Printf("%s Markdown written to: %s", permiflow.Emoji("📄"), mdPath)
			}
			if csvOut {
				csvPath := filepath.Join(outputDir, outputPrefix+".csv")
				permiflow.WriteCSV(bindings, csvPath)
				log.Printf("%s CSV written to: %s", permiflow.Emoji("📊"), csvPath)
			}

			if JSONOut {
				jsonPath := filepath.Join(outputDir, outputPrefix+".json")
				if err := permiflow.WriteJSON(bindings, summary, outputDir, outputPrefix); err != nil {
					return fmt.Errorf("failed to write JSON report: %w", err)
				}
				log.Printf("%s JSON written to: %s", permiflow.Emoji("📦"), jsonPath)
			}
		}
		log.Printf("%s Report complete. %d bindings scanned.", permiflow.Emoji("✅"), len(bindings))

		log.Printf("%s Summary:", permiflow.Emoji("📊"))
		log.Printf("   - Found %d cluster-admin binding(s)", summary.ClusterAdminBindings)
		log.Printf("   - Found %d wildcard verb usage(s)", summary.WildcardVerbs)
		log.Printf("   - Found %d subject(s) with secrets access", summary.SecretsAccess)
		log.Printf("   - Found %d privilege escalation(s)", summary.PrivilegeEscalation)
		log.Printf("   - Found %d exec access(es)", summary.ExecAccess)
		log.Printf("   - Found %d config read secrets access(es)", summary.ConfigReadSecrets)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVar(&kubeconfig, "kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "Path to the kubeconfig file (default: $HOME/.kube/config)")
	scanCmd.Flags().BoolVar(&markdownOut, "markdown", true, "Enable Markdown report output (default: true; use --markdown=false to disable)")
	scanCmd.Flags().BoolVar(&csvOut, "csv", true, "Enable CSV report output (default: true; use --csv=false to disable)")
	scanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Run scan without writing any files (default: false)")
	scanCmd.Flags().StringVar(&outputDir, "out-dir", ".", "Directory to write reports into (default: current directory)")
	scanCmd.Flags().StringVar(&outputPrefix, "prefix", "report", "Base name for output files (without extension). Example: 'audit' → audit.md (default: 'report')")
	scanCmd.Flags().BoolVar(&logTimestamps, "log-timestamps", false, "Include timestamps in output (useful for debugging/logging)")
	scanCmd.Flags().BoolVar(&JSONOut, "json", true, "Enable JSON report output (default: true; use --json=false to disable)")
}
