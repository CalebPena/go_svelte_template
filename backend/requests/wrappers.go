package requests

import "github.com/gin-gonic/gin"

func ResponseWrapper(data gin.H) gin.H {
	return gin.H{"response": data}
}

func ErrorResponseWrapper(message string) gin.H {
	return gin.H{"error": message}
}
