# `resources` Command

The `resources` command lets you inspect Kubernetes API resources available to your cluster. It’s a useful discovery tool when designing RBAC policies or troubleshooting excluded permissions.

---

## 🔧 Basic Usage

```bash
permiflow resources
```

This will print a grouped and sorted list of Kubernetes resources by API group and version.

---

## 🚧 What It Does

- Queries the Kubernetes Discovery API
- Groups resources by `GroupVersion`
- Lists their names, whether they are namespaced, and supported verbs
- Supports filtering by group, version, and namespaced-only mode

---

## 🏠 Key Flags

| Flag                | Type     | Description                                                             |
| ------------------- | -------- | ----------------------------------------------------------------------- |
| `--group`           | `string` | Filter by API group (e.g. `apps`, `batch`, `rbac.authorization.k8s.io`) |
| `--version`         | `string` | Filter by API version (e.g. `v1`, `v1beta1`)                            |
| `--namespaced-only` | `bool`   | Only show namespaced resources                                          |
| `--json`            | `bool`   | Output results as JSON instead of human-readable table                  |

---

## 🔢 Example: List All Namespaced Resources

```bash
permiflow resources --namespaced-only
```

---

## 🔢 Example: Filter by API Group

```bash
permiflow resources --group apps
```

---

## 🔢 Example: JSON Output

```bash
permiflow resources --json > k8s-resources.json
```

---

## 🚀 Typical Output (Default View)

```
📦 Kubernetes API Resources by GroupVersion:
------------------------------------------------

=== Group: apps/v1 ===
  • deployments                 [scope: namespaced]  verbs: [create, delete, get, list, patch, update, watch]
  • daemonsets                  [scope: namespaced]  verbs: [create, delete, get, list, patch, update, watch]

=== Group: rbac.authorization.k8s.io/v1 ===
  • roles                       [scope: namespaced]  verbs: [create, delete, get, list, update, watch]
  • clusterroles                [scope: cluster-wide] verbs: [get, list, watch]
```

---

## 📄 Use Cases

- Discover which resources are available in a restricted or custom API server
- Use as a reference when building secure roles with `generate-role`
- Understand whether a resource is namespaced or cluster-wide

---

For full command references, see [configuration.md](./configuration.md).
