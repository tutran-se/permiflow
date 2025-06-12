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
- Generates a **human-readable summary** of key findings
- Provides a **scan history** for traceability and future comparisons
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

```bash
go install github.com/tutran-se/permiflow@latest

# Short version
permiflow scan

# Dry run: no files written, no scan history recorded.
permiflow scan --dry-run

# Full version
permiflow scan \
  --kubeconfig ~/.kube/config \
  --out-dir ./audit \
  --prefix report \
```

Requires Go 1.21+

After running, you'll see a **timestamped output folder** like:

```
./audit/2025-06-13T08-17-01Z--d3d57c28/
├── report.md
├── report.csv
├── report.json
├── metadata.json
```

- Each scan gets a unique **Scan ID** like `2025-06-13T08-17-01Z--d3d57c28`
- A `metadata.json` file stores scan time, summary, and output context

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

Scan ID:    2025-06-12T09-21-21Z--7518e75f
Path:       examples/2025-06-12T09-21-21Z--7518e75f
Context:    (default)
Timestamp:  2025-06-12T09:21:21Z

Scan ID:    2025-06-12T19-39-37Z--d3d57c28
Path:       audit/2025-06-12T19-39-37Z--d3d57c28
Context:    (default)
Timestamp:  2025-06-12T19:39:37Z
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

## 🏁 Supported CLI Flags

| Flag           | Type     | Description                                                                                     |
| -------------- | -------- | ----------------------------------------------------------------------------------------------- |
| `--kubeconfig` | `string` | Path to kubeconfig file (default: `~/.kube/config`)                                             |
| `--dry-run`    | `bool`   | Run scan without writing output files                                                           |
| `--out-dir`    | `string` | Output directory for reports                                                                    |
| `--prefix`     | `string` | Base name for output files (without extension). Example: 'audit' → audit.md (default: 'report') |

---

## 🌐 MCP Server

Permiflow includes an MCP (Model Context Protocol) server that exposes RBAC scanning capabilities through a standard interface, making it easy to integrate with other tools and services.

### Features

- **Multiple Transport Protocols**: Supports both HTTP and STDIO transports
- **Standardized Interface**: Implements the Model Context Protocol specification
- **RBAC Scanning**: Exposes the same powerful RBAC scanning capabilities as the CLI
- **Graceful Shutdown**: Cleanly handles shutdown signals and resource cleanup

### Getting Started

1. **Build the MCP server**:

```bash
go build -o bin/mcp-server ./cmd/mcp-server
```

2. **Run the MCP server with HTTP transport**:

```bash
./bin/mcp-server --transport http --http-port 8080
```

3. **Run the MCP server with STDIO transport**:

```bash
./bin/mcp-server --transport stdio
```

### Configuration

The MCP server can be configured using command-line flags or environment variables:

| Flag | Type | Description | Environment Variable | Default |
|------|------|-------------|----------------------|---------|
| `--transport` | string | Transport type (http or stdio) | `MCP_TRANSPORT` | `stdio` |
| `--http-port` | int | HTTP port (only used with http transport) | `MCP_HTTP_PORT` | `8080` |
| `--debug` | bool | Enable debug logging | `MCP_DEBUG` | `false` |
| `--kubeconfig` | string | Path to kubeconfig file | `KUBECONFIG` | `~/.kube/config` |
| `--context` | string | Kubernetes context to use | `KUBE_CONTEXT` | Current context |

### API Documentation

The MCP server exposes the following tools:

#### scan_rbac

Scans Kubernetes RBAC rules and generates a report.

**Input Schema:**

```json
{
  "output_format": "json" | "markdown" | "csv",
  "namespaces": ["string"]
}
```

**Example Request:**

```json
{
  "output_format": "json",
  "namespaces": ["default", "kube-system"]
}
```

**Example Response:**

```json
{
  "report": "...",
  "findings": [
    {
      "namespace": "default",
      "resource": "pods",
      "name": "pod-reader",
      "rules": [
        {
          "verbs": ["get", "list", "watch"],
          "resources": ["pods"],
          "api_groups": [""]
        }
      ],
      "risk_level": "LOW"
    }
  ]
}
```

### Integration Example

Here's an example of how to use the MCP server from a Go application:

```go
package main

import (
	"context"
	"fmt"
	"log"

	mcp "github.com/mark3labs/mcp-go/mcp"
)

func main() {
	// Create a new MCP client
	client, err := mcp.NewClient("http://localhost:8080")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Call the scan_rbac tool
	result, err := client.CallTool(context.Background(), "scan_rbac", map[string]interface{}{
		"output_format": "json",
		"namespaces":    []string{"default"},
	})
	if err != nil {
		log.Fatalf("Failed to call scan_rbac: %v", err)
	}

	// Process the result
	fmt.Printf("Scan result: %+v\n", result)
}
```

---

## 📣 License & Acknowledgements

Permiflow is released under the MIT License.

Built with ❤️ for Kubernetes security practitioners.

---
