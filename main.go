package main

import (
	"GoCommunityAPI/config"
	"GoCommunityAPI/databases"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	databases.ConnectToDB()
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
