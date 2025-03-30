package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	invalidJsonMessage   string = "Invalid Request"
	internalErrorMessage string = "Oops, something went wrong"
	forbidenMessage      string = "Forbidden"
)

func InvalidJson(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ErrorResponseWrapper(invalidJsonMessage))
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, ErrorResponseWrapper(forbidenMessage))
	c.Abort()
}

func InternalError(c *gin.Context, error error) {
	c.Error(error)
	c.JSON(http.StatusInternalServerError, ErrorResponseWrapper(internalErrorMessage))
}
