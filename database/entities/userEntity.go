package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Name       string `gorm:"not null"`
	Email      string `gorm:"not null"`
	Password   string `gorm:"not null"`
	Created_at int64  `gorm:"not null"`
	Updated_at int64  `gorm:"not null"`
}
