package permiflow

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ScanMetadata struct {
	ScanID      string   `json:"scan_id"`
	Timestamp   string   `json:"timestamp"`
	Kubeconfig  string   `json:"kubeconfig"`
	OutDir      string   `json:"out_dir"`
	Prefix      string   `json:"prefix"`
	NumBindings int      `json:"num_bindings"`
	Summary     Summary  `json:"summary"`
	OutputFiles []string `json:"output_files"`
}

func WriteMetadata(meta ScanMetadata, folder string) error {
	path := filepath.Join(folder, "metadata.json")
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
