package services

import (
	"GoCommunityAPI/helpers"
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
	"GoCommunityAPI/services"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser_Success(t *testing.T) {
	// Arrange
	user := models.UserModel{
		Email:    "test@example.com",
		Password: "password",
		Name:     "Test User",
	}

	// Mock the repository function
	repositories.AddUser = func(user models.UserModel) error {
		return nil
	}

	// Act
	err := services.CreateUser(user)

	// Assert
	assert.NoError(t, err)
}

func TestLogin_Success(t *testing.T) {
	// Arrange
	user := models.UserModel{
		Email:    "test@example.com",
		Password: "password",
		Name:     "Test User",
	}

	repositories.FindOneUserByEmail = func(email string) (models.UserModel, error) {
		if email == user.Email {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(hashedPassword)
			return user, nil
		}
		return user, nil
	}

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	helpers.GetPrivateKey = func() (*rsa.PrivateKey, error) {
		return privateKey, nil
	}

	// Act
	token, err := services.Login(user)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestLogin_InvalidUser(t *testing.T) {
	// Arrange
	user := models.UserModel{
		Email:    "test@example.com",
		Password: "password",
		Name:     "Test User",
	}

	// Mock the repository function
	repositories.FindOneUserByEmail = func(email string) (models.UserModel, error) {
		return models.UserModel{}, errors.New("User not found")
	}

	// Act
	token, err := services.Login(user)

	// Assert
	assert.Equal(t, errors.New("Invalid email or password"), err)
	assert.Empty(t, token)
}

func TestLogin_InvalidPassword(t *testing.T) {
	// Arrange
	user := models.UserModel{
		Email:    "test@example.com",
		Password: "password",
		Name:     "Test User",
	}

	// Mock the repository function
	repositories.FindOneUserByEmail = func(email string) (models.UserModel, error) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hashedPassword)
		return user, nil
	}

	// Act
	_, err := services.Login(models.UserModel{
		Email:    user.Email,
		Password: "wrong_password",
		Name:     user.Name,
	})

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "Invalid email or password", err.Error())
}

func TestLogin_GetPrivateKeyError(t *testing.T) {
	// Arrange
	user := models.UserModel{
		Email:    "test@example.com",
		Password: "password",
		Name:     "Test User",
	}

	repositories.FindOneUserByEmail = func(email string) (models.UserModel, error) {
		if email == user.Email {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			user.Password = string(hashedPassword)
			return user, nil
		}
		return models.UserModel{}, errors.New("User not found")
	}

	helpers.GetPrivateKey = func() (*rsa.PrivateKey, error) {
		return nil, errors.New("Key not found")
	}

	// Act
	_, err := services.Login(user)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "Key not found", err.Error())
}
