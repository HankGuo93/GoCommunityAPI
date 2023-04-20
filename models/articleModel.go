package models

type ArticleModel struct {
	Id         int
	Title      string
	Content    string
	Author_Id  int
	Created_at int64
	Updated_at int64
}
