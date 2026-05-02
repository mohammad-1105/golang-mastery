package main

import (
	"errors"
	"fmt"
)

// domain model
type User struct {
	Name string
}

// interface (behavior contract)
type UserStore interface {
	Save(u User) error
}

// Real implementation (postgres)
type PostgresStore struct{}

func (p PostgresStore) Save(u User) error {
	if u.Name == "" {
		return errors.New("empty user name")
	}
	fmt.Println("[Postgres] saving user: ", u.Name)
	return nil
}

// Mock implementation (for testing)
type MockStore struct {
	Fail bool
}

func (m MockStore) Save(u User) error {
	if m.Fail {
		return errors.New("mock failure")
	}
	fmt.Println("[Mock] saving user: ", u.Name)
	return nil
}

// service layer
type UserService struct {
	store UserStore
}

// constructor (important pattern)
func NewUserService(store UserStore) UserService {
	return UserService{store: store}
}

// business logic
func (s UserService) Register(name string) error {
	if name == "" {
		return errors.New("Name Cannot be empty")
	}

	user := User{Name: name}

	// Never ignore errors like before
	if err := s.store.Save(user); err != nil {
		return fmt.Errorf("register failed: %w", err)
	}

	return nil
}
func main() {

	// production setup
	pgStore := PostgresStore{}
	service := NewUserService(pgStore)

	if err := service.Register("Alice"); err != nil {
		fmt.Println("error: ", err)
	}

	// Mock : test like scenario
	mock := MockStore{Fail: true}
	mockService := NewUserService(mock)

	if err := mockService.Register("TestUser"); err != nil {
		fmt.Println("expected error: ", err)
	}

}
