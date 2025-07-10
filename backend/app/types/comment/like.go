package types

// CommentLikeReq 评论点赞请求参数
// @Description 评论点赞/取消点赞的请求参数
type CommentLikeReq struct {
	CommentID  uint `json:"comment_id" validate:"required" example:"456" comment:"评论ID"`
	LikeStatus int  `json:"like_status" validate:"required,oneof=0 1" example:"1" comment:"点赞状态(0:取消点赞,1:点赞)"`
}

// CommentLikeResp 评论点赞响应数据
// @Description 评论点赞操作后返回的响应数据
type CommentLikeResp struct {
	OK  int `json:"ok" example:"1" comment:"是否成功"`
}
