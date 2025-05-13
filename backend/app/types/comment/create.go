package types

// CommentCreateReq 创建评论请求参数
// @Description 创建评论的请求参数
type CommentCreateReq struct {
	PostID  uint   `json:"post_id" validate:"required" example:"123" comment:"帖子ID"`
	Content string `json:"content" validate:"required,min=1,max=500" example:"这是一条评论内容" comment:"评论内容"`
}

// CommentCreateResp 创建评论响应数据
// @Description 创建评论成功后返回的响应数据
type CommentCreateResp struct {
	ID uint `json:"id" example:"456" comment:"评论ID"`
}

// SubCommentCreateReq 创建子评论请求参数
// @Description 创建子评论(回复评论)的请求参数
type SubCommentCreateReq struct {
	PostID          uint   `json:"post_id" validate:"required" example:"123" comment:"帖子ID"`
	ParentCommentID uint   `json:"parent" validate:"required" example:"456" comment:"父评论ID"`
	ReplyUserID     uint   `json:"reply" example:"789" comment:"回复的用户ID"`
	Content         string `json:"content" validate:"required,min=1,max=500" example:"这是一条回复评论" comment:"评论内容"`
}

// SubCommentCreateResp 创建子评论响应数据
// @Description 创建子评论成功后返回的响应数据
type SubCommentCreateResp struct {
	ID uint `json:"id" example:"789" comment:"子评论ID"`
}
