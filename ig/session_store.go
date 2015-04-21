package main

import "fmt"

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
	fileStore *FileStore
	Sessions  map[string]Session
}

func NewFileSessionStore(name string) (*FileSessionStore, error) {
	fileStore := NewFileStore(name)

	store := &FileSessionStore{
		Sessions:  map[string]Session{},
		fileStore: fileStore,
	}

	err := fileStore.Read(store)
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
	return store.fileStore.Write(store)
}

func (store *FileSessionStore) Delete(session *Session) error {
	delete(store.Sessions, session.ID)
	return store.fileStore.Write(store)

}
