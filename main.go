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

	apiRouteGroup := goEngin.Group("/api")

	controllers.RegisterUserRoutes(apiRouteGroup.Group("/user"))
	controllers.RegisterArticleRoutes(apiRouteGroup.Group("/article"))
	controllers.RegisterCommentRoutes(apiRouteGroup.Group("/comment"))

	goEngin.Run()
}
