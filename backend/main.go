package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"todo/auth"
	"todo/database"
	"todo/util"
)

var port string = util.GetEnvironmentVariable("BACKEND_PORT")

func init() {
	if !util.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
}

var i int = 0

func Hello(c *gin.Context) {
	debugStr := "off"
	if util.IsDebug {
		debugStr = "on"
	}

	user := auth.GetUser(c)

	fmt.Fprintf(c.Writer, "Hello %d. Debug mode is %s. Your email is %s", i, debugStr, user.Email)
	i++
}

func main() {
	db, err := database.InitDb()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.Use(database.DbMiddleware(db))

	auth.Router(r.Group("/auth"))

	p := r.Group("/")
	p.Use(auth.AuthMiddleware)

	p.GET("/ping", Hello)

	r.Run(":" + port)
}
