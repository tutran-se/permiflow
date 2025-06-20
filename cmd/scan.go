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
			log.SetFlags(0)
		} else {
			log.SetFlags(log.LstdFlags)
		}

		log.Println("Permiflow: Scanning RBAC...")

		client := permiflow.GetKubeClient(kubeconfig)

		start := time.Now()
		bindings, summary := permiflow.ScanRBAC(client)

		// Sort bindings by risk: HIGH > MEDIUM > LOW
		sort.Slice(bindings, func(i, j int) bool {
			return riskRank(bindings[i].RiskLevel) < riskRank(bindings[j].RiskLevel)
		})
		elapsed := time.Since(start)
		log.Printf("Scan completed in %.2fms", elapsed.Seconds()*1000)

		if dryRun {
			log.Println("Dry run enabled — no files will be written.")
		} else {
			scanTime := time.Now().UTC().Format("2006-01-02T15-04-05Z")
			randomSuffix := permiflow.ShortRandomString(8)
			scanID := fmt.Sprintf("%s--%s", scanTime, randomSuffix)

			versionedDir := filepath.Join(outputDir, scanID)
			if err := os.MkdirAll(versionedDir, 0755); err != nil {
				return fmt.Errorf("failed to create output directory: %w", err)
			}

			// META DATA
			meta := permiflow.ScanMetadata{
				ScanID:      scanID,
				Timestamp:   time.Now().UTC().Format(time.RFC3339),
				Kubeconfig:  kubeconfig,
				OutDir:      versionedDir,
				Prefix:      outputPrefix,
				NumBindings: len(bindings),
				Summary:     summary,
				OutputFiles: []string{
					outputPrefix + ".md",
					outputPrefix + ".csv",
					outputPrefix + ".json",
				},
			}

			if err := permiflow.WriteMetadata(meta, versionedDir); err != nil {
				log.Printf("Failed to write metadata.json: %v", err)
			} else {
				log.Printf("Metadata written to: %s", filepath.Join(versionedDir, "metadata.json"))
			}

			// MD
			mdPath := filepath.Join(versionedDir, outputPrefix+".md")
			permiflow.WriteMarkdown(bindings, mdPath, summary)
			log.Printf("Markdown written to: %s", mdPath)

			// CSV
			csvPath := filepath.Join(versionedDir, outputPrefix+".csv")
			permiflow.WriteCSV(bindings, csvPath)
			log.Printf("CSV written to: %s", csvPath)

			// JSON
			jsonPath := filepath.Join(versionedDir, outputPrefix+".json")
			if err := permiflow.WriteJSON(bindings, summary, versionedDir, outputPrefix); err != nil {
				return fmt.Errorf("failed to write JSON report: %w", err)
			}
			log.Printf("JSON written to: %s", jsonPath)

			// HISTORY
			contextName := permiflow.GetCurrentContext(kubeconfig)
			if err := permiflow.AppendToHistory(scanID, versionedDir, contextName); err != nil {
				log.Printf("Failed to update scan history: %v", err)
			} else {
				log.Println("Scan history updated: .permiflow/history.json")
			}
		}

		log.Printf("Report complete. %d bindings scanned.", len(bindings))

		// SUMMARY LOG
		log.Printf("Summary:")
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
	scanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "No files written, no scan history recorded (default: false)")
	scanCmd.Flags().StringVar(&outputDir, "out-dir", ".", "Directory to write reports into (default: current directory)")
	scanCmd.Flags().StringVar(&outputPrefix, "prefix", "report", "Base name for output files (without extension). Example: 'audit' → audit.md (default: 'report')")
	scanCmd.Flags().BoolVar(&logTimestamps, "log-timestamps", false, "Include timestamps in output (useful for debugging/logging)")
}
