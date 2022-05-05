package models

import "time"

type Question struct {
	ID          int64     `json:"id,string" db:"post_id"`                                   // 帖子id
	UserName    string    `json:"username" db:"username"`                                   // 作者name
	CommunityID int64     `json:"community_id,string" db:"community_id" binding:"required"` // 社区id
	Status      int32     `json:"status" db:"status"`                                       // 帖子状态
	Title       string    `json:"title" db:"title" `                                        // 帖子标题
	Content     string    `json:"content" db:"content"`                                     // 帖子内容
	Location    string    `json:"location" db:"location" binding:"required"`
	AudioPath   string    `json:"audio_path" db:"audio_path"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type QueList struct {
	ID         int64     `json:"id,string" db:"post_id"` // 帖子id
	UserName   string    `json:"username" db:"username"`
	Title      string    `json:"title" db:"title"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type QueCh struct {
	Ch       int64  `json:"ch" db:"community_id" binding:"required"`
	Location string `json:"location" db:"location" binding:"required"`
}
