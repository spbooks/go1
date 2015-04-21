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
