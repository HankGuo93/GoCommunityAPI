package main

import (
	"GoCommunityAPI/config"
	"GoCommunityAPI/controllers"
	"GoCommunityAPI/database"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	goEngin := gin.Default()

	goEngin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRouteGroup := goEngin.Group("/api")

	controllers.RegisterUserRoutes(apiRouteGroup.Group("/user"))

	goEngin.Run()
}
