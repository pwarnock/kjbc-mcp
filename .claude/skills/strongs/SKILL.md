---
name: strongs
description: Get the full definition for a Strong's number (e.g., G26 or H430)
---

# /strongs <number>

Get the full definition for a specific Strong's number.

1. Parse the number - G prefix = Greek, H prefix = Hebrew
2. Use `mcp__kjbc__entry` with the number and appropriate language
3. Display: Strong's number, original word, pronunciation, complete definition

Example: `/strongs G26` returns the full entry for agape.

If no prefix is provided, ask whether they mean Greek (G) or Hebrew (H).
