package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show scan history from .permiflow/history.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		history, err := permiflow.LoadHistory()
		if err != nil {
			return fmt.Errorf("failed to load scan history: %w", err)
		}
		if len(history) == 0 {
			fmt.Println("No scan history found. Run a scan first.")
			return nil
		}

		fmt.Println("Scan History")
		fmt.Println("--------------------------------------------")
		for _, entry := range history {
			fmt.Printf("Scan ID:    %s\n", entry.ScanID)
			fmt.Printf("Path:       %s\n", entry.Path)
			if entry.Context != "" {
				fmt.Printf("Context:    %s\n", entry.Context)
			} else {
				fmt.Println("Context:    (default)")
			}
			fmt.Printf("Timestamp:  %s\n\n", entry.Timestamp)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
