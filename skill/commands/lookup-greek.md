# Look Up Greek Word

You are helping the user look up a Greek word in Strong's Concordance.

## Instructions

1. Use the `mcp__kjbc__word` tool to search for the English word in Greek
   - Set `word` parameter to the word the user provided (or ask if not provided)
   - Set `language` parameter to "greek"

2. The tool will return Strong's numbers (e.g., G25, G26) that match the word

3. For each Strong's number returned:
   - Use the `mcp__kjbc__entry` tool to get the full definition
   - Set `entry_number` to the Strong's number (e.g., "G26")
   - Set `language` to "greek"

4. Present the results in a clear, readable format showing:
   - The Strong's number
   - The Greek word (transliteration)
   - The pronunciation guide
   - The definition/meaning
   - Any additional context from the entry

## Example Interaction

User: `/lookup-greek love`

You should:
1. Search for "love" in Greek
2. Get definitions for all returned Strong's numbers
3. Display results like:

**G25 - agapao** (ag-ap-ah'-o)
Love, affection, benevolence towards another...

**G26 - agape** (ag-ah'-pay)
Love feast, love, affection, good will...

(continue for all matches)

## Tips

- If the user doesn't provide a word, ask them what word they want to look up
- Present multiple results clearly so the user can see the different nuances
- Keep definitions concise but complete
