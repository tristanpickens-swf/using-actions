package main

import (
	"path/filepath"
	"testing"
)

func TestAddAndList(t *testing.T) {
	// Create a temp directory so we don't mess up real data
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.json")

	store, err := NewStorage(dbPath)
	if err != nil {
		t.Fatalf("Failed to create storage: %v", err)
	}

	// Test Adding
	c := Contact{Name: "Tester", Phone: "555-0199"}
	added, err := store.AddContact(c)
	if err != nil {
		t.Errorf("Failed to add contact: %v", err)
	}

	if added.ID != 1 {
		t.Errorf("Expected ID 1, got %d", added.ID)
	}

	// Test Listing
	list := store.ListContacts()
	if len(list) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(list))
	}
	if list[0].Name != "Tester" {
		t.Errorf("Expected Name 'Tester', got %s", list[0].Name)
	}
}
