# Example CLI Workflows

These examples showcase how to use Permiflow for audits, CI pipelines, developer workflows, and more.

---

## ✅ Basic RBAC Scan

```bash
permiflow scan
```

- Uses default kubeconfig
- Outputs: `report.md`, `report.csv`, `report.json`, `metadata.json`

---

## 📝 Audit with Custom Output Path

```bash
permiflow scan \
  --kubeconfig ~/.kube/prod-config \
  --out-dir ./audit/prod \
  --prefix prod-audit
```

---

## 🔍 Detect Drift Between Two Snapshots

```bash
permiflow diff \
  --before ./audit/scan1/report.json \
  --after ./audit/scan2/report.json \
  --out-dir ./diffs
```

---

## ❌ CI/CD Gate: Fail if HIGH Risk Added

```bash
permiflow diff \
  --before ./baseline.json \
  --after ./latest.json \
  --fail-on high
```

Exit code will be `1` if any new HIGH risk bindings are introduced.

---

## 🚀 Automate Pre/Post Deployment Checks

```bash
permiflow scan \
  --out-dir ./pre-deploy --prefix report

# Apply your deployment
kubectl apply -f my-app.yaml

permiflow scan \
  --out-dir ./post-deploy --prefix report

permiflow diff \
  --before ./pre-deploy/report.json \
  --after ./post-deploy/report.json \
  --fail-on medium
```

---

## 👁️ View Scan History

```bash
permiflow history
```

Displays recent scans with IDs and timestamps.

---

## 💡 Export Minimal Role for Dev Bots

```bash
permiflow generate-role \
  --allow-verbs get,list,watch \
  --exclude-resources secrets,pods/exec \
  --name dev-reader \
  --out dev-reader.yaml
```

---

## 🚀 Fast Test on a Kind Cluster

```bash
permiflow scan \
  --kubeconfig ~/.kube/kind-config-kind \
  --dry-run
```

---

For in-depth CLI reference, continue to [configuration.md](./configuration.md).
