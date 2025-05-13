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
	LikeCount  int `json:"like_count" example:"10" comment:"当前点赞总数"`
	LikeStatus int `json:"like_status" example:"1" comment:"当前点赞状态(0:未点赞,1:已点赞)"`
}

// SubCommentLikeReq 子评论点赞请求参数
// @Description 子评论点赞/取消点赞的请求参数
type SubCommentLikeReq struct {
	SubCommentID uint `json:"subcomment_id" validate:"required" example:"789" comment:"子评论ID"`
	LikeStatus   int  `json:"like_status" validate:"required,oneof=0 1" example:"1" comment:"点赞状态(0:取消点赞,1:点赞)"`
}

// SubCommentLikeResp 子评论点赞响应数据
// @Description 子评论点赞操作后返回的响应数据
type SubCommentLikeResp struct {
	LikeCount  int `json:"like_count" example:"5" comment:"当前点赞总数"`
	LikeStatus int `json:"like_status" example:"1" comment:"当前点赞状态(0:未点赞,1:已点赞)"`
}
