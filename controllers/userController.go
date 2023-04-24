package controllers

import (
	"GoCommunityAPI/dtos"
	"GoCommunityAPI/models"
	"GoCommunityAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.GET("/:email", FindUser)
	router.POST("/", UserRegistration)
}

func FindUser(c *gin.Context) {
	email := c.Param("email")
	user, err := services.GetUser(email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name":     user.Name,
		"email":    user.Email,
		"CreateAt": user.CreatedAt,
		"UpdateAt": user.UpdatedAt,
		"DeleteAt": user.DeletedAt,
	})
}

func UserRegistration(c *gin.Context) {
	var json dtos.UserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}

	if err := services.CreateUser(&models.UserModel{
		Name:     json.Name,
		Email:    json.Email,
		Password: json.Password,
	}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateDetailedErrorDto("database", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"messages": []string{"User created successfully"}})
}