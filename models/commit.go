package models

import "time"

type Commit struct {
	Id         int       `json:"id,omitempty" db:"id"`
	PostId     int64     `json:"post_id,string,omitempty" db:"post_id"`
	Content    string    `json:"content" db:"content"`
	UserName   string    `json:"username" db:"username"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
