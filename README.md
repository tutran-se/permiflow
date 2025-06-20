# Permiflow

[![Release](https://github.com/tutran-se/permiflow/actions/workflows/release.yml/badge.svg)](https://github.com/tutran-se/permiflow/actions/workflows/release.yml)
[![Latest Version](https://img.shields.io/github/v/tag/tutran-se/permiflow?label=version&sort=semver)](https://github.com/tutran-se/permiflow/releases)
[![Homebrew](https://img.shields.io/badge/install-homebrew-brightgreen)](https://github.com/tutran-se/homebrew-tap)
[![Scoop](https://img.shields.io/badge/install-scoop-blue)](https://github.com/tutran-se/scoop-bucket)
[![Go Report Card](https://goreportcard.com/badge/github.com/tutran-se/permiflow)](https://goreportcard.com/report/github.com/tutran-se/permiflow)
[![Go Reference](https://pkg.go.dev/badge/github.com/tutran-se/permiflow.svg)](https://pkg.go.dev/github.com/tutran-se/permiflow)
[![License](https://img.shields.io/github/license/tutran-se/permiflow)](https://github.com/tutran-se/permiflow/blob/main/LICENSE)

---

## 🚀 What is Permiflow?

**Permiflow** is a zero-mutation CLI tool that scans Kubernetes RBAC bindings and generates structured, human-readable reports. Built for security teams, platform engineers, and auditors.

- No CRDs. No agents. No API writes.
- Markdown, CSV, and JSON output
- Risk-aware diff detection between scans
- Safe for production clusters

---

## 📄 Documentation

All documentation has been moved into the [`docs/`](./docs) directory:

| Guide                                              | Description                                       |
| -------------------------------------------------- | ------------------------------------------------- |
| [Getting Started](./docs/getting-started.md)       | Install, scan, and compare in minutes             |
| [Concepts](./docs/concepts.md)                     | Understand RBAC, access bindings, and risk levels |
| [Scan Command](./docs/scan-command.md)             | How to scan your cluster and save results         |
| [Diff Command](./docs/diff-command.md)             | Compare two scans and detect permission drift     |
| [Generate Role](./docs/generate-role-command.md)   | Safely build secure Role or ClusterRole YAML      |
| [Resources Command](./docs/resources-command.md)   | Explore Kubernetes API resources and verbs        |
| [History & Metadata](./docs/history.md)            | Track and audit scans over time                   |
| [Report Formats](./docs/report-formats.md)         | Markdown, JSON, CSV explained                     |
| [Risk Model](./docs/risk-model.md)                 | How Permiflow classifies HIGH, MEDIUM, LOW        |
| [MCP Server](./docs/mcp-server.md)                 | Expose RBAC scanning via JSON-RPC for automation  |
| [Cursor Integration](./docs/cursor-integration.md) | Use Permiflow in Cursor IDE                       |
| [Examples](./docs/examples.md)                     | Common CLI workflows and automation patterns      |
| [Configuration Reference](./docs/configuration.md) | All flags and environment variables               |
| [FAQ](./docs/faq.md)                               | Common questions answered                         |

---

## 🛠️ Install

```bash
go install github.com/tutran-se/permiflow@latest
```

For Homebrew and Scoop options, see [Getting Started](./docs/getting-started.md).

---

## 🌐 License

MIT © [Tutran SE](https://github.com/tutran-se). Built with care for Kubernetes security engineers.
