package entities

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type CommentEntity struct {
	gorm.Model
	Content   string                `gorm:"not null"`
	CreatedAt int64                 `gorm:"autoCreateTime"`
	UpdatedAt int64                 `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete;default:null"`

	User      UserEntity    `gorm:"association_foreignkey:UserId:"`
	UserId    int           `gorm:"not null"`
	Article   ArticleEntiry `gorm:"association_foreignkey:Article_id"`
	ArticleId int           `gorm:"not null"`
}

func (CommentEntity) TableName() string {
	return "comments"
}
