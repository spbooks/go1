package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FileStore struct {
	filename string
}

func NewFileStore(name string) *FileStore {
	return &FileStore{
		filename: name,
	}
}

func (f *FileStore) Read(target interface{}) error {
	contents, err := ioutil.ReadFile(f.filename)

	if err != nil {
		// If it's a matter of the file not existing, that's ok
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	fmt.Printf("Contents of %s: %s\n", f.filename, contents)
	err = json.Unmarshal(contents, &target)
	return err
}

func (f *FileStore) Write(content interface{}) error {
	contents, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return fmt.Errorf("Error saving user store to %s: %s", f.filename, err)
	}

	err = ioutil.WriteFile(f.filename, contents, 0660)
	if err != nil {
		return fmt.Errorf("Error saving user store to %s: %s", f.filename, err)
	}
	return nil
}
