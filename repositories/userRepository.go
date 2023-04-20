package repositories

import (
	"GoCommunityAPI/database"
	"GoCommunityAPI/database/entities"
	"GoCommunityAPI/models"
)

func FindOneUser(email string) (user models.UserModel, err error) {
	db := database.DB
	query := db.Model(&entities.UserEntity{}).Where(&entities.UserEntity{Email: email})
	var entity entities.UserEntity
	err = query.Find(&entity).Error
	user = models.UserModel{
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return user, err
}

func AddUser(user *models.UserModel) error {
	db := database.DB
	entity := entities.UserEntity{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	err := db.Create(&entity).Error
	return err
}