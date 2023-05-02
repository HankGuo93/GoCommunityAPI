package repositories

import (
	"GoCommunityAPI/database"
	"GoCommunityAPI/database/entities"
	"GoCommunityAPI/models"
)

var FindOneUserById func(id int) (user models.UserModel, err error)
var FindOneUserByEmail func(email string) (user models.UserModel, err error)
var AddUser func(user models.UserModel) error

func init() {
	FindOneUserById = findOneUserById
	FindOneUserByEmail = findOneUserByEmail
	AddUser = addUser
}

func findOneUserById(id int) (user models.UserModel, err error) {
	db := database.DB
	var entity entities.UserEntity
	query := db.Model(&entities.UserEntity{})
	err = query.First(&entity, "id = ?", id).Error
	user = models.UserModel{
		Id:        int(entity.ID),
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return user, err
}

func findOneUserByEmail(email string) (user models.UserModel, err error) {
	db := database.DB
	var entity entities.UserEntity
	query := db.Model(&entities.UserEntity{})
	err = query.First(&entity, "email = ?", email).Error
	user = models.UserModel{
		Id:        int(entity.ID),
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return user, err
}

func addUser(user models.UserModel) error {
	db := database.DB
	entity := entities.UserEntity{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	err := db.Create(&entity).Error
	return err
}
