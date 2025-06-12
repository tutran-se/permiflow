package permiflow

import (
	"encoding/json"
	"os"
	"time"
)

const historyPath = ".permiflow/history.json"

type HistoryEntry struct {
	ScanID    string `json:"scan_id"`
	Path      string `json:"path"`      // Folder that contains report + metadata
	Timestamp string `json:"timestamp"` // ISO timestamp
	Context   string `json:"context"`   // kubeconfig context used for scan
}

func AppendToHistory(scanID, scanPath string, contextName string) error {
	// Ensure .permiflow folder exists
	if err := os.MkdirAll(".permiflow", 0755); err != nil {
		return err
	}

	var history []HistoryEntry

	// Load existing history if it exists
	if data, err := os.ReadFile(historyPath); err == nil {
		_ = json.Unmarshal(data, &history) // best effort; ignore failure here
	}

	// Add the new scan
	entry := HistoryEntry{
		ScanID:    scanID,
		Path:      scanPath,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	history = append(history, entry)

	// Optional: keep only latest N entries (e.g., 50)
	const maxEntries = 50
	if len(history) > maxEntries {
		history = history[len(history)-maxEntries:]
	}

	// Save back to disk
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(historyPath, data, 0644)
}

func GetLastTwoScans() (*HistoryEntry, *HistoryEntry, error) {
	data, err := os.ReadFile(historyPath)
	if err != nil {
		return nil, nil, err
	}

	var history []HistoryEntry
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, nil, err
	}

	if len(history) < 2 {
		return nil, nil, nil // Not enough scans to compare
	}

	// Return oldest-to-newest
	return &history[len(history)-2], &history[len(history)-1], nil
}

func LoadHistory() ([]HistoryEntry, error) {
	data, err := os.ReadFile(historyPath)
	if err != nil {
		// No file? Return empty slice, not an error.
		if os.IsNotExist(err) {
			return []HistoryEntry{}, nil
		}
		return nil, err
	}

	var history []HistoryEntry
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}
	return history, nil
}
