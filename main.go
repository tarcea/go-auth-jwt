package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-auth-jwt/controllers"
	"github.com/tarcea/go-auth-jwt/ini"
	"github.com/tarcea/go-auth-jwt/middleware"
)

func init() {
	ini.LoadEnv()
	ini.ConnectToDb()
	ini.Migrate()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
