package entities

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type ArticleEntiry struct {
	gorm.Model
	Title     string                `gorm:"not null"`
	Content   string                `gorm:"not null"`
	CreatedAt int64                 `gorm:"autoCreateTime"`
	UpdatedAt int64                 `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete;default:null"`

	User   UserEntity `gorm:"association_foreignkey:UserId:"`
	UserId int        `gorm:"not null"`
}

func (ArticleEntiry) TableName() string {
	return "articles"
}
