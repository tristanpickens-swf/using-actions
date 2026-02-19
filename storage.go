package main

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email,omitempty"`
}

type Storage struct {
	mu       sync.Mutex
	contacts []Contact
	nextID   int
	path     string
}

func NewStorage(path string) (*Storage, error) {
	// If the configured path is a directory (commonly happens when a host
	// path that should be a file didn't exist and Docker created a dir),
	// use a data.json file inside that directory.
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		path = filepath.Join(path, "data.json")
	}

	s := &Storage{path: path, nextID: 1}
	if err := s.load(); err != nil {
		if err == fs.ErrNotExist || os.IsNotExist(err) {
			return s, nil
		}
		return nil, err
	}
	return s, nil
}

func (s *Storage) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	f, err := os.Open(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	var data []Contact
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		if err == io.EOF {
			s.contacts = nil
			s.nextID = 1
			return nil
		}
		return err
	}
	s.contacts = data
	max := 0
	for _, c := range data {
		if c.ID > max {
			max = c.ID
		}
	}
	s.nextID = max + 1
	return nil
}

func (s *Storage) save() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(s.contacts); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return os.Rename(tmp, s.path)
}

func (s *Storage) AddContact(c Contact) (Contact, error) {
	s.mu.Lock()
	c.ID = s.nextID
	s.nextID++
	s.contacts = append(s.contacts, c)
	s.mu.Unlock()
	if err := s.save(); err != nil {
		return Contact{}, err
	}
	return c, nil
}

func (s *Storage) ListContacts() []Contact {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Contact, len(s.contacts))
	copy(out, s.contacts)
	return out
}
