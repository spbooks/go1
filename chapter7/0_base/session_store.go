package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type SessionStore interface {
	Find(string) (*Session, error)
	Save(*Session) error
	Delete(*Session) error
}

var globalSessionStore SessionStore

func init() {
	sessionStore, err := NewFileSessionStore("./data/sessions.json")
	if err != nil {
		panic(fmt.Errorf("Error creating session store: %s", err))
	}
	globalSessionStore = sessionStore
}

type FileSessionStore struct {
	filename string
	Sessions map[string]Session
}

func NewFileSessionStore(name string) (*FileSessionStore, error) {
	store := &FileSessionStore{
		Sessions: map[string]Session{},
		filename: name,
	}

	contents, err := ioutil.ReadFile(name)

	if err != nil {
		// If it's a matter of the file not existing, that's ok
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}
	err = json.Unmarshal(contents, store)
	if err != nil {
		return nil, err
	}
	return store, err
}

func (s *FileSessionStore) Find(id string) (*Session, error) {
	session, exists := s.Sessions[id]
	if !exists {
		return nil, nil
	}

	return &session, nil
}

func (store *FileSessionStore) Save(session *Session) error {
	store.Sessions[session.ID] = *session
	contents, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(store.filename, contents, 0660)
}

func (store *FileSessionStore) Delete(session *Session) error {
	delete(store.Sessions, session.ID)
	contents, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(store.filename, contents, 0660)
}
