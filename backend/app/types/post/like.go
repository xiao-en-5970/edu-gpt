package types

type LikeReq struct {
	PostID     uint `json:"post_id"`
	LikeStatus int  `json:"like_status"`
}

type LikeResp struct {
	LikeCount int `json:"like_count"`
	LikeStatus int `json:"like_status"`
}

