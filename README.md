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

- A clean, readable Markdown report (ideal for auditors, reviewers, and GRC)
- A machine-parsable CSV/JSON export for analysis or GitOps flows
- **Drift detection between scans** for audits or CI/CD pipelines
- Peace of mind that your cluster was never touched or mutated

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

- Scans `ClusterRoleBindings` and `RoleBindings`
- Expands roles into rules (verbs + resources)
- Classifies risks: `HIGH`, `MEDIUM`, `LOW`
- Exports reports in **Markdown** (with ToC), **CSV**, and **JSON** formats
- Provides a **scan history** for traceability and future comparisons
- Performs **RBAC drift detection** between any two scans
- Flags dangerous permissions like:
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

### Install Permiflow

```bash
go install github.com/tutran-se/permiflow@latest
```

### Scan Your Cluster

```bash
# Basic scan
permiflow scan

# Dry run: no files written, no scan history recorded.
permiflow scan --dry-run

# Full scan
permiflow scan \
  --kubeconfig ~/.kube/config \
  --out-dir ./audit \
  --prefix report
```

### Compare scans (drift detection)

```bash
permiflow diff \
  --before ./audit/scan1/report.json \
  --after ./audit/scan2/report.json \
  --out-dir ./diffs
```

### Fail in CI if high-risk access is introduced

```bash
permiflow diff \
  --before ./baseline/report.json \
  --after ./latest/report.json \
  --fail-on high
```

---

## 🧾 Metadata & Scan History

Permiflow tracks each scan for traceability and future comparison.

### Each scan generates:

A `metadata.json` file containing:

- **Scan ID**
- **Timestamp**
- **Cluster context**
- **Output file names**
- **Risk summary**

### Global history is stored at:

`.permiflow/history.json`

Use the built-in CLI command to view your scan history:

```
> permiflow history
Scan History
--------------------------------------------
Scan ID:    2025-06-12T08-58-17Z--94c7f21f
Path:       audit/2025-06-12T08-58-17Z--94c7f21f
Context:    (default)
Timestamp:  2025-06-12T08:58:17Z

Scan ID:    2025-06-12T09-11-50Z--52c65f0d
Path:       audit/2025-06-12T09-11-50Z--52c65f0d
Context:    (default)
Timestamp:  2025-06-12T09:11:50Z

Scan ID:    2025-06-12T09-20-45Z--8fb8fdf8
Path:       examples/2025-06-12T09-20-45Z--8fb8fdf8
Context:    (default)
Timestamp:  2025-06-12T09:20:45Z
```

---

## 🔍 Example CLI Output

```
> permiflow scan --out-dir audit
Permiflow: Scanning RBAC...
Found 51 ClusterRoleBindings
Scanning RoleBindings in 5 namespaces
Found 0 RoleBindings in namespace: default
Found 0 RoleBindings in namespace: dev
Found 2 RoleBindings in namespace: uat
Found 9 RoleBindings in namespace: stagging
Found 0 RoleBindings in namespace: prod
Scan completed in 403.99ms
Metadata written to: audit/2025-06-12T20-03-59Z--63d5db96/metadata.json
Markdown written to: audit/2025-06-12T20-03-59Z--63d5db96/report.md
CSV written to: audit/2025-06-12T20-03-59Z--63d5db96/report.csv
JSON written to: audit/2025-06-12T20-03-59Z--63d5db96/report.json
Scan history updated: .permiflow/history.json
Report complete. 240 bindings scanned.
Summary:
- Found 2 cluster-admin binding(s)
- Found 3 wildcard verb usage(s)
- Found 8 subject(s) with secrets access
- Found 0 privilege escalation(s)
- Found 16 exec access(es)
- Found 16 config read secrets access(es)
```

```
> permiflow diff --before audit/report-before.json --after audit/report-after.json
RBAC Diff Summary
------------------
+ user-alice gained get access to configmaps in prod (via Role: config-reader) [MEDIUM]
- user-temp lost exec access to pods/exec in prod (via Role: debug-access) [HIGH]

Added: 1, Removed: 1, Changed: 0

Diff written to audit/
Files: diff.md, diff.json
```

---

## 🏁 Supported CLI Flags

### `scan` command

| Flag           | Type     | Description                                                                                     |
| -------------- | -------- | ----------------------------------------------------------------------------------------------- |
| `--kubeconfig` | `string` | Path to kubeconfig file (default: `~/.kube/config`)                                             |
| `--dry-run`    | `bool`   | Run scan without writing output files                                                           |
| `--out-dir`    | `string` | Output directory for reports                                                                    |
| `--prefix`     | `string` | Base name for output files (without extension). Example: 'audit' → audit.md (default: 'report') |

### `diff` command

| Flag        | Type     | Description                                                                     |
| ----------- | -------- | ------------------------------------------------------------------------------- |
| `--before`  | `string` | Path to baseline JSON report                                                    |
| `--after`   | `string` | Path to newer/current JSON report                                               |
| `--out-dir` | `string` | Output directory for diff exports (diff.md, diff.json)                          |
| `--fail-on` | `string` | Fail the command if `HIGH`, `MEDIUM`, or `LOW` risk is newly introduced in diff |

---

## 🌐 MCP Server

Permiflow includes an MCP (Model Context Protocol) server that exposes RBAC scanning capabilities through a standard interface, making it easy to integrate with AI tools like Cursor and other MCP-compatible clients.

### Features

- **Multiple Transport Protocols**: Supports both HTTP and STDIO transports
- **Standardized Interface**: Implements the Model Context Protocol specification
- **RBAC Scanning**: Exposes the same powerful RBAC scanning capabilities as the CLI
- **Automatic Kubeconfig Detection**: Uses default kubeconfig path or environment variables
- **Graceful Shutdown**: Cleanly handles shutdown signals and resource cleanup

### Getting Started

1. **Build the binary**:

```bash
go build -o permiflow .
```

2. **Run the MCP server with HTTP transport**:

```bash
./permiflow mcp --transport http --http-port 8080
```

3. **Run the MCP server with STDIO transport**:

```bash
./permiflow mcp --transport stdio
```

### Configuration

The MCP server can be configured using command-line flags or environment variables:

| Flag           | Type   | Description                               | Environment Variable | Default          |
| -------------- | ------ | ----------------------------------------- | -------------------- | ---------------- |
| `--transport`  | string | Transport type (http or stdio)            | `MCP_TRANSPORT`      | `stdio`          |
| `--http-port`  | int    | HTTP port (only used with http transport) | -                    | `8080`           |
| `--debug`      | bool   | Enable debug logging                      | `MCP_DEBUG`          | `false`          |
| `--kubeconfig` | string | Path to kubeconfig file                   | `KUBECONFIG`         | `~/.kube/config` |
| `--context`    | string | Kubernetes context to use                 | `MCP_KUBE_CONTEXT`   | Current context  |

### Cursor IDE Integration

Permiflow MCP server works seamlessly with Cursor IDE. Add one of these configurations to your Cursor MCP settings:

#### STDIO Transport (Recommended)

```json
{
  "mcpServers": {
    "permiflow": {
      "command": "/path/to/your/permiflow",
      "args": ["mcp", "--transport", "stdio"]
    }
  }
}
```

#### HTTP Transport

First, start the server:

```bash
./permiflow mcp --transport http --http-port 8080
```

Then configure Cursor:

```json
{
  "mcpServers": {
    "permiflow": {
      "url": "http://localhost:8080/mcp"
    }
  }
}
```

### Available Tools

#### scan_rbac

Scans Kubernetes RBAC configurations and identifies potential security risks.

**Parameters:**

- `kubeconfig` (optional): Path to kubeconfig file (defaults to `~/.kube/config`)
- `context` (optional): Kubernetes context to use
- `format` (optional): Output format - `json` for detailed findings, `summary` for overview only

**Example Usage in Cursor:**

Once configured, you can ask Cursor:

- "Scan my Kubernetes RBAC for security issues"
- "Check for privilege escalation risks in my cluster"
- "Show me a summary of RBAC security findings"

**Example JSON Response:**

```json
{
  "findings": [
    {
      "subject": "system:admin",
      "subjectKind": "User",
      "role": "cluster-admin",
      "namespace": "",
      "verbs": ["*"],
      "resources": ["*"],
      "scope": "Cluster",
      "riskLevel": "HIGH",
      "reason": "Wildcard verb or resource detected"
    }
  ],
  "summary": {
    "clusterAdminBindings": 2,
    "wildcardVerbs": 3,
    "secretsAccess": 8,
    "privilegeEscalation": 0,
    "execAccess": 16,
    "configReadSecrets": 16
  }
}
```

### Testing the MCP Server

#### STDIO Transport

```bash
echo '{"jsonrpc":"2.0","method":"tools/list","params":{},"id":1}' | ./permiflow mcp --transport stdio
```

#### HTTP Transport

```bash
# Start server
./permiflow mcp --transport http --http-port 8080 --debug

# Test in another terminal
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/list","params":{},"id":1}'
```

### Integration Example

Here's an example of how to call the scan_rbac tool via JSON-RPC:

```bash
# Get available tools
echo '{"jsonrpc":"2.0","method":"tools/list","params":{},"id":1}' | ./permiflow mcp --transport stdio

# Scan RBAC with summary format
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"scan_rbac","arguments":{"format":"summary"}},"id":2}' | ./permiflow mcp --transport stdio

# Scan RBAC with JSON format
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"scan_rbac","arguments":{"format":"json"}},"id":3}' | ./permiflow mcp --transport stdio
```

---

## 📣 License & Acknowledgements

Permiflow is released under the MIT License.

Built with ❤️ for Kubernetes security practitioners.

---
