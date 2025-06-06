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

### 🔧 Optional: Disable Emoji

Some users prefer plain terminal output (e.g. for CI logs or `grep`).  
You can disable all emoji in CLI messages with:

```bash
PERMIFLOW_NO_EMOJI=true go run main.go
```

## Sample Output

See example/report.md and example/report.csv

## Requirements

- Go 1.21+

- Kubernetes access (~/.kube/config)

## License

MIT
