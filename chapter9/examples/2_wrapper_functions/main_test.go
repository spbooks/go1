package main

import (
	"errors"
	"os"
	"testing"
)

func TestWriteStuff(t *testing.T) {
	testStr := "Test String"
	testLoc := "testfile.txt"

	// Store the current writeFile and reinstate it at the end of the test
	oldWriteFile := writeFile
	defer func() {
		writeFile = oldWriteFile
	}()

	// Create a new function that tests the passed parameters
	writeFile = func(loc string, data []byte, perm os.FileMode) error {
		if loc != testLoc {
			t.Error("Expected loc to be", testLoc, "but got", loc)
		}
		if string(data) != testStr {
			t.Error("Expected data to be", testStr, "but got", string(data))
		}
		return nil
	}

	err := WriteStuff(testStr, testLoc)
	if err != nil {
		t.Error("Expected no error, but got", err)
	}
}

func TestWriteStuffError(t *testing.T) {
	testStr := "Test String"
	testLoc := "testfile.txt"
	testErr := errors.New("Ah! An error.")

	// Store the current writeFile and reinstate it at the end of the test
	oldWriteFile := writeFile
	defer func() {
		writeFile = oldWriteFile
	}()

	// create a new function that returns an error
	writeFile = func(loc string, data []byte, perm os.FileMode) error {
		return testErr
	}

	err := WriteStuff(testStr, testLoc)
	if err != testErr {
		t.Error("Expected an error, but got", err)
	}
}
