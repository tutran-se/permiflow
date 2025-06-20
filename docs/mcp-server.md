# MCP Server

Permiflow includes an optional MCP (Model Context Protocol) server that exposes its scanning capabilities via a standard interface — useful for IDE plugins, external tools, or programmatic integrations.

---

## 🚀 What Is MCP?

The Model Context Protocol is a JSON-RPC based protocol used by tools like [Cursor IDE](https://www.cursor.so) to interact with local context providers.

Permiflow’s MCP implementation exposes the `scan_rbac` tool, making it easy to:

- Ask AI tools for RBAC summaries
- Trigger scans via API calls
- Integrate into rich developer workflows

---

## 💡 Key Features

- Supports **HTTP** and **STDIO** transports
- Works with any `kubeconfig`
- Same risk classification logic as the CLI
- CLI flag or env var configuration

---

## 🚧 Starting the Server

### STDIO (Recommended for IDEs like Cursor)

```bash
permiflow mcp --transport stdio
```

### HTTP (For tools that speak JSON-RPC over HTTP)

```bash
permiflow mcp --transport http --http-port 8080
```

---

## ⚙️ CLI Flags

| Flag           | Description               | Env Var            | Default           |
| -------------- | ------------------------- | ------------------ | ----------------- |
| `--transport`  | `http` or `stdio`         | `MCP_TRANSPORT`    | `stdio`           |
| `--http-port`  | Port to bind (HTTP only)  | -                  | `8080`            |
| `--debug`      | Enable debug logging      | `MCP_DEBUG`        | `false`           |
| `--kubeconfig` | Kubeconfig path           | `KUBECONFIG`       | `~/.kube/config`  |
| `--context`    | Kubernetes context to use | `MCP_KUBE_CONTEXT` | (current context) |

---

## 🔢 Tool: `scan_rbac`

This is the core tool exposed over MCP. It accepts the following parameters:

| Param        | Type      | Description                            |
| ------------ | --------- | -------------------------------------- |
| `kubeconfig` | string    | Optional. Path to kubeconfig           |
| `context`    | string    | Optional. Kube context to use          |
| `format`     | string    | `json` (default) or `summary`          |
| `namespaces` | string\[] | Optional. List of namespaces to filter |

### Example Response (`format: json`)

```json
{
  "findings": [ ... ],
  "summary": {
    "clusterAdminBindings": 1,
    "wildcardVerbs": 2,
    ...
  }
}
```

### Example Response (`format: summary`)

```
=== RBAC Security Scan Summary ===
Cluster admin bindings: 2
Wildcard verbs: 3
Secrets access: 8
...
```

---

## 🔍 Testing the Server

### STDIO Test

```bash
echo '{"jsonrpc":"2.0","method":"tools/list","id":1}' | ./permiflow mcp --transport stdio
```

### HTTP Test

```bash
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"scan_rbac"},"id":1}'
```

---

## 🧰 Cursor IDE Integration

Add this to your `.cursor/mcp.json` config:

### STDIO

```json
{
  "mcpServers": {
    "permiflow": {
      "command": "/path/to/permiflow",
      "args": ["mcp", "--transport", "stdio"]
    }
  }
}
```

### HTTP

Start the server first:

```bash
./permiflow mcp --transport http --http-port 8080
```

Then configure:

```json
{
  "mcpServers": {
    "permiflow": {
      "url": "http://localhost:8080/mcp"
    }
  }
}
```

---

Next up: [cursor-integration.md](./cursor-integration.md) for hands-on examples.
