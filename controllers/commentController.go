package controllers

import (
	"GoCommunityAPI/dtos"
	"GoCommunityAPI/models"
	"GoCommunityAPI/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(router *gin.RouterGroup) {
	router.GET("/articleId/:articleId", GetCommentsByArticleId)
	router.POST("/", UploadComment)
	router.DELETE("/:id", DeleteComment)
}

func GetCommentsByArticleId(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid comment id"))
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSuze"))
	if err != nil {
		pageSize = 5
	}
	comments, err := services.GetCommentsByArticleId(articleId, page, pageSize)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"articles": dtos.CreateCommentPageResponse(comments),
	})
}

func UploadComment(c *gin.Context) {
	var json dtos.CommentDto
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	err := services.UploadComment(models.CommentModel{
		ArticleId: json.ArticleId,
		Content:   json.Content,
		UserId:    json.UserId,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database_error", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"messages": []string{"Comment created successfully"}})
}

func DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid comment id"))
	}
	err = services.DeleteComment(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database_error", err))
		return
	}
}
