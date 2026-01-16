---
name: lookup-greek
description: Look up an English word in the Greek Strong's Concordance
---

# /lookup-greek <word>

Look up an English word in the Greek Strong's Concordance.

1. Use the `mcp__kjbc__word` tool with `language: "greek"` and the word provided
2. For each Strong's number returned, use `mcp__kjbc__entry` to get the full definition
3. Present results showing: Strong's number, Greek word, pronunciation, and definition

Example: `/lookup-greek love` returns agapao (G25), agape (G26), phileo (G5368), etc.
