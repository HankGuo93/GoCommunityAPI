package models

type CommentModel struct {
	Content    string
	Author_id  int
	Article_id int
	Created_at int64
	Updated_at int64
}
