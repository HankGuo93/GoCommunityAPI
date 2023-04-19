package main

import (
	"GoCommunityAPI/config"
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
	goEngin.Run()
}
