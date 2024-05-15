package models

type User struct {
	UserID   int64  `db:"user_id,string"`
	UserName string `db:"username"`
	PassWord string `db:"password"`
	Token    string
}
