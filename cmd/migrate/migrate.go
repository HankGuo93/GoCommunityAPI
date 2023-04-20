package main

import (
	"GoCommunityAPI/config"
	"GoCommunityAPI/database"
	"GoCommunityAPI/database/entities"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	database.DB.AutoMigrate(&entities.UserEntity{})
	database.DB.AutoMigrate(&entities.ArticleEntiry{})
	database.DB.AutoMigrate(&entities.CommentEntity{})
}
