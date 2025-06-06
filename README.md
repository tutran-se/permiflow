# Permiflow

Permiflow is a simple read-only CLI tool that scans your Kubernetes RBAC configuration and outputs clean Markdown and CSV reports. Useful for SOC 2 prep, access reviews, and cluster audits.

## Features

- 🔍 Flag risky RBAC bindings (e.g. `cluster-admin`, `secrets`, `wildcards`)
- 📄 Generate Markdown for review meetings
- 📊 Export CSVs for auditors

## Usage

```bash
go run .
```

## Sample Output

See example/report.md and example/report.csv

## Requirements

- Go 1.21+

- Kubernetes access (~/.kube/config)

## License

MIT

## Folder Structure

```plaintext
permiflow/
├── main.go
├── client.go
├── scanner.go
├── risk.go
├── render_markdown.go
├── render_csv.go
├── go.mod
├── README.md
├── example/
│   ├── report.md
│   └── report.csv
```
