package dtos

import "GoCommunityAPI/models"

type UserDto struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

func CreateUserDto(user *models.UserModel) UserDto {
	return UserDto{
		Id:        user.Id,
		Name:      user.Email,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
