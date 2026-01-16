---
name: lookup-hebrew
description: Look up an English word in the Hebrew Strong's Concordance
---

# /lookup-hebrew <word>

Look up an English word in the Hebrew Strong's Concordance.

1. Use the `mcp__kjbc__word` tool with `language: "hebrew"` and the word provided
2. For each Strong's number returned, use `mcp__kjbc__entry` to get the full definition
3. Present results showing: Strong's number, Hebrew word, pronunciation, and definition

Example: `/lookup-hebrew god` returns elohim (H430), el (H410), eloah (H433), etc.
