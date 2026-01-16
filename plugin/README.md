# KJBC Plugin

Strong's Concordance MCP for Claude Code. Provides grounded Greek/Hebrew word lookups.

## Installation

Install via the pwarnock marketplace:

```
/plugin marketplace add pwarnock/pwarnock-cc-plugins
/plugin install kjbc
```

Or add manually to your Claude Code settings:

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

## What It Does

Gives Claude access to Strong's Concordance data:
- Greek/Hebrew word definitions
- Strong's number lookups
- Verse references

## Example

> "What's the Greek word for 'love' in John 3:16?"

Claude can look up Strong's G25 (agapao) with the full definition, pronunciation, and usage.

## Requirements

- Node.js 16+ (for npx)
- Internet connection (first run downloads the binary)
