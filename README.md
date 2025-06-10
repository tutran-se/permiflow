# Permiflow

[![Release](https://github.com/tutran-se/permiflow/actions/workflows/release.yml/badge.svg)](https://github.com/tutran-se/permiflow/actions/workflows/release.yml)
[![Latest Version](https://img.shields.io/github/v/tag/tutran-se/permiflow?label=version&sort=semver)](https://github.com/tutran-se/permiflow/releases)
[![Homebrew](https://img.shields.io/badge/install-homebrew-brightgreen)](https://github.com/tutran-se/homebrew-tap)
[![Scoop](https://img.shields.io/badge/install-scoop-blue)](https://github.com/tutran-se/scoop-bucket)
[![Go Report Card](https://goreportcard.com/badge/github.com/tutran-se/permiflow)](https://goreportcard.com/report/github.com/tutran-se/permiflow)
[![Go Reference](https://pkg.go.dev/badge/github.com/tutran-se/permiflow.svg)](https://pkg.go.dev/github.com/tutran-se/permiflow)
[![License](https://img.shields.io/github/license/tutran-se/permiflow)](https://github.com/tutran-se/permiflow/blob/main/LICENSE)

# 🚦 Permiflow

**Permiflow** is a zero-mutation CLI tool that scans Kubernetes RBAC bindings and generates structured, human-readable reports — perfect for security reviews, SOC 2 audits, and internal compliance snapshots.

---

## 🔧 What It Does

- 📊 Scans `ClusterRoleBindings` and `RoleBindings`
- 🔍 Expands roles into rules (verbs + resources)
- 🧠 Classifies risks: `HIGH`, `MEDIUM`, `LOW`
- 📄 Exports reports in **Markdown** (with ToC) and **CSV**
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
📦 Scanning cluster-wide bindings
📍 Scan scope: full cluster (all ClusterRoleBindings + all RoleBindings across namespaces)
🔍 Found 51 ClusterRoleBindings
📦 Scanning RoleBindings in 5 namespaces
🔍 Found 0 RoleBindings in namespace: default
🔍 Found 0 RoleBindings in namespace: dev
🔍 Found 2 RoleBindings in namespace: uat
🔍 Found 9 RoleBindings in namespace: prod
🔍 Found 0 RoleBindings in namespace: demo
⏱ Scan completed in 405.73ms
📄 Markdown written to: report.md
📊 CSV written to: report.csv
✅ Report complete. 240 bindings scanned.
📊 Summary:
   - Found 2 cluster-admin binding(s)
   - Found 3 wildcard verb usage(s)
   - Found 8 subject(s) with secrets access
```

## 🏁 Scan Modes

| Command                          | Behavior                                               |
| -------------------------------- | ------------------------------------------------------ |
| `permiflow scan`                 | Scans entire cluster (all namespaces, all roles)       |
| `permiflow scan --namespace xyz` | Scans only permissions that affect the `xyz` namespace |

## 🏁 Supported CLI Flags

| Flag           | Type     | Description                                                                                     |
| -------------- | -------- | ----------------------------------------------------------------------------------------------- |
| `--kubeconfig` | `string` | Path to kubeconfig file (default: `~/.kube/config`)                                             |
| `--namespace`  | `string` | Scan a specific namespace only (optional)                                                       |
| `--markdown`   | `bool`   | Generate Markdown output (default: true; use --markdown=false to disable)                       |
| `--csv`        | `bool`   | Generate CSV output (default: true; use --csv=false to disable)                                 |
| `--dry-run`    | `bool`   | Run scan without writing output files                                                           |
| `--plain`      | `bool`   | Disable emojis in output                                                                        |
| `--out-dir`    | `string` | Output directory for reports                                                                    |
| `--prefix`     | `string` | Base name for output files (without extension). Example: 'audit' → audit.md (default: 'report') |
| `--version`    | `bool`   | Show version and exit                                                                           |

### 🧪 Emoji Toggle

Disable emojis using the `--plain` flag or:

```bash
export PERMIFLOW_NO_EMOJI=true
```

---

## 📣 License & Acknowledgements

MIT License.

---
