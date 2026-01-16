# @pwarnock/kjbc-mcp

Strong's Concordance MCP server for AI assistants. Provides grounded Greek/Hebrew word lookups from the King James Bible Concordance.

## Installation

```bash
npx @pwarnock/kjbc-mcp
```

Or add to your MCP configuration:

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

- **word** - Find Strong's entries for an English word
- **entry** - Get details for a Strong's number (definition, pronunciation)
- **langs** - List supported languages (Greek, Hebrew)

## Example

Ask your AI assistant:
> "What's the Greek word for 'love' in the New Testament?"

The assistant can look up Strong's G25 (agapao) and G26 (agape) with full definitions.

## License

GPL-3.0
