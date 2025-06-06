package main

import "fmt"

func main() {
  fmt.Println("🔍 Permiflow: Scanning RBAC...")
  client := GetKubeClient()
  bindings := ScanRBAC(client)
  WriteMarkdown(bindings, "report.md")
  WriteCSV(bindings, "report.csv")
  fmt.Printf("✅ Report complete. %d bindings written.\n", len(bindings))
}