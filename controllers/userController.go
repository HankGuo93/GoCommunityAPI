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
	router.POST("/login", Login)
}

// @Summary Find user by email
// @Description Get user by email
// @Tags User
// @Param email path string true "User Email"
// @Success 200 {object} gin.H
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/user/{email} [get]
func FindUser(c *gin.Context) {
	email := c.Param("email")
	user, err := services.GetUser(email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"name":     user.Name,
		"email":    user.Email,
		"CreateAt": user.CreatedAt,
		"UpdateAt": user.UpdatedAt,
		"DeleteAt": user.DeletedAt,
	})
}

// @Summary Create user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param input body dtos.UserDto true "User object to be created"
// @Success 201 {object} gin.H
// @Failure 422 {object} dtos.ErrorDto
// @Router /api/user/ [post]
func UserRegistration(c *gin.Context) {
	var json dtos.UserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}

	if err := services.CreateUser(models.UserModel{
		Name:     json.Name,
		Email:    json.Email,
		Password: json.Password,
	}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dtos.CreateErrorDto("database", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"messages": []string{"User created successfully"}})
}

// @Summary User login
// @Description Login with user email and password
// @Tags User
// @Accept json
// @Produce json
// @Param input body dtos.UserDto true "User email and password"
// @Success 200 {object} gin.H
// @Failure 401 {object} dtos.ErrorDto
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	var json dtos.UserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}

	tokenString, err := services.Login(models.UserModel{
		Email:    json.Email,
		Password: json.Password,
	})

	if err != nil {
		c.JSON(http.StatusForbidden, dtos.CreateErrorDto("login", err))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}
