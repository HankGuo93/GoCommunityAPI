package models

type CommentModel struct {
	Id        int
	Content   string
	UserId    int
	ArticleId int
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64

	User UserModel
}
