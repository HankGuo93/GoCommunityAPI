package entities

import "gorm.io/gorm"

type CommentEntity struct {
	gorm.Model
	Content    string `gorm:"not null"`
	Author_id  int    `gorm:"not null"`
	Article_id int    `gorm:"not null"`
	Created_at int64  `gorm:"not null"`
	Updated_at int64  `gorm:"not null"`
}
