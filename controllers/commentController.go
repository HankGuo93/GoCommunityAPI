package controllers

import (
	"GoCommunityAPI/dtos"
	"GoCommunityAPI/middlewares"
	"GoCommunityAPI/models"
	"GoCommunityAPI/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(router *gin.RouterGroup) {
	router.GET("/articleId/:articleId", FetchCommentPageByArticleId)
	router.POST("/", middlewares.RequereAuth, UploadComment)
	router.DELETE("/:id", middlewares.RequereAuth, DeleteComment)
}

// @Summary Fetch comments by article ID
// @Description Get a list of comments for a given article
// @Tags Comment
// @Accept json
// @Produce json
// @Param articleId path int true "Article ID"
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of comments per page"
// @Success 200 {object} gin.H
// @Failure 400 {object} dtos.ErrorDto
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/comment/articleId/{articleId} [get]
func FetchCommentPageByArticleId(c *gin.Context) {
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
	comments, err := services.FetchCommentPageByArticleId(articleId, page, pageSize)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"articles": dtos.CreateCommentPageResponse(comments),
	})
}

// @Summary Upload a comment
// @Description Create a new comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param comment body dtos.CommentDto true "Comment object that needs to be added"
// @Success 201 {object} gin.H
// @Failure 400 {object} dtos.ErrorDto
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/comment/ [post]
func UploadComment(c *gin.Context) {
	userId, _ := c.Get("userId")
	var json dtos.CommentDto
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	err := services.UploadComment(models.CommentModel{
		ArticleId: json.ArticleId,
		Content:   json.Content,
		UserId:    userId.(int),
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database_error", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"messages": []string{"Comment created successfully"}})
}

// @Summary Delete a comment
// @Description Delete a comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Comment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} dtos.ErrorDto
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid comment id"))
	}
	err = services.DeleteComment(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database_error", err))
		return
	}
}
