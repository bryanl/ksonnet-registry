package store

import (
	"sync"

	"github.com/pkg/errors"
)

// Store manages files
type Store interface {
	Write(key string, data []byte) error
	Read(key string) ([]byte, error)
	Delete(key string) error
	Exists(key string) bool
}

// MemoryStore stores contents in memory
type MemoryStore struct {
	dict map[string][]byte

	sync.RWMutex
}

var _ Store = (*MemoryStore)(nil)

// NewMemoryStore creates an instance of MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		dict: make(map[string][]byte),
	}
}

// Write writes bytes to a key in the MemoryStore.
func (s *MemoryStore) Write(key string, data []byte) error {
	s.Lock()
	defer s.Unlock()

	if s.Exists(key) {
		return errors.Errorf("key %q exists", key)
	}

	s.dict[key] = data
	return nil
}

// Read returns the bytes at a key in the MemoryStore.
func (s *MemoryStore) Read(key string) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	if !s.Exists(key) {
		return nil, errors.Errorf("key %q does not exist", key)
	}

	return s.dict[key], nil
}

// Delete removes a key from the MemoryStore.
func (s *MemoryStore) Delete(key string) error {
	s.Lock()
	defer s.Unlock()

	if !s.Exists(key) {
		return errors.Errorf("key %q does not exist", key)
	}

	delete(s.dict, key)
	return nil
}

// Exists returns if a key exists in the MemoryStore.
func (s *MemoryStore) Exists(key string) bool {
	_, ok := s.dict[key]
	return ok
}
