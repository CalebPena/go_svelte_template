package auth

import (
	"errors"
	"time"
	"todo/util"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte = []byte(util.GetEnvironmentVariable("JWT_SECRET"))

type claims struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Version int    `json:"version"`

	jwt.RegisteredClaims
}

func parseToken(tokenStr string) (User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(t *jwt.Token) (any, error) { return secret, nil })

	if err != nil {
		return User{}, err
	}

	c, ok := token.Claims.(*claims)

	if !ok {
		return User{}, errors.New("Invalid claim")
	}

	if c.Version != Version {
		return User{}, errors.New("Token is an invalid version out of date")
	}

	return User{c.Email, c.Id}, nil
}

func tokenize(user User) (string, error) {
	c := claims{
		user.Id,
		user.Email,
		Version,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, ExpireDays)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
