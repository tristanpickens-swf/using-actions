package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func makeContactsHandler(store *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "invalid body", http.StatusBadRequest)
				return
			}
			var c Contact
			if err := json.Unmarshal(body, &c); err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			if c.Name == "" || c.Phone == "" {
				http.Error(w, "name and phone required", http.StatusBadRequest)
				return
			}
			created, err := store.AddContact(c)
			if err != nil {
				http.Error(w, "failed to save", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(created)
		case http.MethodGet:
			out := store.ListContacts()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(out)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
