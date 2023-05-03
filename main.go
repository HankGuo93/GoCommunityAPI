package main

import (
	"GoCommunityAPI/config"
	"GoCommunityAPI/controllers"
	"GoCommunityAPI/database"
	"GoCommunityAPI/database/entities"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDB()
	migrate()
}

func main() {
	goEngin := gin.Default()

	apiRouteGroup := goEngin.Group("/api")

	controllers.RegisterUserRoutes(apiRouteGroup.Group("/user"))
	controllers.RegisterArticleRoutes(apiRouteGroup.Group("/article"))
	controllers.RegisterCommentRoutes(apiRouteGroup.Group("/comment"))

	goEngin.Run()
}

func migrate() {
	if !database.DB.Migrator().HasTable(&entities.UserEntity{}) {
		database.DB.AutoMigrate(&entities.UserEntity{})
	}

	if !database.DB.Migrator().HasTable(&entities.ArticleEntity{}) {
		database.DB.AutoMigrate(&entities.ArticleEntity{})
	}

	if !database.DB.Migrator().HasTable(&entities.CommentEntity{}) {
		database.DB.AutoMigrate(&entities.CommentEntity{})
	}
}
