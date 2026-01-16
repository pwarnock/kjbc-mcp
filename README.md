# kjbc-mcp

[![npm](https://img.shields.io/npm/v/@pwarnock/kjbc-mcp)](https://www.npmjs.com/package/@pwarnock/kjbc-mcp)
[![License](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](LICENSE)

Strong's Concordance MCP server. Provides grounded Greek/Hebrew word lookups for AI assistants.

## Quick Start

### Any MCP-compatible agent

Add to your MCP configuration:

```json
{
  "mcpServers": {
    "kjbc": {
      "command": "npx",
      "args": ["-y", "@pwarnock/kjbc-mcp"]
    }
  }
}
```

### Claude Code

Install the plugin:

```
/plugin marketplace add pwarnock/pwarnock-cc-plugins
/plugin install kjbc
```

### Claude Desktop

Add to `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "kjbc": {
      "command": "npx",
      "args": ["-y", "@pwarnock/kjbc-mcp"]
    }
  }
}
```

## Tools

| Tool | Description |
|------|-------------|
| `word` | Find Strong's numbers for an English word |
| `entry` | Get definition for a Strong's number |
| `langs` | List supported languages (Greek, Hebrew) |

### word

Find Strong's entries for an English word.

```
word: "love", language: "greek"
→ G25, G26, G5360, G5361, G5362, G5363
```

### entry

Get the full definition for a Strong's number.

```
entry_number: "G26", language: "greek"
→ agape (ag-ah'-pay): love, affection, benevolence
```

## Alternative Installation

### Go developers

```bash
go install github.com/pwarnock/kjbc-mcp/cmd/kjbc-mcp@latest
```

### Manual binary

Download from [releases](https://github.com/pwarnock/kjbc-mcp/releases).

## Attribution

Wraps the [kjbc](https://github.com/jrswab/kjbc) library by [@jrswab](https://github.com/jrswab).

Strong's Concordance data is public domain.

## License

GPL-3.0
