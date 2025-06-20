# Cursor IDE Integration

Permiflow’s MCP server works seamlessly with [Cursor IDE](https://www.cursor.so), enabling secure and intelligent RBAC analysis right inside your coding environment.

---

## 🧑‍💻 Why Integrate?

- Ask Cursor to "scan Kubernetes RBAC for risks"
- Get inline summaries of cluster permissions
- Reduce context switching during security reviews

---

## ⚙️ Setup Instructions

### 1. Build Permiflow (if not already installed)

```bash
go build -o permiflow .
```

### 2. Add to `.cursor/mcp.json`

#### Option A: STDIO Transport (Recommended)

```json
{
  "mcpServers": {
    "permiflow": {
      "command": "/absolute/path/to/permiflow",
      "args": ["mcp", "--transport", "stdio"]
    }
  }
}
```

#### Option B: HTTP Transport

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

## 🔢 Supported Tools

### `scan_rbac`

Scans Kubernetes RBAC and returns either a summary or full list of risky bindings.

**Parameters:**

- `kubeconfig` (optional)
- `context` (optional)
- `format`: `summary` or `json`

---

## 💡 Example Prompts in Cursor

You can say:

- "Scan my Kubernetes RBAC for privilege escalation risks"
- "What service accounts have exec access in the staging cluster?"
- "Show me all users with secrets access"

Cursor sends a structured JSON-RPC call to Permiflow MCP behind the scenes.

---

## 🚀 Example Terminal Tests

### STDIO

```bash
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"scan_rbac","arguments":{"format":"summary"}},"id":1}' | ./permiflow mcp --transport stdio
```

### HTTP

```bash
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"scan_rbac","arguments":{"format":"summary"}},"id":1}'
```

---

Continue to [examples.md](./examples.md) for more real-world CLI patterns and workflows.
