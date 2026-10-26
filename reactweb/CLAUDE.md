# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

- Install dependencies: `npm install`
- Dev server: `npm run dev` (serves the page at http://localhost:5173)
- Production build: `npm run build` (output in `dist/`)
- Preview the build: `npm run preview` (serves the page at http://localhost:4173)

There are no tests or linter configured.

The page needs the Go backend in the sibling directory `../restapi` running on `:8080` (`cd ../restapi && go run .`).

## Architecture

A single-page Vite + React 18 (plain JSX, no TypeScript) frontend for the substitution-cipher REST API in `../restapi`.

- `src/api.js` — `transform(mode, text, key)` POSTs `{text, key}` to `/encrypt` or `/decrypt` and returns `result`. On a non-2xx response it throws an `Error` carrying the API's `{"error": "..."}` message, which the UI shows as-is.
- `src/App.jsx` — all UI state: text, key (prefilled with the README example key `QWERTYUIOPASDFGHJKLZXCVBNM`), encrypt/decrypt mode, and a result label that shows either the result or the error.

The API requires a 26-letter key containing each A–Z exactly once and rejects unknown JSON fields, so the request body must stay exactly `{text, key}`.

### Backend connectivity

The Go service sends no CORS headers. Requests use relative URLs (`/encrypt`, `/decrypt`) and rely on the proxy in `vite.config.js` forwarding them to `http://localhost:8080`. This only works under `npm run dev` or `npm run preview`. Serving `dist/` from anywhere else needs either CORS added to the Go service or a reverse proxy in front of both.
