# Substitution Cipher REST API

A small Go HTTP service with two JSON endpoints:

- `POST /encrypt`
- `POST /decrypt`

Each endpoint accepts a 26-letter substitution key that contains every ASCII
letter from `A` to `Z` exactly once:

```json
{
  "text": "Hello, World!",
  "key": "QWERTYUIOPASDFGHJKLZXCVBNM"
}
```

The response contains the transformed text:

```json
{
  "result": "Itssg, Vgksr!"
}
```

The key maps plaintext `A`–`Z` to the corresponding key letters. Decryption
uses the inverse mapping. Letter case is preserved, and characters other than
ASCII letters are left unchanged.

Run the service with `go run .`; it listens on port `8080`. Run tests with
`go test ./...`.
