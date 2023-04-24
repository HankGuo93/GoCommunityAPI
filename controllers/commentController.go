package controllers

import "github.com/gin-gonic/gin"

func RegisterCommentRoutes(router *gin.RouterGroup) {
	router.GET("/:articleId", GetCommentsByArticleId)
	router.POST("/", UploadComment)
	router.DELETE("/:commentId", CommentDelete)
}

func GetCommentsByArticleId(c *gin.Context) {

}

func UploadComment(c *gin.Context) {

}

func CommentDelete(c *gin.Context) {

}
