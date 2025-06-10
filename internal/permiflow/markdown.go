package permiflow

import (
	"fmt"
	"os"
	"strings"
)

func WriteMarkdown(bindings []AccessBinding, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create markdown file: %v\n", err)
		return
	}
	defer f.Close()

	// Header
	_, _ = fmt.Fprintln(f, "# Permiflow RBAC Audit Report")

	// Table of Contents
	_, _ = fmt.Fprintln(f, "## 📘 Table of Contents")
	for _, b := range bindings {
		slug := slugify(fmt.Sprintf("%s-%s", b.Subject, b.SubjectKind))
		_, _ = fmt.Fprintf(f, "- [%s (%s)](#%s)\n", b.Subject, b.SubjectKind, slug)
	}
	_, _ = fmt.Fprintln(f, "\n---")

	// Sections
	for _, b := range bindings {
		slug := slugify(fmt.Sprintf("%s-%s", b.Subject, b.SubjectKind))
		_, _ = fmt.Fprintf(f, "## <a name=\"%s\"></a>Subject: %s (%s)\n", slug, b.Subject, b.SubjectKind)
		_, _ = fmt.Fprintf(f, "- Namespace: `%s`\n", b.Namespace)
		_, _ = fmt.Fprintf(f, "- Role: `%s`\n", b.Role)
		_, _ = fmt.Fprintf(f, "- Verbs: `%s`\n", formatList(b.Verbs))
		_, _ = fmt.Fprintf(f, "- Resources: `%s`\n", formatList(b.Resources))
		_, _ = fmt.Fprintf(f, "- Scope: `%s`\n", b.Scope)
		_, _ = fmt.Fprintf(f, "- Risk Level: **%s**\n\n", b.RiskLevel)

		if b.Namespace == "" && b.SubjectKind != "ServiceAccount" {
			_, _ = fmt.Fprintf(f, "- ℹ️ Note: This is a cluster-wide subject (%s) granted access to this namespace.\n", b.SubjectKind)
		}
		_, _ = fmt.Fprintln(f, "")
	}
}

// Helper to slugify Markdown anchors
func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "/", "-")
	return s
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
