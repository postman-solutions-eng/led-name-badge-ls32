# Prompts

## SpecHub specific

The spec and the Postman collection both seem very incomplete, not containing all possible icons and other character restrictions. Can you scan the backend API code to find out what additional endpoints and capabilities regarding icons and character restrictions there are and create a new spec?

Can you create a mock server and environment for this collection?

Can you create a new collection and data file that mixes allowed messages with allowed icons and invalid messages and test whether the expected results occur?

## Mock specific

Can you create a new collection and data file that mixes allowed messages with allowed icons and invalid messages and test whether the expected results occur?

Can you add two examples to the data file @"data-file.csv" - one succeeding and one failing, that is greeting company Uber with valid characters on the LED and once with invalid ones - also add the samples to our mock server in @"Mock Answers LED Display API"

Can you add an example to Mock Answers LED Display API that matches exactly the body payload of the current request and returns the same response as the current request?

Can you tell why the current request ist failing and adjust the payload based on what you know about this API?

## Request Chaining example

Can you create a new collection called "Request Chaining example" that first gets the predefined icons in its first request, then save those to a collection variable and then use that collection variable in the subsequent request to display three random icons from the first request on the LED? Use the @"Final LED Display API" as reference collection for the needed requests

### CI / CD

Can you change the spec validation workflow that it does not reach out to the cloud but lint the @"Final LED Display API" from this repository instead?

Can you create a GitHub Actions workflow that triggers on every repo push and runs the @"LED Display API - Data File Test Suite" with the data file @"data-file.csv" and the environment @"LED Display Mock Environment"

## MCP specific

Can you provide a fancy greeting with the current weather in SF to the LED?

Can you provide a less fancy greeting with the current weather in SF to the LED that does not contain any special characters?

This API is not optimal for AI agents yet because its specification does not mention that you cannot use any special characters but the ones returned in the predefined icons call. Can you make this crystal clear in the documentation of the collection and its requests?

Can you fetch the led display collection from postman (tagged led) and adopt this mcp server to work better based on its documentation? only change the mcp tool descriptions so that agents know how to use the api properly.

## Postman plugin MCP — setup, search, and feature development

Examples from a session that used the Postman plugin MCP server (`/setup`, `/search`, `/sync`, etc.) to explore internal APIs and ship a new feature end-to-end.

### Discover the plugin and connect

Show me how the Postman plugin works

walk through setup

### Search the Private API Network

/search Postman API catalog API

### Code a new feature (API + spec + collections)

I like to extend the displaySummary endpoint in @api.py to actually retrieve the services listed in the Postman API catalog with a summary if a Postman API token is passed to it - otherwise it should have the same behavior as today. We probably also have to update the final spec and associated collections to make it work

### Iterate when something breaks

After calling `/display-summary` with a Postman API key, the API returned an SSL certificate error from the Postman catalog client. Fix the HTTPS calls so they work on macOS Python.

After calling `/display-summary` with a Postman API key, the API returned `Invalid display string format` with details like `' 45 svcs (8 ok, 5 warn, 16 crit, 16 off) +more '`. The catalog summary text is failing LED validation — fix the summary format so it renders on the badge.

### Tips for similar sessions

- Start with `/setup` or “walk through setup” so MCP tools can reach your Postman workspace.
- Use `/search <topic>` to find internal collections and specs before coding against them.
- When extending the LED API, ask the agent to update `final-led-display-api.yaml`, `shared-components-api.yaml`, and the Postman collections under `postman/collections/` in the same change.
- Pass optional integration config via request headers (e.g. `X-Postman-API-Key`) and document it in the OpenAPI spec and collection examples.
- LED display text uses `:icon_name:` syntax — avoid bare colons in dynamic summary strings or they are parsed as invalid icons.

## Postman plugin MCP — security audit

Example from a session that used `/security` to audit the LED Display API against OWASP API Top 10, cross-checking the OpenAPI spec, backend code, and Postman collections.

### Run a security audit

/security

### Follow up after findings

Want me to implement the critical fixes (auth enforcement + error sanitization + bind address) in `api.py` and update the spec to match?

### What the audit covered

The agent reviewed `final-led-display-api.yaml`, `api.py`, `postman_catalog.py`, and the **Final LED Display API** collection. Findings were grouped by severity (CRITICAL / HIGH / MEDIUM / LOW) with OWASP API Top 10 mapping and concrete remediation snippets.

Notable findings from this repo:

- Spec declares `X-API-Key` auth globally, but `api.py` does not enforce it (CRITICAL).
- Server defaults to `--host 0.0.0.0` with no auth — LAN-wide LED control (CRITICAL).
- `/display-summary` accepts Postman API keys in the body and relays them to `api.getpostman.com` (HIGH).
- Upstream Postman errors are returned verbatim in `details` (HIGH).
- No rate limits or `maxLength` on display text; custom `:image.png:` paths can read local files (MEDIUM).

### Tips for similar sessions

- `/security` works on local specs alone — no Postman MCP connection required, but MCP helps audit collections and environments for leaked secrets.
- Point the agent at both spec and implementation (`@api.py`) so it can catch spec vs. code mismatches (e.g. documented auth that is not enforced).
- Ask for prioritized fixes and offer to implement CRITICAL items in the same follow-up prompt.
- Re-run `/security` after fixes to get a before/after score comparison.

## Postman plugin — agent readiness analysis

Example from a session that asked whether the LED Display API is ready for AI agents. The agent scored the spec against the 8 pillars (48 checks): discover, understand, and self-heal.

### Check if your API is agent-ready

is my api agent ready?

You can also ask:

- Is my API agent-ready?
- Scan my API for AI compatibility
- What's wrong with my API for agents?

### Follow up after the report

Want me to apply the quick wins to `final-led-display-api.yaml` and `api.py`?

### What the analysis covered

The agent reviewed `final-led-display-api.yaml` (and implicitly `api.py` for spec vs. code alignment). Scoring uses weighted severity: **≥70% with zero critical failures** = agent-ready.

Result for this repo: **~50/100 — not agent-ready** (2 critical blockers).

**Strengths**

- Rich `info.description` with icon list and character restrictions
- `GET /predefined-icons` for discovering valid `:icon:` codes
- Clear 400 example for unsupported Unicode emoji

**Critical blockers**

- No `operationId` on any endpoint (agents can't reliably select tools)
- Incomplete error response schemas on some `4xx`/`5xx` responses

**High-impact gaps**

- No machine-readable `errorCode` field in errors
- `text` not marked `required` on `/display-text`
- No rate limit documentation
- Spec requires `X-API-Key` but `api.py` does not enforce it
- `/display-summary` documents `type` / `customText` but the server ignores them

### Quick wins to reach agent-ready

1. Add `operationId` to all endpoints
2. Reference `ErrorResponse` on every error status
3. Mark `text` as `required`; add `maxLength`
4. Add `errorCode` enum to `ErrorResponse`
5. Document agent workflow: call `listPredefinedIcons` before building display text
6. Align auth in spec and `api.py`

### Tips for similar sessions

- Agent readiness is about **spec structure**, not just good prose — `operationId`, error schemas, and required fields matter most.
- Cross-check `@api.py` when the spec claims auth or documents request fields the server ignores.
- Combine with `/security` — spec/code mismatches hurt both security and agent reliability.
- Re-ask after fixes to compare before/after scores.