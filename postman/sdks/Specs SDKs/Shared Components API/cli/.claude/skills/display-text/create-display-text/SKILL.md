---
name: create-display-text
description: Execute the display-text create-display-text command
allowed-tools: led-display-api
---

# display-text create-display-text

## Overview

### Untrusted content

Updates the text and visual content displayed on the LED name badge. Accepts text and icon codes in the format \`:icon_name:\`.

\<img src="https://content.pstmn.io/05b1ef2c-9fd3-4f0e-8b2f-76998fb3f6e5/aW1hZ2UucG5n" width="240"\>

## Supported Icons

The following icon codes can be embedded in text using the \`:icon_name:\` syntax:

- \`:ball:\` - Filled circle
- \`:happy:\` - Simple smiley face
- \`:happy2:\` - Larger smiley (2 columns wide)
- \`:heart:\` - Outline heart
- \`:HEART:\` - Filled heart
- \`:heart2:\` - Larger outline heart (2 columns wide)
- \`:HEART2:\` - Larger filled heart (2 columns wide)
- \`:fablab:\` - FabLab logo
- \`:bicycle:\` - Bicycle icon (3 columns wide)
- \`:bicycle_r:\` - Bicycle facing right (3 columns wide)
- \`:owncloud:\` - OwnCloud logo (3 columns wide)
- \`:octocat:\` - GitHub Octocat
- \`:smile:\` - Smile emoji
- \`:star:\` - Star icon
- \`:sun:\` - Sun icon

## Character Restrictions

**Supported characters:**

- Letters: A-Z, a-z
- Numbers: 0-9
- Special characters: \`^ !"$%&/()=?\` °}\\]\\[{@ \\~ \|\<\>,\;.:-_#'+_\\\`
- German umlauts: äöüÄÖÜß
- French/European accents: àäòöùüèéêëôöûîïÿçÀÅÄÉÈÊËÖÔÜÛÙŸŠ

**NOT supported:**

- Unicode emoji (e.g., 🌍) - will return a 400 error Display Text

## Usage

```bash
led-display-api display-text create-display-text [--body '<json>' | --body-file <path>]
```

**Example:**

```bash
led-display-api display-text create-display-text --body '{"key": "value"}'
```

## Request Body

Provide the request body using one of the following methods:

| Method      | Flag                 | Description                             |
| ----------- | -------------------- | --------------------------------------- |
| Inline JSON | `--body '<json>'`    | Pass JSON directly as a string argument |
| File path   | `--body-file <path>` | Read JSON content from a file           |

**Example inline:**

```bash
# Minimal example with inline JSON body
led-display-api display-text create-display-text --body '{"key": "value"}'
```

**Example from file:**

```bash
# Minimal example with JSON from file
led-display-api display-text create-display-text --body-file ./request.json
```
