package permiflow

import (
	"context"
	"fmt"
	"log"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type AccessBinding struct {
	Subject   string
	Role      string
	Namespace string
	Verbs     []string
	Resources []string
	RiskLevel string
	Scope     string
}

type Summary struct {
	ClusterAdminBindings int
	WildcardVerbs        int
	SecretsAccess        int
}

func ScanRBAC(clientset kubernetes.Interface, namespace string) ([]AccessBinding, Summary) {
	ctx := context.Background()
	var results []AccessBinding
	var summary Summary

	// Cache all ClusterRoles and Roles up front
	clusterRoles, err := clientset.RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("%s Failed to list ClusterRoles: %v", Emoji("❌"), err)
		return results, summary
	}
	clusterRoleMap := make(map[string]rbacv1.ClusterRole)
	for _, cr := range clusterRoles.Items {
		clusterRoleMap[cr.Name] = cr
	}

	roleMap := make(map[string]rbacv1.Role)
	if namespace != "" {
		roles, err := clientset.RbacV1().Roles(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Printf("%s Failed to list Roles in namespace %s: %v", Emoji("❌"), namespace, err)
			return results, summary
		}
		for _, r := range roles.Items {
			roleMap[r.Name] = r
		}
	}

	// Scan ClusterRoleBindings
	crbs, err := clientset.RbacV1().ClusterRoleBindings().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Printf("%s Failed to list ClusterRoleBindings: %v", Emoji("❌"), err)
		return results, summary
	}
	fmt.Printf("%s Found %d ClusterRoleBindings\n", Emoji("🔍"), len(crbs.Items))

	for _, crb := range crbs.Items {
		role, ok := clusterRoleMap[crb.RoleRef.Name]
		if !ok {
			log.Printf("%s Skipping ClusterRoleBinding %s: ClusterRole %s not found\n", Emoji("⚠️"), crb.Name, crb.RoleRef.Name)
			continue
		}
		extractBindingsFromRole(crb.Subjects, crb.RoleRef.Name, "Cluster", role.Rules, &results, &summary)
	}

	// Scan RoleBindings (if namespace provided)
	if namespace != "" {
		_, err := clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			log.Printf("%s Warning: Namespace '%s' does not exist. No RoleBindings will be found.\n", Emoji("⚠️"), namespace)
			return results, summary
		} else if err != nil {
			log.Printf("%s Failed to validate namespace existence: %v\n", Emoji("❌"), err)
			return results, summary
		}

		rbs, err := clientset.RbacV1().RoleBindings(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Printf("%s Failed to list RoleBindings in %s: %v\n", Emoji("❌"), namespace, err)
			return results, summary
		}
		fmt.Printf("%s Found %d RoleBindings in namespace: %s\n", Emoji("🔍"), len(rbs.Items), namespace)

		for _, rb := range rbs.Items {
			var rules []rbacv1.PolicyRule
			switch rb.RoleRef.Kind {
			case "ClusterRole":
				role, ok := clusterRoleMap[rb.RoleRef.Name]
				if !ok {
					log.Printf("%s Skipping RoleBinding %s: ClusterRole %s not found\n", Emoji("⚠️"), rb.Name, rb.RoleRef.Name)
					continue
				}
				rules = role.Rules
			case "Role":
				role, ok := roleMap[rb.RoleRef.Name]
				if !ok {
					log.Printf("%s Skipping RoleBinding %s: Role %s not found in namespace %s\n", Emoji("⚠️"), rb.Name, rb.RoleRef.Name, namespace)
					continue
				}
				rules = role.Rules
			default:
				log.Printf("%s Unknown RoleRef kind: %s in RoleBinding %s\n", Emoji("⚠️"), rb.RoleRef.Kind, rb.Name)
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
			binding := AccessBinding{
				Subject:   subject.Name,
				Role:      roleName,
				Namespace: subject.Namespace,
				Verbs:     rule.Verbs,
				Resources: rule.Resources,
				RiskLevel: ClassifyRisk(rule.Verbs, rule.Resources),
				Scope:     scope,
			}
			*out = append(*out, binding)
			updateSummary(summary, binding)
		}
	}
}

func updateSummary(summary *Summary, binding AccessBinding) {
	if binding.Role == "cluster-admin" {
		summary.ClusterAdminBindings++
	}
	if contains(binding.Verbs, "*") {
		summary.WildcardVerbs++
	}
	if contains(binding.Resources, "secrets") {
		summary.SecretsAccess++
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

func ClassifyRisk(verbs, resources []string) string {
	if contains(verbs, "*") || contains(resources, "*") {
		return "HIGH"
	}
	if contains(resources, "secrets") {
		return "MEDIUM"
	}
	return "LOW"
}
