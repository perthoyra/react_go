package cipher

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testKey = "QWERTYUIOPASDFGHJKLZXCVBNM"

func TestSubstitute(t *testing.T) {
	encrypted, err := substitute("Hello, World! 123", testKey, false)
	if err != nil {
		t.Fatal(err)
	}
	if want := "Itssg, Vgksr! 123"; encrypted != want {
		t.Fatalf("encrypt = %q, want %q", encrypted, want)
	}

	decrypted, err := substitute(encrypted, testKey, true)
	if err != nil {
		t.Fatal(err)
	}
	if want := "Hello, World! 123"; decrypted != want {
		t.Fatalf("decrypt = %q, want %q", decrypted, want)
	}
}

func TestSubstituteRejectsInvalidKey(t *testing.T) {
	for _, key := range []string{"", "ABC", "ABCDEFGHIJKLMNOPQRSTUVWXYA", "ABCDEFGHIJKLMNOPQRSTUVWXY!"} {
		if _, err := substitute("text", key, false); err == nil {
			t.Errorf("substitute accepted invalid key %q", key)
		}
	}
}

func TestHandler(t *testing.T) {
	handler := Handler(false)
	body, err := json.Marshal(cipherRequest{Text: "Hello", Key: testKey})
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/encrypt", bytes.NewReader(body))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d: %s", response.Code, http.StatusOK, response.Body)
	}
	var result cipherResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if result.Result != "Itssg" {
		t.Fatalf("result = %q, want %q", result.Result, "Itssg")
	}
}

func TestHandlerRejectsMethod(t *testing.T) {
	response := httptest.NewRecorder()
	Handler(false).ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/encrypt", nil))
	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusMethodNotAllowed)
	}
}
