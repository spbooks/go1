package main

import (
	"fmt"
	"strings"
)

type UserStore interface {
	Find(string) (*User, error)
	FindByEmail(string) (*User, error)
	FindByUsername(string) (*User, error)
	Save(*User) error
}

type FileUserStore struct {
	fileStore *FileStore
	Users     map[string]User
}

var globalUserStore UserStore

func init() {
	userStore, err := NewFileUserStore("./data/users.json")
	if err != nil {
		panic(fmt.Errorf("Error creating user store: %s", err))
	}

	globalUserStore = userStore
}

func NewFileUserStore(name string) (*FileUserStore, error) {
	fileStore := NewFileStore(name)

	store := &FileUserStore{
		Users:     map[string]User{},
		fileStore: fileStore,
	}

	err := fileStore.Read(store)
	return store, err
}

func (store FileUserStore) Find(id string) (*User, error) {
	user, ok := store.Users[id]
	if ok {
		return &user, nil
	}
	return nil, nil
}

func (store FileUserStore) FindByUsername(username string) (*User, error) {
	if username == "" {
		return nil, nil
	}

	for _, user := range store.Users {
		if strings.ToLower(username) == strings.ToLower(user.Username) {
			return &user, nil
		}
	}
	return nil, nil
}

func (store FileUserStore) FindByEmail(email string) (*User, error) {
	if email == "" {
		return nil, nil
	}

	for _, user := range store.Users {
		if strings.ToLower(email) == strings.ToLower(user.Email) {
			return &user, nil
		}
	}
	return nil, nil
}

// Should I just be writing to a file here without keeping it open?
// If I was to do that I'd probably need to add an access mutex yea?
func (store FileUserStore) Save(user *User) error {
	store.Users[user.ID] = *user
	return store.fileStore.Write(store)
}
