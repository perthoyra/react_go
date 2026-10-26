# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

- Run the server (listens on `:8080`): `go run .`
- Run all tests: `go test ./...`
- Run a single test: `go test -run TestSubstitute ./internal/cipher`
- Format / vet: `gofmt -w . && go vet ./...`

Go 1.22, module `restapi`, standard library only (no third-party dependencies).

## Architecture

An HTTP service for a monoalphabetic substitution cipher. `main.go` only wires up routes; all logic lives in `internal/cipher`:

- `cipher.Handler(decrypt bool)` returns the handler for both routes: `/encrypt` uses `false`, `/decrypt` uses `true`. Request handling (POST-only, 1 MiB body limit, strict JSON decoding with `DisallowUnknownFields` and a check that the body contains exactly one JSON object) is shared, so changes to request handling apply to both endpoints.
- `substitute(text, key, decrypt)` holds the cipher logic and does not depend on HTTP. It upper-cases and validates the key (exactly 26 ASCII letters, each A–Z once), builds a 26-byte mapping, and inverts it for decryption. It preserves letter case and leaves everything except ASCII letters unchanged. It works byte by byte, so multi-byte UTF-8 passes through untouched.
- Key validation errors from `substitute` are returned to the client verbatim as `400` responses, so their messages are part of the API.
- All errors are JSON `{"error": "..."}` via `writeError`; successful responses are `{"result": "..."}`.

Tests in `internal/cipher/cipher_test.go` are in the same package, so they call `substitute` directly and exercise handlers through `httptest` without starting a server. The README example (`"Hello, World!"` with key `QWERTYUIOPASDFGHJKLZXCVBNM` → `"Itssg, Vgksr!"`) matches the test fixture.
