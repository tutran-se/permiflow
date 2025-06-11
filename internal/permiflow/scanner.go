package permiflow

import (
	"context"
	"log"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type AccessBinding struct {
	Subject     string
	SubjectKind string // e.g., "User", "ServiceAccount", "Group"
	Role        string
	Namespace   string
	Verbs       []string
	Resources   []string
	Scope       string // "Cluster" or "Namespaced"
	RiskLevel   string
	Reason      string // Explanation for risk classification
}

type Summary struct {
	ClusterAdminBindings int
	WildcardVerbs        int
	SecretsAccess        int
	PrivilegeEscalation  int
	ExecAccess           int
	ConfigReadSecrets    int
}

func ScanRBAC(clientset kubernetes.Interface) ([]AccessBinding, Summary) {
	ctx := context.Background()
	var results []AccessBinding
	var summary Summary

	// Cache all ClusterRoles
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("%s Failed to list ClusterRoles: %v", Emoji("❌"), err)
		return results, summary
	}
	clusterRoleMap := make(map[string]rbacv1.ClusterRole)
	for _, cr := range clusterRoles.Items {
		clusterRoleMap[cr.Name] = cr
	}

	// Scan ClusterRoleBindings (always)
	crbs, err := clientset.RbacV1().ClusterRoleBindings().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("%s Failed to list ClusterRoleBindings: %v", Emoji("❌"), err)
		return results, summary
	}
	log.Printf("%s Found %d ClusterRoleBindings", Emoji("🔍"), len(crbs.Items))
	for _, crb := range crbs.Items {
		role, ok := clusterRoleMap[crb.RoleRef.Name]
		if !ok {
			log.Printf("%s Skipping ClusterRoleBinding %s: ClusterRole %s not found", Emoji("⚠️"), crb.Name, crb.RoleRef.Name)
			continue
		}

		extractBindingsFromRole(crb.Subjects, crb.RoleRef.Name, "Cluster", role.Rules, &results, &summary)
	}

	// Namespace scan: either just one, or all
	namespaces := []string{}
	nsList, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("%s Failed to list namespaces: %v", Emoji("❌"), err)
		return results, summary
	}
	for _, ns := range nsList.Items {
		namespaces = append(namespaces, ns.Name)
	}
	log.Printf("%s Scanning RoleBindings in %d namespaces", Emoji("📦"), len(namespaces))

	for _, ns := range namespaces {
		roleList, err := clientset.RbacV1().Roles(ns).List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Printf("%s Failed to list Roles in namespace %s: %v", Emoji("❌"), ns, err)
			continue
		}
		roleMap := make(map[string]rbacv1.Role)
		for _, r := range roleList.Items {
			roleMap[r.Name] = r
		}

		rbs, err := clientset.RbacV1().RoleBindings(ns).List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Printf("%s Failed to list RoleBindings in %s: %v", Emoji("❌"), ns, err)
			continue
		}
		log.Printf("%s Found %d RoleBindings in namespace: %s", Emoji("🔍"), len(rbs.Items), ns)

		for _, rb := range rbs.Items {
			var rules []rbacv1.PolicyRule
			switch rb.RoleRef.Kind {
			case "ClusterRole":
				role, ok := clusterRoleMap[rb.RoleRef.Name]
				if !ok {
					log.Printf("%s Skipping RoleBinding %s: ClusterRole %s not found", Emoji("⚠️"), rb.Name, rb.RoleRef.Name)
					continue
				}
				rules = role.Rules
			case "Role":
				role, ok := roleMap[rb.RoleRef.Name]
				if !ok {
					log.Printf("%s Skipping RoleBinding %s: Role %s not found in namespace %s", Emoji("⚠️"), rb.Name, rb.RoleRef.Name, ns)
					continue
				}
				rules = role.Rules
			default:
				log.Printf("%s Unknown RoleRef kind: %s in RoleBinding %s", Emoji("⚠️"), rb.RoleRef.Kind, rb.Name)
				continue
			}
			extractBindingsFromRole(rb.Subjects, rb.RoleRef.Name, "Namespaced", rules, &results, &summary)
		}
	}

	return results, summary
}

func extractBindingsFromRole(subjects []rbacv1.Subject, roleName, scope string, rules []rbacv1.PolicyRule, out *[]AccessBinding, summary *Summary) {
	for _, subject := range subjects {
		for _, rule := range rules {
			riskLevel, reason := ClassifyRisk(rule.Verbs, rule.Resources)
			binding := AccessBinding{
				Subject:     subject.Name,
				SubjectKind: subject.Kind,
				Role:        roleName,
				Namespace:   subject.Namespace,
				Verbs:       rule.Verbs,
				Resources:   rule.Resources,
				Scope:       scope,
				RiskLevel:   riskLevel,
				Reason:      reason,
			}
			*out = append(*out, binding)
			updateSummary(summary, binding)
		}
	}
}

func updateSummary(summary *Summary, binding AccessBinding) {
	var verbs = binding.Verbs
	var resources = binding.Resources
	var role = binding.Role

	if role == "cluster-admin" {
		summary.ClusterAdminBindings++
	}
	if contains(verbs, "*") {
		summary.WildcardVerbs++
	}
	if contains(resources, "secrets") {
		summary.SecretsAccess++
	}
	if (contains(verbs, "create") && contains(resources, "rolebindings")) ||
		(contains(verbs, "update") && contains(resources, "roles")) {
		summary.PrivilegeEscalation++
	}
	if contains(resources, "pods/exec") || contains(resources, "nodes") {
		summary.ExecAccess++
	}
	if contains(verbs, "get") && contains(resources, "configmaps") {
		summary.ConfigReadSecrets++
	}
}

func contains(list []string, item string) bool {
	for _, x := range list {
		if x == item {
			return true
		}
	}
	return false
}

func ClassifyRisk(verbs, resources []string) (string, string) {
	// High Risk Checks
	if contains(verbs, "*") || contains(resources, "*") {
		return "HIGH", "Wildcard verb or resource detected"
	}
	if contains(verbs, "create") && contains(resources, "rolebindings") {
		return "HIGH", "Can create rolebindings (potential privilege escalation)"
	}
	if contains(verbs, "update") && contains(resources, "roles") {
		return "HIGH", "Can update roles (potential privilege escalation)"
	}
	if contains(verbs, "update") && contains(resources, "rolebindings") {
		return "HIGH", "Can update rolebindings (potential privilege escalation)"
	}
	if contains(verbs, "create") && contains(resources, "roles") {
		return "HIGH", "Can create roles (potential privilege escalation)"
	}
	if contains(verbs, "create") && contains(resources, "clusterroles") {
		return "HIGH", "Can create clusterroles (potential privilege escalation)"
	}
	if contains(verbs, "update") && contains(resources, "clusterroles") {
		return "HIGH", "Can update clusterroles (potential privilege escalation)"
	}
	if contains(resources, "pods/exec") {
		return "HIGH", "Access to pods/exec (enables remote command execution)"
	}
	if contains(resources, "nodes") {
		return "HIGH", "Access to nodes (host-level access)"
	}

	// Medium Risk Checks
	if contains(resources, "secrets") {
		return "MEDIUM", "Can access secrets (sensitive credential exposure)"
	}
	if contains(verbs, "get") && contains(resources, "configmaps") {
		return "MEDIUM", "Can read configmaps (often contains credentials or API keys)"
	}

	// Default to Low Risk
	return "LOW", "No known high-risk permissions"
}
