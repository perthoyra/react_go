package main

import (
	"log"
	"net/http"

	"restapi/internal/cipher"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/encrypt", cipher.Handler(false))
	mux.HandleFunc("/decrypt", cipher.Handler(true))

	log.Println("REST API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
