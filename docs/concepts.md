# Core Concepts in Permiflow

Understanding the foundational elements of Kubernetes RBAC and how Permiflow interprets them is key to making the most of your scans and reports.

---

## 🔑 RBAC 101

Kubernetes Role-Based Access Control (RBAC) defines **who** (subjects) can perform **what actions** (verbs) on **which resources**.

A complete RBAC permission consists of:

- **Subject**: A User, Group, or ServiceAccount
- **Role/ClusterRole**: Defines allowed actions
- **Binding**: Grants the role to the subject, optionally within a namespace

Permiflow inspects all `RoleBindings` and `ClusterRoleBindings`, expanding them into granular access rules.

---

## 🔐 AccessBinding (Permiflow Representation)

Permiflow converts raw RBAC bindings into structured `AccessBinding` entries with:

| Field         | Description                                     |
| ------------- | ----------------------------------------------- |
| `Subject`     | Identity (user, service account, etc.)          |
| `SubjectKind` | Type of subject: User, Group, ServiceAccount    |
| `Role`        | Name of Role or ClusterRole bound               |
| `Namespace`   | Namespace for namespaced bindings               |
| `Verbs`       | Allowed actions: get, list, watch, etc.         |
| `Resources`   | Affected Kubernetes resource types              |
| `Scope`       | `Cluster` or `Namespaced`                       |
| `RiskLevel`   | HIGH / MEDIUM / LOW based on rule analysis      |
| `Reason`      | Textual explanation for the risk classification |

---

## ⚠️ Risk Classification

Permiflow uses a built-in heuristic to classify each binding’s risk level:

### HIGH Risk

- Wildcard verbs or resources (`*`)
- Ability to create/update roles or bindings (privilege escalation)
- Access to `pods/exec` or `nodes`

### MEDIUM Risk

- Read access to sensitive resources like `secrets` or `configmaps`

### LOW Risk

- Read-only or scoped permissions on non-sensitive resources

Learn more in [risk-model.md](./risk-model.md).

---

## 📝 Output Types

Each scan produces multiple output formats:

- **Markdown (`.md`)**: Human-readable report, ideal for audits and GRC
- **CSV (`.csv`)**: Tabular format for spreadsheets and automation
- **JSON (`.json`)**: Full data for programmatic use or CI/CD
- **Metadata (`metadata.json`)**: Scan context, timestamp, and file listing

---

## ⏳ Scan History

All scan metadata is persisted in `.permiflow/history.json`. You can view scan history with:

```bash
permiflow history
```

Use this to track drift, audit trends, or restore previous states.

---

## 🚪 Read-Only Guarantee

Permiflow does **not mutate** your cluster. It only uses read permissions via your kubeconfig. That makes it safe for:

- Production audits
- CI/CD gates
- Compliance snapshots

---

## ✅ Summary

Permiflow turns opaque RBAC rules into structured, auditable insights. The core concepts of **bindings**, **roles**, **subjects**, and **risk levels** underpin all scan reports and drift comparisons.

Ready to run your first deep scan? See [scan-command.md](./scan-command.md).
