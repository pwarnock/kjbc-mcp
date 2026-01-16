# KJBC Skill for Claude Code

A Claude Code skill that provides convenient commands for looking up Greek and Hebrew words using Strong's Concordance.

## Installation

```bash
/skill install pwarnock/kjbc-mcp
```

## Commands

### `/lookup-greek <word>`

Look up an English word in the Greek Strong's Concordance.

**Example:**
```
/lookup-greek love
```

Returns all Greek Strong's numbers and definitions for "love" (e.g., agapao, agape, phileo).

### `/lookup-hebrew <word>`

Look up an English word in the Hebrew Strong's Concordance.

**Example:**
```
/lookup-hebrew god
```

Returns all Hebrew Strong's numbers and definitions for "god" (e.g., elohim, el, eloah).

### `/strongs <number>`

Get the full definition for a specific Strong's number.

**Example:**
```
/strongs G26
```

Returns the complete definition for Strong's number G26 (agape).

Accepts both Greek (G###) and Hebrew (H###) numbers.

## Use Cases

- Bible study and original language research
- Understanding nuances between different Greek/Hebrew words translated to the same English word
- Looking up specific Strong's references mentioned in study materials
- Exploring word meanings in their original context

## Prerequisites

The KJBC MCP server must be installed and configured. The skill will use the MCP tools to perform lookups.

See the [main README](../README.md) for MCP server installation instructions.
