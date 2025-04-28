package types

import "time"

type PostListReq struct{
	Offset int `json:"last_pid"`
	Limit int `json:"size"`
}

type BriefPost struct {
	Nickname     string    `json:"poster_nickname"`
	ID           uint      `json:"id"`
	PosterID     uint      `json:"poster_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ViewCount    int       `json:"view_count"`
	LikeCount    int       `json:"like_count"`
	CollectCount int       `json:"collect_count"`
	CommentCount int       `json:"comment_count"`
	CreateAt     time.Time `json:"create_at"`
}


type PostListResp []BriefPost