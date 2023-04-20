package entities

import "gorm.io/gorm"

type ArticleEntiry struct {
	gorm.Model
	Id         int    `gorm:"not null"`
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	Author_Id  int    `gorm:"not null"`
	Created_at int64  `gorm:"not null"`
	Updated_at int64  `gorm:"not null"`
}
