# Substitution Cipher Web App

A small React frontend for the substitution cipher API. It lets you encrypt or decrypt text using a 26-letter key that contains every letter from `A` to `Z` exactly once.

## Features

- Toggle between encrypt and decrypt modes
- Enter custom text and key
- Submit requests to the Go API
- See the transformed result immediately
- Validation for invalid key inputs from the backend

## Prerequisites

- Node.js and npm
- Go

## Run the app

1. Start the backend service from the `restapi` folder:

   ```bash
   cd ../restapi
   go run .
   ```

   The API listens on `http://localhost:8080`.

2. In the `reactweb` folder, install dependencies:

   ```bash
   cd ../reactweb
   npm install
   ```

3. Start the Vite dev server:

   ```bash
   npm run dev
   ```

4. Open the app in your browser at:

   ```text
   http://localhost:5173
   ```

The frontend uses Vite's dev-server proxy to forward `/encrypt` and `/decrypt` requests to the Go API on port `8080`, so the browser sees the requests as same-origin calls.

## Build for production

```bash
npm run build
```

This produces a production bundle in the `dist/` directory.
