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
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"articles": dtos.CreateArticlePageResponse(articles),
	})
}

func GetArticleDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid article id"))
	}
	article, err := services.GetArticleDetail(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database", err))
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
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database_error", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"messages": []string{"Article created successfully"}})
}

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
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database_error", err))
		return
	}
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must provide a valid article id"))
	}
	err = services.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database_error", err))
		return
	}
}
