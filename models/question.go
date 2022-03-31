package models

import "time"

type Question struct {
	ID          int64     `json:"id,string" db:"post_id"`                            // 帖子id
	UserName    string    `json:"username" db:"username"`                            // 作者name
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区id
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	Title       string    `json:"title" db:"title" binding:"required"`               // 帖子标题
	Content     string    `json:"content" db:"content" binding:"required"`           // 帖子内容
	Location    string    `json:"location" db:"location" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type QueList struct {
	ID         int64     `json:"id,string" db:"post_id"` // 帖子id
	Title      string    `json:"title" db:"title"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type QueCh struct {
	Ch       int64  `json:"ch" db:"community_id" binding:"required"`
	Location string `json:"location" db:"location" binding:"required"`
}
