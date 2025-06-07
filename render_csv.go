package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func WriteCSV(bindings []AccessBinding, filename string) {
	f, _ := os.Create(filename)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Header
	writer.Write([]string{
		"Subject", "Namespace", "Role", "Verbs", "Resources", "RiskLevel", "Scope",
	})

	for _, b := range bindings {
		record := []string{
			b.Subject,
			b.Namespace,
			b.Role,
			strings.Join(b.Verbs, "|"),
			strings.Join(b.Resources, "|"),
			b.RiskLevel,
			b.Scope,
		}
		writer.Write(record)
	}
}
