# `scan` Command

The `scan` command is the core of Permiflow. It collects and analyzes all RoleBindings and ClusterRoleBindings in your cluster, producing a structured permissions report.

---

## 🔧 Basic Usage

```bash
permiflow scan
```

This performs a full RBAC scan using the default kubeconfig and writes:

- `report.md` (human-readable)
- `report.csv` (machine-readable)
- `report.json` (full structured output)
- `metadata.json` (scan metadata)

---

## 🏠 Key Flags

| Flag               | Type     | Description                                                       |
| ------------------ | -------- | ----------------------------------------------------------------- |
| `--kubeconfig`     | `string` | Path to kubeconfig file (default: `~/.kube/config`)               |
| `--dry-run`        | `bool`   | Do not write output files or update history                       |
| `--out-dir`        | `string` | Directory to write output files into (default: current directory) |
| `--prefix`         | `string` | Base name for output files (e.g. `audit` -> `audit.md`)           |
| `--log-timestamps` | `bool`   | Enable timestamps in logs for debugging                           |

---

## 📅 Output Artifacts

### 1. `report.md`

A clean Markdown file that includes:

- Summary of risks
- Table of Contents by subject
- Detailed breakdown of each binding

### 2. `report.csv`

- Tabular view: subject, verbs, resources, risk
- Easy to analyze in Excel, Google Sheets, or scripts

### 3. `report.json`

- Complete structured representation
- Can be diffed or parsed in CI/CD pipelines

### 4. `metadata.json`

Includes:

- Timestamp
- Scan ID (with random suffix)
- Context name
- Number of bindings
- List of generated files

---

## ⌚ Performance

Even large clusters scan quickly — typically under 1 second. Permiflow prioritizes:

- Read-only API calls
- Lightweight analysis
- Zero cluster-side dependencies

---

## 💪 Real-World Examples

```bash
# Dry run with custom kubeconfig
permiflow scan \
  --dry-run \
  --kubeconfig ~/.kube/staging-config

# Save audit to a timestamped folder
permiflow scan \
  --out-dir ./audit \
  --prefix staging-audit
```

---

## 📄 Scan History

Each scan writes to `.permiflow/history.json` unless `--dry-run` is used.
View history with:

```bash
permiflow history
```

---

## ❇ Tip

You can use `scan` in automation or pre-deploy hooks to snapshot current permissions before changes are applied.

---

Continue to [diff-command.md](./diff-command.md) to learn how to compare scans.
