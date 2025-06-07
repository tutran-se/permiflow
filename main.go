package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const Version = "v0.1.3"

var (
	kubeconfig   string
	namespace    string
	markdownOut  bool
	csvOut       bool
	dryRun       bool
	noEmoji      bool
	outputDir    string
	outputPrefix string
	showVersion  bool
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "Path to the kubeconfig file")
	flag.StringVar(&namespace, "namespace", "", "Limit scan to a specific namespace (optional)")
	flag.BoolVar(&markdownOut, "markdown", true, "Enable Markdown report output")
	flag.BoolVar(&csvOut, "csv", true, "Enable CSV report output")
	flag.BoolVar(&dryRun, "dry-run", false, "Run scan without writing any files")
	flag.BoolVar(&noEmoji, "plain", false, "Use plain text output (no emoji icons)")
	flag.StringVar(&outputDir, "out-dir", ".", "Directory to write reports into")
	flag.StringVar(&outputPrefix, "prefix", "report", "Filename prefix for output files")
	flag.BoolVar(&showVersion, "version", false, "Print version and exit")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println("Permiflow version:", Version)
		os.Exit(0)
	}

	if os.Getenv("PERMIFLOW_NO_EMOJI") == "true" {
		noEmoji = true
	}

	emoji := func(s string) string {
		if noEmoji {
			return ""
		}
		return s
	}

	fmt.Println(emoji("🔍") + " Permiflow: Scanning RBAC...")

	client := GetKubeClient(kubeconfig)
	bindings, summary := ScanRBAC(client, namespace)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	if dryRun {
		fmt.Println(emoji("🧪") + " Dry run enabled — no files will be written.")
	} else {
		if markdownOut {
			mdPath := filepath.Join(outputDir, outputPrefix+".md")
			WriteMarkdown(bindings, mdPath)
			fmt.Println(emoji("📄")+" Markdown written to:", mdPath)
		}
		if csvOut {
			csvPath := filepath.Join(outputDir, outputPrefix+".csv")
			WriteCSV(bindings, csvPath)
			fmt.Println(emoji("📊")+" CSV written to:", csvPath)
		}
	}

	fmt.Printf("%s Summary:\n", emoji("📊"))
	fmt.Printf("- Found %d cluster-admin binding(s)\n", summary.ClusterAdminBindings)
	fmt.Printf("- Found %d wildcard verb usage(s)\n", summary.WildcardVerbs)
	fmt.Printf("- Found %d subject(s) with secrets access\n", summary.SecretsAccess)

	fmt.Printf("%s Report complete. %d bindings scanned.\n", emoji("✅"), len(bindings))
}
