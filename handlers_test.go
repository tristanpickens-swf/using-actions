package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestPostContactHandler(t *testing.T) {
	// Setup storage
	tmpPath := filepath.Join(t.TempDir(), "test.json")
	store, _ := NewStorage(tmpPath)
	handler := makeContactsHandler(store)

	// Create a mock request
	contact := Contact{Name: "Bob", Phone: "123"}
	body, _ := json.Marshal(contact)
	req := httptest.NewRequest(http.MethodPost, "/contacts", bytes.NewBuffer(body))

	// Create a response recorder (acts like a ResponseWriter)
	rr := httptest.NewRecorder()

	// Execute handler
	handler.ServeHTTP(rr, req)

	// Assertions
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var returned Contact
	json.NewDecoder(rr.Body).Decode(&returned)
	if returned.Name != "Bob" {
		t.Errorf("Expected name Bob, got %s", returned.Name)
	}
}
