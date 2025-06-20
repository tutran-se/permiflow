# Risk Model

Permiflow uses a built-in heuristic to assign a risk level to each RBAC binding it analyzes. These classifications help you prioritize which permissions require review or remediation.

---

## ⚠️ Risk Levels

| Level  | Color     | Meaning                                                                           |
| ------ | --------- | --------------------------------------------------------------------------------- |
| HIGH   | 🔴 Red    | Dangerous permissions. Can lead to privilege escalation or remote code execution. |
| MEDIUM | 🔷 Orange | Sensitive access. Potential for lateral movement or data leakage.                 |
| LOW    | 🔵 Blue   | Low-risk or routine access. Still worth tracking for drift.                       |

---

## ⚖️ HIGH Risk Conditions

Bindings that match **any** of the following are classified as HIGH:

- `verbs: ["*"]` (wildcard)
- `resources: ["*"]` (wildcard)
- Access to `pods/exec` → Remote shell / command execution
- Access to `nodes` → Host-level access
- Verbs like `create`/`update` on:

  - `roles`
  - `rolebindings`
  - `clusterroles`

**Why it matters:** These permissions enable cluster-wide control, escalation, or access to node internals.

---

## 📊 MEDIUM Risk Conditions

- Access to `secrets`
- `get` on `configmaps`

**Why it matters:** Secrets and configmaps often contain sensitive data like tokens, API keys, or credentials.

---

## 📲 LOW Risk Conditions

- Scopes that are:

  - Read-only (`get`, `list`, `watch`)
  - Non-sensitive resources (e.g. `pods`, `services`, `deployments`)

**Why it matters:** These are operational permissions that are usually required, but still worth logging.

---

## 📊 Risk Summary Output

Every scan includes a summarized count of risky bindings:

```txt
Summary:
- Found 2 cluster-admin binding(s)
- Found 3 wildcard verb usage(s)
- Found 8 subject(s) with secrets access
- Found 0 privilege escalation(s)
- Found 16 exec access(es)
- Found 16 config read secrets access(es)
```

---

## 🧹 Use Cases

- Gate deployments if new HIGH risk bindings appear
- Highlight HIGH/MEDIUM roles in reviews
- Tune roles using `generate-role` with safer presets

---

Next: [mcp-server.md](./mcp-server.md) to expose scans over HTTP or STDIO for automation tools.
