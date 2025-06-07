package main

import (
	"context"

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
}

type Summary struct {
	ClusterAdminBindings int
	WildcardVerbs        int
	SecretsAccess        int
}

func ScanRBAC(clientset *kubernetes.Clientset, namespace string) ([]AccessBinding, Summary) {
	var results []AccessBinding
	var summary Summary

	crbs, _ := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	for _, crb := range crbs.Items {
		roleName := crb.RoleRef.Name
		for _, subject := range crb.Subjects {
			role, _ := clientset.RbacV1().ClusterRoles().Get(context.TODO(), roleName, metav1.GetOptions{})
			for _, rule := range role.Rules {
				binding := AccessBinding{
					Subject:   subject.Name,
					Role:      roleName,
					Namespace: subject.Namespace,
					Verbs:     rule.Verbs,
					Resources: rule.Resources,
					RiskLevel: ClassifyRisk(rule.Verbs, rule.Resources),
				}
				results = append(results, binding)

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
		}
	}

	return results, summary
}
