# Look Up Strong's Number

You are helping the user get the definition for a specific Strong's Concordance number.

## Instructions

1. Get the Strong's number from the user (if not already provided)
   - Should be in format G### for Greek or H### for Hebrew
   - Examples: G26, H430

2. Determine the language from the prefix:
   - G = Greek
   - H = Hebrew

3. Use the `mcp__kjbc__entry` tool to get the definition
   - Set `entry_number` to the Strong's number (e.g., "G26")
   - Set `language` to "greek" or "hebrew" based on the prefix

4. Present the full definition in a clear format showing:
   - The Strong's number
   - The original language word (transliteration)
   - The pronunciation guide
   - The complete definition/meaning
   - Any additional context from the entry

## Example Interaction

User: `/strongs G26`

You should:
1. Recognize this is Greek (G prefix)
2. Look up entry G26 in Greek
3. Display the result like:

**G26 - agape** (ag-ah'-pay)

Love, affection, benevolence, good will, esteem, a love feast. The word refers to the love of God and Christ, as well as the love Christians should have for one another and for God.

## Tips

- If the user provides just a number without prefix, ask whether they mean Greek (G) or Hebrew (H)
- Handle both uppercase and lowercase prefixes (g26 = G26)
- If the entry is not found, inform the user clearly and suggest they verify the number
