package entities

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type UserEntity struct {
	gorm.Model
	Name      string                `gorm:"not null"`
	Email     string                `gorm:"not null;unique"`
	Password  string                `gorm:"not null"`
	CreatedAt int64                 `gorm:"autoCreateTime"`
	UpdatedAt int64                 `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete;default:null"`
}

func (UserEntity) TableName() string {
	return "users"
}
