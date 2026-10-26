import { useState } from 'react'
import { transform } from './api.js'

const DEFAULT_KEY = 'QWERTYUIOPASDFGHJKLZXCVBNM'

export default function App() {
  const [text, setText] = useState('')
  const [key, setKey] = useState(DEFAULT_KEY)
  const [mode, setMode] = useState('encrypt')
  const [result, setResult] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSubmit(event) {
    event.preventDefault()
    setLoading(true)
    setError('')
    try {
      setResult(await transform(mode, text, key))
    } catch (err) {
      setResult('')
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <main>
      <h1>Substitution Cipher</h1>
      <form onSubmit={handleSubmit}>
        <div className="toggle" role="group" aria-label="Mode">
          {['encrypt', 'decrypt'].map((m) => (
            <button
              key={m}
              type="button"
              className={mode === m ? 'active' : ''}
              aria-pressed={mode === m}
              onClick={() => setMode(m)}
            >
              {m === 'encrypt' ? 'Encrypt' : 'Decrypt'}
            </button>
          ))}
        </div>

        <label htmlFor="text">Text</label>
        <textarea
          id="text"
          rows={4}
          value={text}
          onChange={(e) => setText(e.target.value)}
        />

        <label htmlFor="key">Key (26 letters, A–Z once each)</label>
        <input
          id="key"
          value={key}
          maxLength={26}
          spellCheck={false}
          onChange={(e) => setKey(e.target.value)}
        />

        <button type="submit" className="send" disabled={loading}>
          {loading ? 'Sending…' : 'Send'}
        </button>
      </form>

      <label className={`result ${error ? 'error' : ''}`} aria-live="polite">
        {error || result || 'Result will appear here'}
      </label>
    </main>
  )
}
