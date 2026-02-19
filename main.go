package main

import (
	"log"
	"net/http"
)

func main() {
	store, err := NewStorage("data.json")
	if err != nil {
		log.Fatalf("failed to create storage: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/contacts", makeContactsHandler(store))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
