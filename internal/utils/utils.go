package utils

import "sync"

type Set struct {
	mu sync.RWMutex
	m  map[string]bool
}

func NewSet() *Set {
	return &Set{m: make(map[string]bool)}
}

func (s *Set) Add(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[id] = true
}

func (s *Set) Contains(id string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.m[id]
	return exists
}

func (s *Set) Remove(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.m[id]; exists {
		delete(s.m, id)
		return true
	}
	return false
}
