package models

type CommentModel struct {
	Content   string
	UserId    int
	ArticleId int
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}
