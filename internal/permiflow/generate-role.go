package permiflow

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

// GenerateRBACRoleOptions defines input flags
type GenerateRBACRoleOptions struct {
	AllowedVerbs      []string
	ExcludedResources []string
	OutputPath        string
	RoleName          string
	RoleKind          string // "ClusterRole" or "Role"
	Namespace         string // required if RoleKind == "Role"
}

// Minimal catalog of common resources grouped by API group
var defaultResources = map[string][]string{
	"":                          {"pods", "services", "secrets", "configmaps", "nodes", "persistentvolumeclaims", "pods/exec", "pods/proxy", "pods/portforward"},
	"apps":                      {"deployments", "statefulsets", "daemonsets"},
	"batch":                     {"jobs", "cronjobs"},
	"rbac.authorization.k8s.io": {"roles", "rolebindings", "clusterroles", "clusterrolebindings"},
}

// GenerateRBACRole builds and writes the ClusterRole or Role YAML to disk
func GenerateRBACRole(opts GenerateRBACRoleOptions) error {
	f, err := os.Create(opts.OutputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()
	return encodeRBACRole(opts, f)
}

// GenerateRBACRoleToStdout prints the ClusterRole or Role YAML to stdout
func GenerateRBACRoleToStdout(opts GenerateRBACRoleOptions) error {
	return encodeRBACRole(opts, os.Stdout)
}

func encodeRBACRole(opts GenerateRBACRoleOptions, out io.Writer) error {
	excludeSet := make(map[string]bool)
	for _, r := range opts.ExcludedResources {
		excludeSet[strings.ToLower(r)] = true
	}

	rules := []rbacv1.PolicyRule{}

	// 1. Sort API groups deterministically
	apiGroups := make([]string, 0, len(defaultResources))
	for group := range defaultResources {
		apiGroups = append(apiGroups, group)
	}
	sort.Strings(apiGroups)

	// 2. Build rules with sorted resources
	for _, apiGroup := range apiGroups {
		resources := defaultResources[apiGroup]
		allowed := []string{}
		for _, r := range resources {
			if !excludeSet[strings.ToLower(r)] {
				allowed = append(allowed, r)
			}
		}
		if len(allowed) > 0 {
			sort.Strings(allowed) // sort allowed resources per rule
			rules = append(rules, rbacv1.PolicyRule{
				APIGroups: []string{apiGroup},
				Resources: allowed,
				Verbs:     opts.AllowedVerbs,
			})
		}
	}

	// 3. Sort entire rules list by group + resource + verb
	sort.SliceStable(rules, func(i, j int) bool {
		a, b := rules[i], rules[j]

		groupA := ""
		if len(a.APIGroups) > 0 {
			groupA = a.APIGroups[0]
		}
		groupB := ""
		if len(b.APIGroups) > 0 {
			groupB = b.APIGroups[0]
		}
		if groupA != groupB {
			return groupA < groupB
		}

		resA := ""
		if len(a.Resources) > 0 {
			resA = a.Resources[0]
		}
		resB := ""
		if len(b.Resources) > 0 {
			resB = b.Resources[0]
		}
		if resA != resB {
			return resA < resB
		}

		verbA := ""
		if len(a.Verbs) > 0 {
			verbA = a.Verbs[0]
		}
		verbB := ""
		if len(b.Verbs) > 0 {
			verbB = b.Verbs[0]
		}
		return verbA < verbB
	})

	// 4. Encode final object
	yamlSerializer := json.NewYAMLSerializer(json.DefaultMetaFactory, nil, nil)

	switch strings.ToLower(opts.RoleKind) {
	case "role":
		if opts.Namespace == "" {
			return fmt.Errorf("namespace is required when RoleKind is 'Role'")
		}
		role := rbacv1.Role{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Role",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      opts.RoleName,
				Namespace: opts.Namespace,
			},
			Rules: rules,
		}
		role.CreationTimestamp = metav1.Time{}
		return yamlSerializer.Encode(&role, out)
	default:
		clusterRole := rbacv1.ClusterRole{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ClusterRole",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: opts.RoleName,
			},
			Rules: rules,
		}
		clusterRole.CreationTimestamp = metav1.Time{}
		return yamlSerializer.Encode(&clusterRole, out)
	}
}

// GenerateRBACRoleToStdoutWithWriter is test-friendly version for internal use
func GenerateRBACRoleToStdoutWithWriter(opts GenerateRBACRoleOptions, out *bytes.Buffer) error {
	return encodeRBACRole(opts, out)
}

func DefaultResourceList() []string {
	flat := []string{}
	for _, res := range defaultResources {
		flat = append(flat, res...)
	}
	return flat
}
