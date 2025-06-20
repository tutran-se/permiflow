# Configuration Reference

Permiflow is designed to be flexible and automation-friendly. You can configure it via CLI flags or environment variables.

---

## 🔧 CLI Flags

### Global

| Flag        | Type   | Description                            |
| ----------- | ------ | -------------------------------------- |
| `--dry-run` | `bool` | Skip writing files or updating history |

---

### `scan` Command

| Flag               | Type     | Description                                         |
| ------------------ | -------- | --------------------------------------------------- |
| `--kubeconfig`     | `string` | Path to kubeconfig file                             |
| `--out-dir`        | `string` | Directory to write output files                     |
| `--prefix`         | `string` | Prefix for output files (e.g. `audit` → `audit.md`) |
| `--log-timestamps` | `bool`   | Include timestamps in log output                    |

---

### `diff` Command

| Flag        | Type     | Description                                                                |
| ----------- | -------- | -------------------------------------------------------------------------- |
| `--before`  | `string` | Path to baseline JSON report                                               |
| `--after`   | `string` | Path to newer/current JSON report                                          |
| `--out-dir` | `string` | Directory to write diff output (Markdown and JSON)                         |
| `--fail-on` | `string` | Risk level threshold to trigger non-zero exit code (e.g. `high`, `medium`) |

---

### `generate-role` Command

| Flag                  | Type     | Description                                                       |
| --------------------- | -------- | ----------------------------------------------------------------- |
| `--allow-verbs`       | `string` | Comma-separated verbs to allow (e.g. `get,list,watch`)            |
| `--exclude-resources` | `string` | Comma-separated resources to exclude (e.g. `secrets,pods/exec`)   |
| `--scope`             | `string` | Optional. Set to `namespaced` to exclude cluster-wide resources   |
| `--out`               | `string` | Path to write the output YAML file                                |
| `--name`              | `string` | Name of the Role or ClusterRole                                   |
| `--kind`              | `string` | `Role` or `ClusterRole`                                           |
| `--namespace`         | `string` | Required if kind is `Role`                                        |
| `--dry-run`           | `bool`   | Print YAML to stdout instead of writing a file                    |
| `--explain`           | `bool`   | Print a summary without generating the file                       |
| `--profile`           | `string` | Optional. Use a preset profile: `safe-cluster-admin`, `read-only` |

---

### `resources` Command

| Flag                | Type     | Description                                                    |
| ------------------- | -------- | -------------------------------------------------------------- |
| `--group`           | `string` | Filter by API group (e.g. `apps`, `rbac.authorization.k8s.io`) |
| `--version`         | `string` | Filter by API version (e.g. `v1`)                              |
| `--namespaced-only` | `bool`   | Only show namespaced resources                                 |
| `--json`            | `bool`   | Output as JSON instead of formatted table                      |

---

### `history` Command

- No flags required.

---

### `version` Command

- No flags required.

---

### `completion` Command

| Argument   | Values                              | Description                       |
| ---------- | ----------------------------------- | --------------------------------- |
| Shell name | `bash`, `zsh`, `fish`, `powershell` | Generates shell completion script |

---

### `mcp` Command

| Flag           | Type     | Description                          |
| -------------- | -------- | ------------------------------------ |
| `--transport`  | `string` | `http` or `stdio`                    |
| `--http-port`  | `int`    | Port to bind if using HTTP transport |
| `--debug`      | `bool`   | Enable debug logging                 |
| `--kubeconfig` | `string` | Kubeconfig path                      |
| `--context`    | `string` | Kubernetes context to use            |

---

## 🌍 Environment Variables

| Env Var            | Used By | Description                           |
| ------------------ | ------- | ------------------------------------- |
| `KUBECONFIG`       | All     | Overrides kubeconfig path             |
| `MCP_TRANSPORT`    | `mcp`   | `http` or `stdio`                     |
| `MCP_DEBUG`        | `mcp`   | Set to `true` to enable debug logging |
| `MCP_KUBE_CONTEXT` | `mcp`   | Overrides Kubernetes context          |

---

For real-world usage examples, see [examples.md](./examples.md).
