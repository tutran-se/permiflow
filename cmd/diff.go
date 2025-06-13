package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

var beforePath string
var afterPath string
var outDir string
var failOnLevel string

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show permission differences between two scans",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load both scan reports
		bindingsBefore, err := permiflow.LoadBindingsFromReport(beforePath)
		if err != nil {
			return fmt.Errorf("failed to load --before report: %w", err)
		}

		bindingsAfter, err := permiflow.LoadBindingsFromReport(afterPath)
		if err != nil {
			return fmt.Errorf("failed to load --after report: %w", err)
		}

		// Compute diff
		diff := permiflow.DiffBindings(bindingsBefore, bindingsAfter)

		// Print results
		permiflow.PrintDiff(diff)

		// If an output directory is specified, write the diff to files
		if outDir != "" {
			if err := os.MkdirAll(outDir, 0755); err != nil {
				return fmt.Errorf("failed to create output directory: %w", err)
			}

			if err := permiflow.WriteDiffMarkdown(diff, filepath.Join(outDir, "diff.md")); err != nil {
				return fmt.Errorf("failed to write Markdown: %w", err)
			}
			if err := permiflow.WriteDiffJSON(diff, filepath.Join(outDir, "diff.json")); err != nil {
				return fmt.Errorf("failed to write JSON: %w", err)
			}

			fmt.Println()

			if outDir != "." {
				outDir = filepath.Clean(outDir)
			}
			outDir = filepath.Base(outDir)

			// Print output directory
			fmt.Printf("Diff written to %s/\n", outDir)
			fmt.Println("Files: diff.md, diff.json")
		}

		// Check for failOnLevel
		if failOnLevel != "" {
			level := strings.ToUpper(failOnLevel)
			if permiflow.ContainsRiskLevel(diff, level) {
				cmd.SilenceUsage = true
				fmt.Printf("\nDetected %s risk binding(s). Failing as requested by --fail-on.\n", level)
				return fmt.Errorf("permission drift triggered failure: %s risk level found", level)
			}
		}

		return nil
	},
}

func init() {
	diffCmd.Flags().StringVar(&beforePath, "before", "", "Path to the older scan (baseline)")
	diffCmd.Flags().StringVar(&afterPath, "after", "", "Path to the newer scan (latest/current)")
	diffCmd.Flags().StringVar(&outDir, "out-dir", "", "Directory to write diff output (optional)")
	diffCmd.Flags().StringVar(&failOnLevel, "fail-on", "", "Exit with code 1 if any new or changed binding matches this risk level (e.g. 'high')")

	if err := diffCmd.MarkFlagRequired("before"); err != nil {
		log.Printf("warn: failed to mark --before as required: %v", err)
	}
	if err := diffCmd.MarkFlagRequired("after"); err != nil {
		log.Printf("warn: failed to mark --after as required: %v", err)
	}

	rootCmd.AddCommand(diffCmd)
}
