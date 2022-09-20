package models

import (
	"sync"
	"time"
)

type Commit struct {
	Id         int       `json:"id,omitempty" db:"id"`
	PostId     int64     `json:"post_id,string,omitempty" db:"post_id"`
	Content    string    `json:"content" db:"content"`
	UserName   string    `json:"username" db:"username"`
	ToUserName string    `json:"to_user" db:"to_user"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type NameMap struct {
	M map[string]string
	S sync.Mutex
}

var NM = NameMap{
	M: make(map[string]string),
}
