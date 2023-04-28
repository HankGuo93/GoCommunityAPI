package services

import (
	"GoCommunityAPI/helpers"
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(email string) (models.UserModel, error) {
	user, err := repositories.FindOneUserByEmail(email)
	return user, err
}

func CreateUser(user models.UserModel) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	user.Password = string(hash)
	err = repositories.AddUser(user)
	return err
}

func Login(user models.UserModel) (string, error) {
	var tokenString string
	userInfo, err := repositories.FindOneUserByEmail(user.Email)
	if err != nil {
		return tokenString, errors.New("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(user.Password))
	if err != nil {
		return tokenString, errors.New("Invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, jwt.MapClaims{
		"sub": userInfo.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	privateKey, err := helpers.GetPrivateKey()
	if err != nil {
		return tokenString, errors.New(err.Error())
	}
	tokenString, err = token.SignedString(privateKey)
	if err != nil {
		return tokenString, errors.New("Failed to create token")
	}

	return tokenString, err
}
