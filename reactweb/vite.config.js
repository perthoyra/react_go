import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// The Go REST API (../restapi) listens on :8080 and sends no CORS headers,
// so the dev server proxies its endpoints to keep requests same-origin.
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/encrypt': 'http://localhost:8080',
      '/decrypt': 'http://localhost:8080',
    },
  },
})
