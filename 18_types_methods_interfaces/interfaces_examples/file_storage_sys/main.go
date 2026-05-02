package main

import (
	"errors"
	"fmt"
)

// File storage system

// behavior contract
type Storage interface {
	Save(key, value string) error
	Load(key string) (string, error)
}

// implementation 1 - in-memory (fast for testing)

type MemoryStorage struct {
	data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]string)}
}

func (m *MemoryStorage) Save(key, value string) error {
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Load(key string) (string, error) {
	val, ok := m.data[key]

	if !ok {
		return "", errors.New("not found")
	}
	return val, nil
}

// Implementation 2 - File storage (simulated)
type FileStorage struct{}

func (f *FileStorage) Save(key, value string) error {
	fmt.Println("[File] saving: ", key, value)
	return nil
}

func (f *FileStorage) Load(key string) (string, error) {
	fmt.Println("[File] loading: ", key)
	return "some-data", nil
}

// Business Logic layer
type ConfigService struct {
	store Storage
}

func NewConfigService(store Storage) *ConfigService {
	return &ConfigService{store: store}
}

func (c *ConfigService) SetConfig(key, value string) error {
	return c.store.Save(key, value)
}

func (c *ConfigService) GetConfig(key string) (string, error) {
	return c.store.Load(key)
}

// usage here
func main() {

	// dev / test
	memStore := NewMemoryStorage()
	config := NewConfigService(memStore)

	_ = config.SetConfig("theme", "dark")
	val, _ := config.GetConfig("theme")
	fmt.Println("memory: ", val)

	// production
	fileStore := &FileStorage{}
	configProd := NewConfigService(fileStore)

	_ = configProd.SetConfig("theme", "light")

}
