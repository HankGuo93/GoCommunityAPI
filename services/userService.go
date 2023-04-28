package services

import (
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
	"errors"
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
