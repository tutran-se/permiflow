package main

import "fmt"

func main() {
	fmt.Println(emoji("🔍"), "Permiflow: Scanning RBAC...")
	client := GetKubeClient()
	bindings, summary := ScanRBAC(client)

	fmt.Printf("%s Found %d RBAC bindings\n", emoji("✅"), len(bindings))
	WriteMarkdown(bindings, "report.md")
	WriteCSV(bindings, "report.csv")
	fmt.Printf("%s Report written to report.md and report.csv\n", emoji("📄"))

	fmt.Println()

	fmt.Printf("%s Summary:\n", emoji("📊"))
	fmt.Printf("- Found %d cluster-admin binding(s)\n", summary.ClusterAdminBindings)
	fmt.Printf("- Found %d wildcard verb usage(s)\n", summary.WildcardVerbs)
	fmt.Printf("- Found %d subject(s) with secrets access\n", summary.SecretsAccess)
}
