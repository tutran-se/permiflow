package permiflow

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// JSONReport contains the full scan output: bindings and summary
type JSONReport struct {
	Summary  Summary         `json:"summary"`
	Bindings []AccessBinding `json:"bindings"`
}

// WriteJSON writes the results to a JSON file in the specified output directory.
func WriteJSON(bindings []AccessBinding, summary Summary, outDir, prefix string) error {
	report := JSONReport{
		Summary:  summary,
		Bindings: bindings,
	}

	// Ensure the output directory exists
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	filename := filepath.Join(outDir, prefix+".json")

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create JSON file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
