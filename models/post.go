package models

import "time"

type Post struct {
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"authorID" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
