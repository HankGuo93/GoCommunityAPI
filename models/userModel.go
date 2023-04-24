package models

type UserModel struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}
