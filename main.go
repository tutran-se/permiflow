package main

import "fmt"

func main() {
  fmt.Println(emoji("🔍"), "Permiflow: Scanning RBAC...")
  client := GetKubeClient()
  bindings := ScanRBAC(client)
  WriteMarkdown(bindings, "report.md")
  WriteCSV(bindings, "report.csv")
  fmt.Printf("%s Report complete. %d bindings written.\n", emoji("✅"), len(bindings))
}