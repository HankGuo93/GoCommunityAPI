package services

import (
	"GoCommunityAPI/models"
	"GoCommunityAPI/repositories"
)

func GetUser(email string) (models.UserModel, error) {
	user, err := repositories.FindOneUserByEmail(email)
	return user, err
}

func CreateUser(user *models.UserModel) error {
	err := repositories.AddUser(user)
	return err
}
