# Scan History & Metadata

Permiflow automatically tracks each scan you perform to ensure traceability, support comparisons, and aid in compliance documentation.

---

## 🔍 What Gets Tracked

Each time you run `permiflow scan`, a **metadata file** is created, and an entry is added to a **global scan history log**.

This allows you to:

- Audit changes over time
- Track scan frequency
- Detect drift between known states

---

## 📅 Viewing History

```bash
permiflow history
```

This command reads `.permiflow/history.json` and prints a chronological list of previous scans.

**Example:**

```
Scan History
--------------------------------------------
Scan ID:    2025-06-12T08-58-17Z--94c7f21f
Path:       audit/2025-06-12T08-58-17Z--94c7f21f
Context:    (default)
Timestamp:  2025-06-12T08:58:17Z
```

---

## 📃 File: `metadata.json`

Each scan creates a file like `audit/2025-06-12T08-58-17Z--94c7f21f/metadata.json` containing:

| Field          | Description                                      |
| -------------- | ------------------------------------------------ |
| `scan_id`      | Unique ID with timestamp and suffix              |
| `timestamp`    | UTC time of scan                                 |
| `kubeconfig`   | Path used for cluster access                     |
| `context`      | Kube context (e.g., `prod-cluster`)              |
| `out_dir`      | Folder containing the outputs                    |
| `prefix`       | Base name for all output files                   |
| `num_bindings` | Number of RBAC bindings found                    |
| `summary`      | Object with counts of risks (e.g., wildcard use) |
| `output_files` | Markdown, CSV, and JSON filenames                |

---

## 📄 File: `.permiflow/history.json`

A central log of the **last 50 scans**, stored in your workspace root. It includes:

- Scan ID
- Path to output folder
- Timestamp
- Kube context used

This file is updated unless you run with `--dry-run`.

---

## ⚠️ Git Hygiene Tip

If you're storing reports in Git, consider **ignoring** `.permiflow/history.json` to prevent local state from polluting version control.

---

## ❇ Use Cases

- Pair with `diff` to always compare the last 2 scans
- Document scans as part of audit evidence
- Trigger alerts if no recent scan has occurred

---

Ready to explore output formats? See [report-formats.md](./report-formats.md).
