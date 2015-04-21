package main

import "code.google.com/p/go.crypto/bcrypt"

type User struct {
	ID             string
	Email          string
	HashedPassword string
	Username       string
}

const (
	hashCost       = 10
	passwordLength = 6
	userIDLength = 16
)

func NewUser(username, email, password string) (User, error) {
	user := User{
		Email:    email,
		Username: username,
	}
	if username == "" {
		return user, errNoUsername
	}

	if email == "" {
		return user, errNoEmail
	}

	if password == "" {
		return user, errNoPassword
	}

	if len(password) < passwordLength {
		return user, errPasswordTooShort
	}

	// Check if the username exists
	existingUser, err := globalUserStore.FindByUsername(username)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errUsernameExists
	}

	// Check if the email exists
	existingUser, err = globalUserStore.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	user.HashedPassword = string(hashedPassword)
	user.ID = GenerateID("usr", userIDLength)
	return user, err
}
