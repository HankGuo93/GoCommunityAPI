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

func RegisterArticleRoutes(router *gin.RouterGroup) {
	router.GET("", FetchArticlePage)
	router.GET("/:id", GetArticleDetail)
	router.POST("/", middlewares.RequereAuth, UploadArticle)
	router.PUT("/:id", middlewares.RequereAuth, UpdateArticle)
	router.DELETE("/:id", middlewares.RequereAuth, DeleteArticle)
}

// FetchArticlePage godoc
// @Description Retrieve a page of articles
// @Tags Article
// @Summary Get a page of articles
// @ID fetch-article-page
// @Produce  json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} gin.H
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/article [get]
func FetchArticlePage(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 5
	}
	articles, err := services.FetchArticlePage(page, pageSize)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"articles": dtos.CreateArticlePageResponse(articles),
	})
}

// GetArticleDetail godoc
// @Description Retrieve article by ID
// @Tags Article
// @Summary Get article by ID
// @ID get-article-by-id
// @Produce  json
// @Param id path int true "Article ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} dtos.ErrorDto
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/article/{id} [get]
func GetArticleDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid article id"))
	}
	article, err := services.GetArticleDetail(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Id":      article.Id,
		"Title":   article.Title,
		"Content": article.Content,
		"User": map[string]interface{}{
			"Id":    article.User.Id,
			"Name":  article.User.Name,
			"Email": article.User.Email,
		},
	})
}

// UploadArticle godoc
// @Description Upload a new article
// @Tags Article
// @Summary Upload a new article
// @ID upload-article
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param article body dtos.ArticleDto true "Article object"
// @Success 201 {object} gin.H
// @Failure 400 {object} dtos.ErrorDto
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/article [post]
func UploadArticle(c *gin.Context) {
	userId, _ := c.Get("userId")
	var json dtos.ArticleDto
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	err := services.UploadArticle(models.ArticleModel{
		Title:   json.Title,
		Content: json.Content,
		UserId:  userId.(int),
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database_error", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"messages": []string{"Article created successfully"}})
}

// UpdateArticle godoc
// @Description Update an article by ID
// @Tags Article
// @Summary Update an article by ID
// @ID update-article-by-id
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Article ID"
// @Param article body dtos.ArticleDto true "Article object"
// @Success 200 {string} string
// @Failure 400 {object} dtos.ErrorDto
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/article/{id} [put]
func UpdateArticle(c *gin.Context) {
	userId, _ := c.Get("userId")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid article id"))
	}
	var json dtos.ArticleDto
	if err = c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	err = services.UpdateArticle(models.ArticleModel{
		Id:      id,
		Title:   json.Title,
		Content: json.Content,
		UserId:  userId.(int),
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database_error", err))
		return
	}
}

// DeleteArticle godoc
// @Description Delete an article by ID
// @Tags Article
// @Summary Delete an article by ID
// @ID delete-article-by-id
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Article ID"
// @Success 200 {string} string
// @Failure 400 {object} dtos.ErrorDto
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid article id"))
	}
	err = services.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database_error", err))
		return
	}
}
