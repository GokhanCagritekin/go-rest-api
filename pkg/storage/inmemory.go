package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
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

func (s *Storage) Save() error {
	_, err := json.Marshal(s.DB)
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(s.DB, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
	return nil
}

func (s *Storage) Recover() error {
	c := make(map[string]json.RawMessage)

	absPath, _ := filepath.Abs("data.json")
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	e := json.Unmarshal(data, &c)
	if e != nil {
		panic(e)
	}

	i := 0
	for k, v := range c {
		s.Set(k, string(v))
		i++
	}
	return nil
}
