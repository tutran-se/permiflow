# Frequently Asked Questions (FAQ)

This document addresses the most common questions users have when using Permiflow.

---

## ❓ Does Permiflow modify my cluster?

**No.** Permiflow is strictly **read-only**. It uses your `kubeconfig` to scan RBAC resources but never creates, updates, or deletes anything.

---

## ❓ Do I need cluster-admin permissions to run it?

You only need permissions to **list** the following:

- ClusterRoles
- ClusterRoleBindings
- Roles (across namespaces)
- RoleBindings (across namespaces)

If your kubeconfig doesn't provide access to these, the scan may be incomplete.

---

## ❓ Can I use Permiflow in CI/CD pipelines?

**Yes.** It supports dry runs, file output, and exit codes based on risk level via `--fail-on`.

---

## ❓ What formats does it support?

- Markdown (`.md`)
- CSV (`.csv`)
- JSON (`.json`)
- Metadata (`metadata.json`)

All formats are generated in one run and can be stored, versioned, or diffed.

---

## ❓ Where is scan history stored?

Permiflow writes history to `.permiflow/history.json`. You can view it using:

```bash
permiflow history
```

---

## ❓ Can I exclude sensitive resources?

Yes. When generating roles with `generate-role`, use `--exclude-resources`.

Example:

```bash
permiflow generate-role \
  --allow-verbs get,list \
  --exclude-resources secrets,pods/exec
```

---

## ❓ How does the risk model work?

Bindings are evaluated against a heuristic:

- Wildcards or privilege escalation → **HIGH**
- Secrets access → **MEDIUM**
- Read-only access to non-sensitive resources → **LOW**

See [risk-model.md](./risk-model.md) for full details.

---

## ❓ Can I use it with Cursor IDE?

Yes. Permiflow includes an MCP server that integrates with Cursor using either STDIO or HTTP.

See [cursor-integration.md](./cursor-integration.md).

---

## ❓ Is it safe for production clusters?

**Absolutely.** Permiflow is safe for production usage as it only reads cluster state. It never mutates anything.

---

For deeper reference, see the [architecture.md](./architecture.md) or open a GitHub issue if your question isn’t covered.
