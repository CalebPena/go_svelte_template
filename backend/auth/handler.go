package auth

import (
	"net/http"
	"net/mail"
	"strings"
	"todo/database"
	"todo/requests"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func loginHandler(c *gin.Context) {
	var data loginRequest

	err := c.ShouldBindJSON(&data)

	if err != nil {
		requests.InvalidJson(c)
		return
	}

	data.Email = strings.TrimSpace(data.Email)

	db := database.GetDb(c)
	user, err := login(db, data.Email, data.Password)

	if err == bcrypt.ErrMismatchedHashAndPassword || err == errUserNotFound {
		c.JSON(http.StatusUnauthorized, requests.ErrorResponseWrapper("Invalid email or password"))
		return
	}
	if err != nil {
		requests.InternalError(c, err)
		return
	}

	token, err := tokenize(user)

	if err != nil {
		requests.InternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.ResponseWrapper(gin.H{"token": token}))
}

type registerRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func registerHandler(c *gin.Context) {
	var data registerRequest

	err := c.ShouldBindJSON(&data)

	if err != nil {
		requests.InvalidJson(c)
		return
	}

	_, err = mail.ParseAddress(data.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, requests.ErrorResponseWrapper("Invalid email"))
		return
	}
	if len(data.Password) < MinPasswordLength {
		c.JSON(http.StatusBadRequest, requests.ErrorResponseWrapper("Password is too short"))
		return
	}

	data.Email = strings.TrimSpace(data.Email)

	db := database.GetDb(c)

	user, err := registerUser(db, data.Email, data.Password)

	if err == bcrypt.ErrPasswordTooLong {
		c.JSON(http.StatusBadRequest, requests.ErrorResponseWrapper("Password is too long"))
		return
	}
	if err == errUserAlreadyExists {
		c.JSON(http.StatusBadRequest, requests.ErrorResponseWrapper("User with that email already exists"))
		return
	}
	if err != nil {
		requests.InternalError(c, err)
		return
	}

	token, err := tokenize(user)

	if err != nil {
		requests.InternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.ResponseWrapper(gin.H{"token": token}))
}

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	bearer := "Bearer "

	if len(token) <= len(bearer) || token[:len(bearer)] != bearer {
		requests.Forbidden(c)
		return
	}

	token = token[len(bearer):]

	user, err := parseToken(token)

	if err != nil {
		requests.Forbidden(c)
		return
	}

	c.Set(UserKey, user)
}

func GetUser(c *gin.Context) User {
	return c.MustGet(UserKey).(User)
}
