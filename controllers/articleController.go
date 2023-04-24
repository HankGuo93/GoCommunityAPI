package controllers

import "github.com/gin-gonic/gin"

func RegisteArticleRoutes(router *gin.RouterGroup) {
	router.GET("/", GetArticleList)
	router.GET("/:id", GetArticleDetail)
	router.POST("/", UploadArticle)
	router.POST("/:id", UpdateArticle)
	router.DELETE("/:id", ArticleDelete)
}

func GetArticleList(c *gin.Context) {

}

func GetArticleDetail(c *gin.Context) {

}

func UploadArticle(c *gin.Context) {

}

func UpdateArticle(c *gin.Context) {

}

func ArticleDelete(c *gin.Context) {

}
