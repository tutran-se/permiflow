package permiflow

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteCSV(bindings []AccessBinding, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create CSV file: %v\n", err)
		return
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	if err := writer.Write([]string{"Subject", "Kind", "Namespace", "Role", "Verbs", "Resources", "Scope", "RiskLevel"}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write CSV header: %v\n", err)
		return
	}

	for _, b := range bindings {
		row := []string{
			b.Subject,
			b.SubjectKind,
			b.Namespace,
			b.Role,
			join(b.Verbs, " "),
			join(b.Resources, " "),
			b.Scope,
			b.RiskLevel,
		}
		if err := writer.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write CSV row: %v\n", err)
		}
	}
}
