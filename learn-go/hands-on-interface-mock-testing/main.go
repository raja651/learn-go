package main

import (
	"fmt"
	"log"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type MockDataStore struct {
	Users map[int]User
}

func (m *MockDataStore) GetUser(id int) (User, error) {
	user, ok := m.Users[id]

	if !ok {
		return User{}, fmt.Errorf("user %v not found", id)
	}

	return user, nil
}

func (m *MockDataStore) SaveUser(u User) error {
	_, ok := m.Users[u.ID]

	if ok {
		return fmt.Errorf("user %v already exists", u.ID)
	}

	m.Users[u.ID] = u
	return nil
}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	md := make(map[int]User)

	db := &MockDataStore{
		Users: md,
	}

	user1 := User{
		ID:   1,
		Name: "Krithu",
		Age:  2,
	}

	// ser := Service{
	// 	ds: db,
	// }

	err := db.SaveUser(user1)

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	u1, err := db.GetUser(user1.ID)

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	fmt.Println(user1)
	fmt.Println(u1)
}
