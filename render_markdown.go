package main

import (
	"fmt"
	"os"
)

func WriteMarkdown(bindings []AccessBinding, filename string) {
  f, _ := os.Create(filename)
  defer f.Close()

  fmt.Fprintf(f, "# Kubernetes RBAC Report\n\n")
  for _, b := range bindings {
    fmt.Fprintf(f, "## Subject: %s (Namespace: %s)\n", b.Subject, b.Namespace)
    fmt.Fprintf(f, "- Role: `%s`\n", b.Role)
    fmt.Fprintf(f, "- Verbs: `%v`\n", b.Verbs)
    fmt.Fprintf(f, "- Resources: `%v`\n", b.Resources)
    fmt.Fprintf(f, "- Risk Level: **%s**\n\n", b.RiskLevel)
  }
}