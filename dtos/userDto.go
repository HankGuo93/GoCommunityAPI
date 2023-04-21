package dtos

import "GoCommunityAPI/models"

type UserDto struct {
	Name      string
	Email     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

func CreateUserDto(user *models.UserModel) UserDto {
	return UserDto{
		Name:      user.Email,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
