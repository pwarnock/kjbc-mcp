---
name: lookup-hebrew
description: Look up a Hebrew word in Strong's Concordance
---

# Look Up Hebrew Word

You are helping the user look up a Hebrew word in Strong's Concordance.

## Instructions

1. Use the `mcp__kjbc__word` tool to search for the English word in Hebrew
   - Set `word` parameter to the word the user provided (or ask if not provided)
   - Set `language` parameter to "hebrew"

2. The tool will return Strong's numbers (e.g., H430, H7965) that match the word

3. For each Strong's number returned:
   - Use the `mcp__kjbc__entry` tool to get the full definition
   - Set `entry_number` to the Strong's number (e.g., "H430")
   - Set `language` to "hebrew"

4. Present the results in a clear, readable format showing:
   - The Strong's number
   - The Hebrew word (transliteration)
   - The pronunciation guide
   - The definition/meaning
   - Any additional context from the entry

## Example Interaction

User: `/lookup-hebrew peace`

You should:
1. Search for "peace" in Hebrew
2. Get definitions for all returned Strong's numbers
3. Display results like:

**H7965 - shalom** (shaw-lome')
Peace, completeness, soundness, welfare, safety...

(continue for all matches)

## Tips

- If the user doesn't provide a word, ask them what word they want to look up
- Present multiple results clearly so the user can see the different nuances
- Keep definitions concise but complete
