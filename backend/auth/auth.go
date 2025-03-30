package auth

import (
	"database/sql"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	ExpireDays        = 30
	Version           = 0
	MinPasswordLength = 8
	UserKey           = "user"
)

type User struct {
	Email string
	Id    int
}

func login(db *sql.DB, email, password string) (User, error) {
	user, passwordHash, err := getUserByEmail(db, email)

	if err == sql.ErrNoRows {
		return User{}, errUserNotFound
	}
	if err != nil {
		return User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func registerUser(db *sql.DB, email, password string) (User, error) {
	email = strings.TrimSpace(email)
	_, _, err := getUserByEmail(db, email)

	if err == nil {
		return User{}, errUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return User{}, err
	}

	user, err := createUser(db, email, string(hashedPassword))

	if err != nil {
		return User{}, err
	}

	return user, nil
}
