package permiflow_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tutran-se/permiflow/internal/permiflow"
)

func TestGenerateRBACRoleToStdout_BasicClusterRole(t *testing.T) {
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs:      []string{"get", "list"},
		ExcludedResources: []string{"secrets"},
		RoleName:          "test-clusterrole",
		RoleKind:          "ClusterRole",
	}

	var buf bytes.Buffer
	err := testencodeRBACRole(opts, &buf)
	assert.NoError(t, err)
	out := buf.String()

	assert.Contains(t, out, "kind: ClusterRole")
	assert.Contains(t, out, "name: test-clusterrole")
	assert.NotContains(t, out, "secrets")
	assert.Contains(t, out, "- get")
	assert.Contains(t, out, "- list")
}

func TestGenerateRBACRoleToStdout_BasicRole(t *testing.T) {
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs:      []string{"get", "watch"},
		ExcludedResources: []string{"nodes", "secrets"},
		RoleName:          "test-role",
		RoleKind:          "Role",
		Namespace:         "dev",
	}

	var buf bytes.Buffer
	err := testencodeRBACRole(opts, &buf)
	assert.NoError(t, err)
	out := buf.String()

	assert.Contains(t, out, "kind: Role")
	assert.Contains(t, out, "namespace: dev")
	assert.Contains(t, out, "name: test-role")
	assert.Contains(t, out, "- get")
	assert.Contains(t, out, "- watch")
	assert.NotContains(t, out, "nodes")
	assert.NotContains(t, out, "secrets")
}

func TestGenerateRBACRoleToStdout_ErrorOnMissingNamespace(t *testing.T) {
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs:      []string{"get"},
		ExcludedResources: []string{},
		RoleName:          "bad-role",
		RoleKind:          "Role",
		Namespace:         "",
	}

	var buf bytes.Buffer
	err := testencodeRBACRole(opts, &buf)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "namespace is required")
}

func TestGenerateRBACRole_FileOutput(t *testing.T) {
	path := "test-role.yaml"
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs:      []string{"create", "delete"},
		ExcludedResources: []string{"pods/exec"},
		RoleName:          "file-output",
		RoleKind:          "ClusterRole",
		OutputPath:        path,
	}

	err := permiflow.GenerateRBACRole(opts)
	assert.NoError(t, err)

	data, err := os.ReadFile(path)
	assert.NoError(t, err)
	content := string(data)
	assert.Contains(t, content, "ClusterRole")
	assert.Contains(t, content, "create")
	assert.NotContains(t, content, "pods/exec")

	_ = os.Remove(path)
}

// Helper to test encoding logic
func testencodeRBACRole(opts permiflow.GenerateRBACRoleOptions, buf *bytes.Buffer) error {
	// swap stdout serializer with mock buffer
	type encoder interface {
		Encode(obj interface{}, w *bytes.Buffer) error
	}
	return permiflow.GenerateRBACRoleToStdoutWithWriter(opts, buf)
}

func TestGenerateRBACRole_InvalidKindDefaultsToClusterRole(t *testing.T) {
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs:      []string{"get"},
		ExcludedResources: []string{"secrets"},
		RoleName:          "default-kind",
		RoleKind:          "UnknownKind",
	}

	var buf bytes.Buffer
	err := permiflow.GenerateRBACRoleToStdoutWithWriter(opts, &buf)
	assert.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "kind: ClusterRole")
	assert.Contains(t, out, "name: default-kind")
}

func TestSafeClusterAdminProfileLogic(t *testing.T) {
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs:      []string{"get", "list", "watch", "create", "update"},
		ExcludedResources: []string{"secrets", "pods/exec", "nodes", "rolebindings", "clusterroles"},
		RoleName:          "safe-cluster-admin",
		RoleKind:          "ClusterRole",
	}

	var buf bytes.Buffer
	err := permiflow.GenerateRBACRoleToStdoutWithWriter(opts, &buf)
	assert.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "safe-cluster-admin")
	assert.NotContains(t, out, "secrets")
	assert.NotContains(t, out, "pods/exec")
}

func TestGenerateRBACRole_NamespacedScopeExcludesClusterResources(t *testing.T) {
	// Simulate what --scope namespaced would do
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs: []string{"get"},
		ExcludedResources: []string{
			"nodes", "clusterroles", "clusterrolebindings", "persistentvolumes",
			"mutatingwebhookconfigurations", "validatingwebhookconfigurations",
		},
		RoleName: "namespaced-only",
		RoleKind: "ClusterRole",
	}

	var buf bytes.Buffer
	err := permiflow.GenerateRBACRoleToStdoutWithWriter(opts, &buf)
	assert.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "ClusterRole")
	assert.NotContains(t, out, "nodes")
	assert.NotContains(t, out, "clusterrolebindings")
	assert.Contains(t, out, "pods")
	assert.Contains(t, out, "configmaps")
}

func TestGenerateRBACRole_DuplicateExcludesHandled(t *testing.T) {
	opts := permiflow.GenerateRBACRoleOptions{
		AllowedVerbs:      []string{"get"},
		ExcludedResources: []string{"secrets", "secrets", "pods/exec"},
		RoleName:          "dedup-check",
		RoleKind:          "ClusterRole",
	}

	var buf bytes.Buffer
	err := permiflow.GenerateRBACRoleToStdoutWithWriter(opts, &buf)
	assert.NoError(t, err)

	out := buf.String()
	assert.NotContains(t, out, "secrets")
	assert.NotContains(t, out, "pods/exec")
}
