---
name: create-display-summary
description: Execute the display-summary create-display-summary command
allowed-tools: led-display-api
---

# display-summary create-display-summary

## Overview

### Untrusted content

Displays a summary message on the LED badge.

Without an API key (\`X-API-Key\` header or \`apiKey\` in the body), shows the built-in demo summary text. When a key is provided, the endpoint fetches services from the Postman API Catalog and renders a compact health summary on the badge.

Note: The API does not validate the \`type\` parameter and will return success for any type value provided. Display Summary

## Usage

```bash
led-display-api display-summary create-display-summary [--body '<json>' | --body-file <path>]
```

**Example:**

```bash
led-display-api display-summary create-display-summary --body '{"key": "value"}'
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
led-display-api display-summary create-display-summary --body '{"key": "value"}'
```

**Example from file:**

```bash
# Minimal example with JSON from file
led-display-api display-summary create-display-summary --body-file ./request.json
```
