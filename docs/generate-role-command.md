# `generate-role` Command

The `generate-role` command helps you safely create Kubernetes `ClusterRole` or `Role` definitions. It excludes sensitive resources by default and is ideal for provisioning minimal, principle-of-least-privilege RBAC for bots, contractors, and staging environments.

---

## ⚙️ Basic Usage

```bash
permiflow generate-role \
  --allow-verbs get,list,watch \
  --exclude-resources secrets,pods/exec \
  --name dev-reader \
  --out dev-reader.yaml
```

---

## 🏠 Key Flags

| Flag                  | Type     | Description                                                             |
| --------------------- | -------- | ----------------------------------------------------------------------- |
| `--allow-verbs`       | `string` | Comma-separated verbs to allow (e.g. `get,list,watch`)                  |
| `--exclude-resources` | `string` | Comma-separated resources to exclude (e.g. `secrets,pods/exec`)         |
| `--scope`             | `string` | Optional. Set to `namespaced` to exclude cluster-wide resources         |
| `--out`               | `string` | Output YAML file path (e.g. `safe-role.yaml`)                           |
| `--name`              | `string` | Name of the Role or ClusterRole (default: `almost-admin`)               |
| `--kind`              | `string` | Either `Role` or `ClusterRole` (default: `ClusterRole`)                 |
| `--namespace`         | `string` | Required if kind is `Role`                                              |
| `--dry-run`           | `bool`   | Print the YAML to stdout instead of writing it to file                  |
| `--explain`           | `bool`   | Print a summary of inputs without generating YAML                       |
| `--profile`           | `string` | Optional. Use a preset profile (e.g. `safe-cluster-admin`, `read-only`) |

---

## 🔍 Preset Profiles

You can use built-in profiles for common use cases:

### `--profile safe-cluster-admin`

Grants broad but safe verbs across most resources:

```yaml
Verbs: get, list, watch, create, update
Excludes: secrets, pods/exec, nodes, rolebindings, clusterroles
```

### `--profile read-only`

Restrictive access suitable for external reviewers:

```yaml
Verbs: get, list, watch
Excludes: secrets, pods/exec, nodes, create, update, delete
```

---

## 🔢 Example: ClusterRole for Audit Bot

```bash
permiflow generate-role \
  --allow-verbs get,list \
  --exclude-resources secrets,nodes,pods/exec \
  --name audit-bot \
  --out audit-bot-role.yaml
```

---

## 🔢 Example: Namespaced Role (Must include --namespace)

```bash
permiflow generate-role \
  --kind Role \
  --namespace dev \
  --allow-verbs get,list \
  --exclude-resources secrets \
  --name dev-role \
  --out dev-role.yaml
```

---

## 🔢 Example: Dry Run with Explanation

```bash
permiflow generate-role \
  --allow-verbs get,list \
  --exclude-resources secrets \
  --dry-run --explain
```

Outputs a YAML preview and a summary of inputs.

---

## 👍 Best Practices

- Use `--scope namespaced` to automatically exclude dangerous cluster-wide resources
- Use `--dry-run` in CI/CD pipelines to verify roles before committing
- Store output YAML in Git to track role evolution over time

---

Continue to [configuration.md](./configuration.md) for full flag reference or [examples.md](./examples.md) to see how this integrates into real workflows.
