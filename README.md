# 🚦 Permiflow

**Permiflow** is a focused, read-only CLI tool that scans Kubernetes RBAC bindings and generates structured, human-readable reports — perfect for security reviews, SOC 2 audits, and internal compliance snapshots.

> > 🚧 MVP: v0.1.x — minimal, offline-compatible, zero-mutation scanning

---

## 🔧 What It Does

- 📊 Scans `ClusterRoleBindings` (and `RoleBindings` in upcoming `v0.2.0`)
- 🔍 Expands roles into rules (verbs + resources)
- 🧠 Classifies risks: `HIGH`, `MEDIUM`, `LOW`
- 📄 Exports reports in **Markdown** and **CSV**
- ✅ Flags dangerous permissions like:
  - `cluster-admin`
  - Wildcard verbs (`*`)
  - Access to sensitive resources (e.g. `secrets`)

---

## 🛡️ Security-First by Design

- Read-only: **no writes to the cluster**
- Offline-compatible: no agents, no CRDs, no API writes
- Works with any `kubeconfig` file or cluster
- No external dependencies — just Go + your config

---

## 🚀 Quick Start

```bash
go install github.com/tutran-se/permiflow@latest

permiflow \
  --kubeconfig ~/.kube/config \
  --markdown \
  --csv \
  --out-dir ./audit \
  --plain
```

Requires Go 1.21+

After running, you'll see:

- `./audit/report.md`
- `./audit/report.csv`

---

## 📦 Output Structure

Permiflow generates two files by default:

| File         | Purpose                                   |
| ------------ | ----------------------------------------- |
| `report.md`  | Human-friendly access summary             |
| `report.csv` | Structured table of roles and permissions |

Customizable via `--out-dir` and `--prefix`.

## 🔍 Example CLI Output

```
🔍 Permiflow: Scanning RBAC...

📄 Markdown written to: ./audit/report.md
📊 CSV written to: ./audit/report.csv

📊 Summary:

- Found 1 cluster-admin binding(s)

- Found 3 wildcard verb usage(s)

- Found 2 subject(s) with secrets access

✅ Report complete. 14 bindings scanned.
```

## 🏁 Supported CLI Flags

| Flag           | Type     | Description                                         |
| -------------- | -------- | --------------------------------------------------- |
| `--kubeconfig` | `string` | Path to kubeconfig file (default: `~/.kube/config`) |
| `--namespace`  | `string` | Scan a specific namespace only (optional)           |
| `--markdown`   | `bool`   | Generate Markdown output (default: true)            |
| `--csv`        | `bool`   | Generate CSV output (default: true)                 |
| `--dry-run`    | `bool`   | Run scan without writing output files               |
| `--plain`      | `bool`   | Disable emojis in output                            |
| `--out-dir`    | `string` | Output directory for reports                        |
| `--prefix`     | `string` | Optional prefix for output filenames                |
| `--version`    | `bool`   | Show version and exit                               |

### 🧪 Emoji Toggle

Disable emojis using the `--plain` flag or:

```bash
export PERMIFLOW_NO_EMOJI=true
```

---

## 📣 License & Acknowledgements

MIT License.

---
