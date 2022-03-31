package models

type Commit struct {
	Id       int    `json:"id" db:"id"`
	PostId   int64  `json:"post_id" db:"post_id"`
	Content  string `json:"content" db:"content"`
	UserName string `json:"username" db:"username"`
}
