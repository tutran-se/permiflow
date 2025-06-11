package permiflow

import (
	"testing"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestScanRBAC_ClusterRoleBinding(t *testing.T) {
	client := fake.NewSimpleClientset(
		&rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{Name: "read-secrets"},
			Rules: []rbacv1.PolicyRule{
				{Verbs: []string{"get"}, Resources: []string{"secrets"}},
			},
		},
		&rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{Name: "crb-1"},
			Subjects: []rbacv1.Subject{
				{Kind: "User", Name: "alice"},
			},
			RoleRef: rbacv1.RoleRef{Kind: "ClusterRole", Name: "read-secrets"},
		},
	)

	results, summary := ScanRBAC(client)
	if len(results) != 1 {
		t.Fatalf("expected 1 binding, got %d", len(results))
	}
	if results[0].RiskLevel != "MEDIUM" {
		t.Errorf("expected risk MEDIUM, got %s", results[0].RiskLevel)
	}
	if summary.SecretsAccess != 1 {
		t.Errorf("expected 1 secrets access, got %d", summary.SecretsAccess)
	}
}

func TestScanRBAC_WildcardVerb(t *testing.T) {
	client := fake.NewSimpleClientset(
		&rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{Name: "admin"},
			Rules: []rbacv1.PolicyRule{
				{Verbs: []string{"*"}, Resources: []string{"pods"}},
			},
		},
		&rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{Name: "crb-admin"},
			Subjects: []rbacv1.Subject{
				{Kind: "ServiceAccount", Name: "svc"},
			},
			RoleRef: rbacv1.RoleRef{Kind: "ClusterRole", Name: "admin"},
		},
	)

	_, summary := ScanRBAC(client)
	if summary.WildcardVerbs != 1 {
		t.Errorf("expected 1 wildcard verb usage, got %d", summary.WildcardVerbs)
	}
}

func TestClassifyRisk(t *testing.T) {
	tests := []struct {
		verbs     []string
		resources []string
		expected  string
	}{
		{[]string{"*"}, []string{"pods"}, "HIGH"},
		{[]string{"get"}, []string{"*"}, "HIGH"},
		{[]string{"get"}, []string{"secrets"}, "MEDIUM"},
		{[]string{"get"}, []string{"pods"}, "LOW"},
	}

	for _, tt := range tests {
		result, _ := ClassifyRisk(tt.verbs, tt.resources)
		if result != tt.expected {
			t.Errorf("ClassifyRisk(%v, %v) = %v; want %v", tt.verbs, tt.resources, result, tt.expected)
		}
	}
}
