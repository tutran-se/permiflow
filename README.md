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

## ❓ Why Permiflow?

Kubernetes RBAC is powerful — but opaque. Most tools either mutate live clusters, dump cryptic JSON, or require complex setups.

**Permiflow** was built to make **RBAC visibility dead simple**, especially for security-conscious teams. With a single command, you get:

- 📄 A clean, readable Markdown report (ideal for auditors, reviewers, and GRC)
- 📊 A machine-parsable CSV/JSON export for analysis or GitOps flows
- 🛡️ Peace of mind that your cluster was never touched or mutated

No CRDs. No agents. No surprises.

---

## 👤 Who Is It For?

Permiflow is made for:

- **Platform Engineers** maintaining secure, multi-tenant clusters
- **Security Engineers** conducting internal reviews or threat modeling
- **Compliance & GRC Teams** prepping for SOC 2, ISO 27001, or FedRAMP audits
- **SREs & DevOps Practitioners** who want clear, actionable permission insights
- Anyone who needs **RBAC clarity** — without modifying the cluster

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
  - Privilege escalation risks

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

# Short version
permiflow scan

# Dry run (no output files)
permiflow scan --dry-run

# Full version
permiflow scan \
  --kubeconfig ~/.kube/config \
  --markdown \
  --csv \
  --json \
  --out-dir ./audit \
  --prefix report \
  --plain
```

Requires Go 1.21+

After running, you'll see:

- `./audit/report.md`
- `./audit/report.csv`
- `./audit/report.json`

---

## 📦 Output Structure

Permiflow generates two files by default:

| File          | Purpose                                   |
| ------------- | ----------------------------------------- |
| `report.md`   | Human-friendly access summary             |
| `report.csv`  | Structured table of roles and permissions |
| `report.json` | Machine-readable JSON format              |

Customizable via `--out-dir` and `--prefix`.

## 🔍 Example CLI Output

```
🔍 Permiflow: Scanning RBAC...
🔍 Found 51 ClusterRoleBindings
📦 Scanning RoleBindings in 5 namespaces
🔍 Found 0 RoleBindings in namespace: default
🔍 Found 0 RoleBindings in namespace: dev
🔍 Found 2 RoleBindings in namespace: uat
🔍 Found 9 RoleBindings in namespace: stagging
🔍 Found 0 RoleBindings in namespace: prod
⏱ Scan completed in 410.46ms
📄 Markdown written to: examples/report.md
📊 CSV written to: examples/report.csv
📦 JSON written to: examples/report.json
✅ Report complete. 240 bindings scanned.
📊 Summary:
   - Found 2 cluster-admin binding(s)
   - Found 3 wildcard verb usage(s)
   - Found 8 subject(s) with secrets access
   - Found 0 privilege escalation(s)
   - Found 16 exec access(es)
   - Found 16 config read secrets access(es)
```

## 🏁 Supported CLI Flags

| Flag           | Type     | Description                                                                                     |
| -------------- | -------- | ----------------------------------------------------------------------------------------------- |
| `--kubeconfig` | `string` | Path to kubeconfig file (default: `~/.kube/config`)                                             |
| `--markdown`   | `bool`   | Generate Markdown output (default: true; use --markdown=false to disable)                       |
| `--csv`        | `bool`   | Generate CSV output (default: true; use --csv=false to disable)                                 |
| `--json`       | `bool`   | Generate JSON output (default: true; use --json=false to disable)                               |
| `--dry-run`    | `bool`   | Run scan without writing output files                                                           |
| `--plain`      | `bool`   | Disable emojis in output                                                                        |
| `--out-dir`    | `string` | Output directory for reports                                                                    |
| `--prefix`     | `string` | Base name for output files (without extension). Example: 'audit' → audit.md (default: 'report') |

### 🧪 Emoji Toggle

Disable emojis using the `--plain` flag or:

```bash
export PERMIFLOW_NO_EMOJI=true
```

---

## 📣 License & Acknowledgements

Permiflow is released under the MIT License.

Built with ❤️ for Kubernetes security practitioners.

---
