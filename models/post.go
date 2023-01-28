package models

import "time"

type Post struct {
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"authorID" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	*Post                   //`json:"post"`
	*CommunityDetail `json:" community_detail"`
}
