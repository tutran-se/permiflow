package permiflow

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type BindingChange struct {
	Before AccessBinding
	After  AccessBinding
}

type DiffResult struct {
	Added   []AccessBinding
	Removed []AccessBinding
	Changed []BindingChange
}

func LoadBindingsFromReport(path string) ([]AccessBinding, error) {
	reportPath := filepath.Clean(path)
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var report JSONReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return report.Bindings, nil
}

func DiffBindings(before, after []AccessBinding) DiffResult {
	result := DiffResult{}
	beforeMap := make(map[string]AccessBinding)
	afterMap := make(map[string]AccessBinding)

	for _, b := range before {
		beforeMap[bindingKey(b)] = b
	}
	for _, b := range after {
		afterMap[bindingKey(b)] = b
	}

	for key, bNew := range afterMap {
		if bOld, found := beforeMap[key]; !found {
			result.Added = append(result.Added, bNew)
		} else if !equalRuleSets(bOld, bNew) {
			result.Changed = append(result.Changed, BindingChange{
				Before: bOld,
				After:  bNew,
			})
		}
	}

	for key, bOld := range beforeMap {
		if _, found := afterMap[key]; !found {
			result.Removed = append(result.Removed, bOld)
		}
	}

	return result
}

func PrintDiff(diff DiffResult) {
	fmt.Println("RBAC Diff Summary")
	fmt.Println("------------------")

	for _, added := range diff.Added {
		ns := added.Namespace
		if added.Scope == "Cluster" {
			ns = "cluster"
		}
		fmt.Printf("+ %s gained %s access to %s in %s (via Role: %s) [%s]\n",
			added.Subject,
			formatSlice(added.Verbs),
			formatSlice(added.Resources),
			ns,
			added.Role,
			added.RiskLevel,
		)
	}

	for _, removed := range diff.Removed {
		ns := removed.Namespace
		if removed.Scope == "Cluster" {
			ns = "cluster"
		}
		fmt.Printf("- %s lost %s access to %s in %s (via Role: %s) [%s]\n",
			removed.Subject,
			formatSlice(removed.Verbs),
			formatSlice(removed.Resources),
			ns,
			removed.Role,
			removed.RiskLevel,
		)
	}

	for _, change := range diff.Changed {
		ns := change.Before.Namespace
		if change.Before.Scope == "Cluster" {
			ns = "cluster"
		}
		fmt.Printf("~ %s's access in %s changed (Role: %s):\n", change.Before.Subject, ns, change.Before.Role)
		fmt.Printf("    BEFORE: %s on %s [%s]\n",
			formatSlice(change.Before.Verbs),
			formatSlice(change.Before.Resources),
			change.Before.RiskLevel,
		)
		fmt.Printf("    AFTER:  %s on %s [%s]\n",
			formatSlice(change.After.Verbs),
			formatSlice(change.After.Resources),
			change.After.RiskLevel,
		)
	}

	// if no changes, print a summary
	if len(diff.Added) == 0 && len(diff.Removed) == 0 && len(diff.Changed) == 0 {
		fmt.Println("No changes detected.")
	} else {
		fmt.Println()
		fmt.Printf("Added: %d, Removed: %d, Changed: %d\n",
			len(diff.Added), len(diff.Removed), len(diff.Changed))

	}
}
func bindingKey(b AccessBinding) string {
	ns := b.Namespace
	if b.Scope == "Cluster" {
		ns = ""
	}
	return fmt.Sprintf("%s|%s|%s|%s|%s", b.SubjectKind, b.Subject, ns, b.Role, b.Scope)
}

func equalRuleSets(a, b AccessBinding) bool {
	return stringSliceEqual(a.Verbs, b.Verbs) &&
		stringSliceEqual(a.Resources, b.Resources) &&
		a.RiskLevel == b.RiskLevel
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]bool)
	for _, val := range a {
		m[val] = true
	}
	for _, val := range b {
		if !m[val] {
			return false
		}
	}
	return true
}

func formatSlice(items []string) string {
	if len(items) == 1 {
		return items[0]
	}
	return fmt.Sprintf("%v", items)
}

func WriteDiffMarkdown(diff DiffResult, path string) error {
	var b strings.Builder
	b.WriteString("# RBAC Diff Summary\n\n")

	for _, added := range diff.Added {
		ns := added.Namespace
		if added.Scope == "Cluster" {
			ns = "cluster"
		}
		fmt.Fprintf(&b, "- **ADDED:** `%s` gained `%s` access to `%s` in `%s` (via Role: `%s`) [%s]\n",
			added.Subject,
			formatSlice(added.Verbs),
			formatSlice(added.Resources),
			ns,
			added.Role,
			added.RiskLevel,
		)
	}

	for _, removed := range diff.Removed {
		ns := removed.Namespace
		if removed.Scope == "Cluster" {
			ns = "cluster"
		}
		fmt.Fprintf(&b, "- **REMOVED:** `%s` lost `%s` access to `%s` in `%s` (via Role: `%s`) [%s]\n",
			removed.Subject,
			formatSlice(removed.Verbs),
			formatSlice(removed.Resources),
			ns,
			removed.Role,
			removed.RiskLevel,
		)
	}

	for _, change := range diff.Changed {
		ns := change.Before.Namespace
		if change.Before.Scope == "Cluster" {
			ns = "cluster"
		}
		fmt.Fprintf(&b, "- **CHANGED:** `%s` in `%s` (Role: `%s`):\n", change.Before.Subject, ns, change.Before.Role)
		fmt.Fprintf(&b, "  - BEFORE: `%s` on `%s` [%s]\n",
			formatSlice(change.Before.Verbs),
			formatSlice(change.Before.Resources),
			change.Before.RiskLevel,
		)
		fmt.Fprintf(&b, "  - AFTER:  `%s` on `%s` [%s]\n",
			formatSlice(change.After.Verbs),
			formatSlice(change.After.Resources),
			change.After.RiskLevel,
		)
	}

	// if no changes, print a summary
	if len(diff.Added) == 0 && len(diff.Removed) == 0 && len(diff.Changed) == 0 {
		b.WriteString("No changes detected.\n")
	} else {
		b.WriteString("\n")
		fmt.Fprintf(&b, "Added: %d, Removed: %d, Changed: %d\n",
			len(diff.Added), len(diff.Removed), len(diff.Changed))
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func WriteDiffJSON(diff DiffResult, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(diff, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func ContainsRiskLevel(diff DiffResult, risk string) bool {
	for _, b := range diff.Added {
		if b.RiskLevel == risk {
			return true
		}
	}
	for _, c := range diff.Changed {
		if c.After.RiskLevel == risk {
			return true
		}
	}
	return false
}
