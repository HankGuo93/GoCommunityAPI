package models

type ArticleModel struct {
	Id        int
	Title     string
	Content   string
	UserId   int
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}
