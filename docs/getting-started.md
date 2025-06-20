# Getting Started with Permiflow

Permiflow is a zero-mutation CLI tool that scans Kubernetes RBAC bindings and generates structured, human-readable reports. This guide walks you through installation and basic usage.

---

## ✨ What You'll Get

- Clean, readable **Markdown reports** for audits
- **CSV and JSON** outputs for automation or GitOps
- **Drift detection** to compare permissions across scans
- Full audit trail with **scan history and metadata**

---

## ⚙️ Installation

### Via Go (Recommended)

```bash
go install github.com/tutran-se/permiflow@latest
```

This places the `permiflow` binary in your `$GOPATH/bin` or `$HOME/go/bin`.

### Homebrew (macOS/Linux)

```bash
brew install tutran-se/tap/permiflow
```

### Scoop (Windows)

```powershell
scoop bucket add tutran-se https://github.com/tutran-se/scoop-bucket.git
scoop install permiflow
```

---

## 🚀 First Scan

Scan your current Kubernetes cluster and generate a Markdown report:

```bash
permiflow scan
```

By default, it uses the kubeconfig at `~/.kube/config` and saves output in the current directory.

### Dry Run (No Output Files)

```bash
permiflow scan --dry-run
```

### Custom Output Directory and File Prefix

```bash
permiflow scan \
  --kubeconfig ~/.kube/config \
  --out-dir ./audit \
  --prefix report
```

---

## ⚖️ Detect Permission Drift

Compare two scans (e.g., before and after a deployment):

```bash
permiflow diff \
  --before ./audit/scan1/report.json \
  --after ./audit/scan2/report.json \
  --out-dir ./diffs
```

### Fail in CI/CD on High-Risk Changes

```bash
permiflow diff \
  --before ./baseline/report.json \
  --after ./latest/report.json \
  --fail-on high
```

---

## 📅 View Scan History

Permiflow tracks scan history in `.permiflow/history.json`. To view:

```bash
permiflow history
```

---

## 📄 Example Output Files

After a scan, you’ll find:

- `report.md` - Human-readable audit
- `report.csv` - Machine-parsable summary
- `report.json` - Full structured output
- `metadata.json` - Metadata describing the scan context

---

## 🚀 You're Ready

You now have full visibility into Kubernetes RBAC configurations without ever mutating your cluster.

Next up: [Understanding Concepts](./concepts.md)
