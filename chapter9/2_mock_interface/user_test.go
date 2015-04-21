package main

import "testing"

func TestNewUserNoUsername(t *testing.T) {
	_, err := NewUser("", "user@example.com", "password")
	if err != errNoUsername {
		t.Error("Expected err to be errNoUsername")
	}
}

func TestNewUserNoPassword(t *testing.T) {
	_, err := NewUser("user", "user@example.com", "")
	if err != errNoPassword {
		t.Error("Expected err to be errNoUsername")
	}
}

type MockUserStore struct {
	findUser         *User
	findEmailUser    *User
	findUsernameUser *User
	saveUser         *User
}

func (store *MockUserStore) Find(string) (*User, error) {
	return store.findUser, nil
}

func (store *MockUserStore) FindByEmail(string) (*User, error) {
	return store.findEmailUser, nil
}

func (store *MockUserStore) FindByUsername(string) (*User, error) {
	return store.findUsernameUser, nil
}

func (store *MockUserStore) Save(user User) error {
	store.saveUser = &user
	return nil
}

func TestNewUserExistingUsername(t *testing.T) {
	globalUserStore = &MockUserStore{
		findUsernameUser: &User{},
	}

	_, err := NewUser("user", "user@example.com", "somepassword")
	if err != errUsernameExists {
		t.Error("Expected err to be errUsernameExists")
	}
}

func TestNewUserExistingEmail(t *testing.T) {
	globalUserStore = &MockUserStore{
		findEmailUser: &User{},
	}

	_, err := NewUser("user", "user@example.com", "somepassword")
	if err != errEmailExists {
		t.Error("Expected err to be errEmailExists")
	}
}
