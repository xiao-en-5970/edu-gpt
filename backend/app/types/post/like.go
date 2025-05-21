package types

// PostLikeReq 帖子点赞请求参数
// @Description 帖子点赞/取消点赞的请求参数
type PostLikeReq struct {
	PostID     uint `json:"post_id" validate:"required" example:"123" comment:"帖子ID"`
	LikeStatus int  `json:"like_status" validate:"required,oneof=0 1" example:"1" comment:"点赞状态(0:取消点赞,1:点赞)"`
}

// PostLikeResp 帖子点赞响应数据
// @Description 帖子点赞操作后返回的响应数据
type PostLikeResp struct {
	OK  int `json:"ok" example:"1" comment:"是否成功"`
}
