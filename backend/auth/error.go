package auth

import "errors"

var errUserNotFound error = errors.New("todo/auth: user not found")
var errUserAlreadyExists error = errors.New("todo/auth: user already exists")
