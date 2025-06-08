package permiflow

import (
	"fmt"
	"os"
)

func WriteMarkdown(bindings []AccessBinding, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create markdown file: %v\n", err)
		return
	}
	defer f.Close()

	_, _ = fmt.Fprintln(f, "# Permiflow RBAC Audit Report")

	for _, b := range bindings {
		_, _ = fmt.Fprintf(f, "## Subject: %s\n", b.Subject)
		_, _ = fmt.Fprintf(f, "- Namespace: `%s`\n", b.Namespace)
		_, _ = fmt.Fprintf(f, "- Role: `%s`\n", b.Role)
		_, _ = fmt.Fprintf(f, "- Verbs: `%s`\n", formatList(b.Verbs))
		_, _ = fmt.Fprintf(f, "- Resources: `%s`\n", formatList(b.Resources))
		_, _ = fmt.Fprintf(f, "- Scope: `%s`\n", b.Scope)
		_, _ = fmt.Fprintf(f, "- Risk Level: **%s**\n\n", b.RiskLevel)
	}
}

func formatList(items []string) string {
	return fmt.Sprintf("[%s]", join(items, " "))
}

func join(slice []string, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := slice[0]
	for _, s := range slice[1:] {
		result += sep + s
	}
	return result
}
