package storage

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
)

type Storage struct {
	DB   map[string]string
	lock sync.RWMutex
}

func NewStorage() *Storage {
	db := new(Storage)
	return db
}

func (s *Storage) Set(key string, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DB == nil {
		s.DB = make(map[string]string)
	}
	s.DB[key] = value
}

func (s *Storage) Get(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	v, ok := s.DB[key]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}
