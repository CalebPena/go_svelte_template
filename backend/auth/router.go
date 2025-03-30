package auth

import (
	"github.com/gin-gonic/gin"
)

func Router(group *gin.RouterGroup) {
	group.POST("/register", registerHandler)
	group.POST("/login", loginHandler)
}
