package cipher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type cipherRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type cipherResponse struct {
	Result string `json:"result"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func Handler(decrypt bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		var request cipherRequest
		if err := decoder.Decode(&request); err != nil {
			writeError(w, http.StatusBadRequest, "request body must contain valid JSON with text and key")
			return
		}
		if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
			writeError(w, http.StatusBadRequest, "request body must contain exactly one JSON object")
			return
		}

		result, err := substitute(request.Text, request.Key, decrypt)
		if err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cipherResponse{Result: result}); err != nil {
			log.Printf("writing response: %v", err)
		}
	}
}

func substitute(text, key string, decrypt bool) (string, error) {
	key = strings.ToUpper(key)
	if len(key) != 26 {
		return "", fmt.Errorf("key must contain exactly 26 ASCII letters")
	}

	var seen [26]bool
	for i := 0; i < len(key); i++ {
		if key[i] < 'A' || key[i] > 'Z' {
			return "", fmt.Errorf("key must contain only ASCII letters")
		}
		index := key[i] - 'A'
		if seen[index] {
			return "", fmt.Errorf("key must contain each letter A-Z exactly once")
		}
		seen[index] = true
	}

	var mapping [26]byte
	for i := range mapping {
		mapping[i] = key[i]
	}
	if decrypt {
		for i, mapped := range mapping {
			mapping[mapped-'A'] = byte('A' + i)
		}
	}

	output := []byte(text)
	for i, char := range output {
		switch {
		case char >= 'A' && char <= 'Z':
			output[i] = mapping[char-'A']
		case char >= 'a' && char <= 'z':
			output[i] = mapping[char-'a'] + ('a' - 'A')
		}
	}
	return string(output), nil
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(errorResponse{Error: message}); err != nil {
		log.Printf("writing error response: %v", err)
	}
}
