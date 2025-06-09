package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v0.3.3"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Permiflow version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Permiflow version:", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
