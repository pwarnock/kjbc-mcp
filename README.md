# kjbc-mcp

A minimal Model Context Protocol (MCP) server for the King James Bible Concordance (KJBC). It exposes a small, read-only tool surface over stdio and relies on the upstream `kjbc` library for data and lookups.

## Quick start

1) Install the binary (recommended)
- Download a release from `https://github.com/pwarnock/kjbc-mcp/releases` and place it on your PATH.

2) Configure your MCP host (example: Claude Desktop)

```json
{
  "mcpServers": {
    "kjbc": {
      "command": "kjbc-mcp",
      "args": []
    }
  }
}
```

## Build from source

```bash
go install github.com/pwarnock/kjbc-mcp/cmd/kjbc-mcp@latest
```

Note: this requires Go 1.24+ because the upstream `kjbc` module declares Go 1.24.

## Tools

- `word`: Find Strong's entries for an English word.
  - Inputs: `word` (string), `language` (greek|hebrew, default greek), `show_verses` (bool, default false)
- `entry`: Get details for a Strong's entry number.
  - Inputs: `entry_number` (string), `language` (greek|hebrew, default greek)
- `langs`: List supported languages.

## Attribution

This MCP server is a thin wrapper around the original `kjbc` tool by @jrswab:
- Source: `https://github.com/jrswab/kjbc`

Strong's Concordance data is public domain and sourced via `kjbc` (see that repository for full attribution).

## License

GPL-3.0. See `LICENSE`.
