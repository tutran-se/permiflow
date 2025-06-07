package main

import (
	"context"
	"fmt"

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
	Scope     string // Cluster vs Namespaced scope
}

type Summary struct {
	ClusterAdminBindings int
	WildcardVerbs        int
	SecretsAccess        int
}

func ScanRBAC(clientset *kubernetes.Clientset, namespace string) ([]AccessBinding, Summary) {
	var results []AccessBinding
	var summary Summary
	ctx := context.TODO()

	// 🔹 1. Scan ClusterRoleBindings
	crbs, err := clientset.RbacV1().ClusterRoleBindings().List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%s Failed to list ClusterRoleBindings: %v\n", emoji("❌"), err)
		return results, summary
	}
	fmt.Printf("%s Found %d ClusterRoleBindings\n", emoji("🔍"), len(crbs.Items))

	for _, crb := range crbs.Items {
		role, err := clientset.RbacV1().ClusterRoles().Get(ctx, crb.RoleRef.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("%s Skipping ClusterRoleBinding %s: ClusterRole %s not found\n", emoji("⚠️"), crb.Name, crb.RoleRef.Name)
			continue
		}
		extractBindingsFromRole(crb.Subjects, crb.RoleRef.Name, "Cluster", role.Rules, &results, &summary)
	}

	// 🔹 2. Scan RoleBindings (if --namespace is specified)
	if namespace != "" {
		// Enhancement: Check if namespace exists
		_, err := clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("%s Warning: Namespace '%s' does not exist. No RoleBindings will be found.\n", emoji("⚠️"), namespace)
		} else if err != nil {
			fmt.Printf("%s Failed to validate namespace existence: %v\n", emoji("❌"), err)
			return results, summary
		}

		// Proceed to list RoleBindings in the specified namespace
		rbs, err := clientset.RbacV1().RoleBindings(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			fmt.Printf("%s Failed to list RoleBindings in %s: %v\n", emoji("❌"), namespace, err)
			return results, summary
		}
		fmt.Printf("%s Found %d RoleBindings in namespace: %s\n", emoji("🔍"), len(rbs.Items), namespace)

		for _, rb := range rbs.Items {
			role, err := clientset.RbacV1().Roles(namespace).Get(ctx, rb.RoleRef.Name, metav1.GetOptions{})
			if err != nil {
				fmt.Printf("%s Skipping RoleBinding %s: Role %s not found in namespace %s\n", emoji("⚠️"), rb.Name, rb.RoleRef.Name, namespace)
				continue
			}
			extractBindingsFromRole(rb.Subjects, rb.RoleRef.Name, namespace, role.Rules, &results, &summary)
		}
	}

	return results, summary
}

func extractBindingsFromRole(
	subjects []rbacv1.Subject,
	roleName string,
	scope string,
	rules []rbacv1.PolicyRule,
	out *[]AccessBinding,
	summary *Summary,
) {
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
	switch binding.RiskLevel {
	case "HIGH":
		if contains(binding.Verbs, "*") {
			summary.WildcardVerbs++
		}
		if binding.Role == "cluster-admin" {
			summary.ClusterAdminBindings++
		}
	case "MEDIUM":
		if contains(binding.Resources, "secrets") {
			summary.SecretsAccess++
		}
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
