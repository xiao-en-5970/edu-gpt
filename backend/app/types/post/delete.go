package types

type PostDeleteReq struct {
	PostID uint `json:"post_id" validate:"required" example:"123" comment:"帖子ID"`
}

type PostDeleteResp struct {
	OK int `json:"ok" example:"1" comment:"是否成功"`
}
