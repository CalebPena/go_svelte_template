package auth

import (
	"database/sql"
)

func createUser(db *sql.DB, email string, password string) (User, error) {
	user := User{Email: email}
	err := db.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id", email, password).Scan(&user.Id)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func getUserByEmail(db *sql.DB, email string) (User, string, error) {
	user := User{Email: email}
	var passwordHash string
	err := db.QueryRow("SELECT id, password FROM users WHERE email = $1", email).Scan(&user.Id, &passwordHash)

	if err != nil {
		return User{}, "", err
	}

	return user, passwordHash, nil
}
