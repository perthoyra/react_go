// Calls POST /encrypt or POST /decrypt on the restapi service.
// Returns the transformed text, or throws with the API's error message.
export async function transform(mode, text, key) {
  const response = await fetch(`/${mode}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text, key }),
  })

  let body
  try {
    body = await response.json()
  } catch {
    throw new Error(`Unexpected response from server (HTTP ${response.status})`)
  }
  if (!response.ok) {
    throw new Error(body.error ?? `HTTP ${response.status}`)
  }
  return body.result
}
