package permiflow

import (
	"fmt"
	"os"
	"strings"
)

func WriteMarkdown(bindings []AccessBinding, filename string, summary Summary) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create markdown file: %v\n", err)
		return
	}
	defer f.Close()

	// Header
	_, _ = fmt.Fprintln(f, "# Permiflow RBAC Audit Report")

	// Summary
	_, _ = fmt.Fprintln(f, "## Summary")
	_, _ = fmt.Fprintf(f, "- Total bindings scanned: **%d**\n", len(bindings))
	_, _ = fmt.Fprintf(f, "- Found %d cluster-admin binding(s)\n", summary.ClusterAdminBindings)
	_, _ = fmt.Fprintf(f, "- Found %d wildcard verb usage(s)\n", summary.WildcardVerbs)
	_, _ = fmt.Fprintf(f, "- Found %d subject(s) with secrets access\n", summary.SecretsAccess)
	_, _ = fmt.Fprintf(f, "- Found %d privilege escalation(s)\n", summary.PrivilegeEscalation)
	_, _ = fmt.Fprintf(f, "- Found %d exec access(es)\n", summary.ExecAccess)
	_, _ = fmt.Fprintf(f, "- Found %d config read secrets access(es)\n", summary.ConfigReadSecrets)
	_, _ = fmt.Fprintln(f, "\n---")

	// Risk Levels
	_, _ = fmt.Fprintln(f, "## Risk Levels")
	_, _ = fmt.Fprintln(f, "- **HIGH**: Wildcard verbs or resources, privilege escalation risks")
	_, _ = fmt.Fprintln(f, "- **MEDIUM**: Sensitive resources with non-wildcard verbs")
	_, _ = fmt.Fprintln(f, "- **LOW**: Non-sensitive resources with non-wildcard verbs")
	_, _ = fmt.Fprintln(f, "\n---")

	// Table of Contents
	_, _ = fmt.Fprintln(f, "## Table of Contents")
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
		_, _ = fmt.Fprintf(f, "- Risk Level: **%s**\n", b.RiskLevel)
		_, _ = fmt.Fprintf(f, "- Reason: %s\n\n", b.Reason)

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
