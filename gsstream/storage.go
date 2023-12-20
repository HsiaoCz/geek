package main

import (
	"fmt"
	"sync"
)

type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type MemoryStore struct {
	// concurency safe
	mu   sync.Mutex
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(b []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

func (s *MemoryStore) Fetch(offset int) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}
	return s.data[offset], nil
}
