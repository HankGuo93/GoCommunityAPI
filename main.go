package main

import (
	"GoCommunityAPI/config"
	"GoCommunityAPI/controllers"
	"GoCommunityAPI/database"
	"GoCommunityAPI/database/entities"

	_ "GoCommunityAPI/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	goEngin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
