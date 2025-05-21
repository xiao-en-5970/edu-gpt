package types

import "time"

// CommentListReq 评论列表请求参数
// @Description 获取评论列表的请求参数
type CommentListReq struct {
	PostID  uint `json:"post_id" validate:"required" example:"123" comment:"帖子ID"`
	LastCid uint `json:"last_cid" example:"100" comment:"上一页最后一条评论ID，用于分页"`
	Limit   int  `json:"size" validate:"required,min=1,max=50" example:"10" comment:"每页数量"`
}

// BriefComment 简要评论信息
// @Description 简要的评论信息展示
type BriefComment struct {
	Nickname   string    `json:"poster_nickname" example:"张三" comment:"评论者昵称"`
	PostID     uint      `json:"post_id" example:"123" comment:"帖子ID"`
	ID         uint      `json:"id" example:"456" comment:"评论ID"`
	PosterID   uint      `json:"poster_id" example:"789" comment:"评论者ID"`
	Content    string    `json:"content" example:"这是一条评论内容" comment:"评论内容"`
	LikeCount  int       `json:"like_count" example:"10" comment:"点赞数"`
	ChildCount int       `json:"comment_count" example:"3" comment:"子评论数量"`
	LikeStatus int       `json:"like_status" example:"1" comment:"当前用户点赞状态(0:未点赞,1:已点赞)"`
	CreateAt   time.Time `json:"create_at" example:"2025-04-27T16:10:08.5Z" comment:"创建时间"`
	Active     string    `json:"active" example:"active" comment:"激活状态"`
	AvatarUrl  string    `json:"avatar_url" example:"https://example.com/avatar.jpg" comment:"评论者头像URL"`
}

// CommentListResp 评论列表响应数据
// @Description 评论列表的响应数据
type CommentListResp []BriefComment

// SubCommentListReq 子评论列表请求参数
// @Description 获取子评论列表的请求参数
type SubCommentListReq struct {
	ParentCommentID uint `json:"parent" validate:"required" example:"456" comment:"父评论ID"`
	LastSCid        uint `json:"last_scid" example:"100" comment:"上一页最后一条子评论ID，用于分页"`
	Limit           int  `json:"size" validate:"required,min=1,max=50" example:"10" comment:"每页数量"`
}

// BriefSubComment 简要子评论信息
// @Description 简要的子评论信息展示
type BriefSubComment struct {
	Nickname        string    `json:"poster_nickname" example:"李四" comment:"评论者昵称"`
	ID              uint      `json:"id" example:"789" comment:"子评论ID"`
	Post_id         uint      `json:"post_id" example:"123" comment:"帖子ID"`
	PosterID        uint      `json:"poster_id" example:"456" comment:"评论者ID"`
	Content         string    `json:"content" example:"这是一条回复评论" comment:"评论内容"`
	ReplyUserID     uint      `json:"reply" example:"101" comment:"回复的用户ID"`
	ParentCommentID uint      `json:"parent" example:"456" comment:"父评论ID"`
	LikeCount       int       `json:"like_count" example:"5" comment:"点赞数"`
	CreateAt        time.Time `json:"create_at" example:"2025-04-27T16:10:08.5Z" comment:"创建时间"`
	Active          string    `json:"active" example:"30分钟前" comment:"活跃时间显示"`
	AvatarUrl       string    `json:"avatar_url" example:"https://example.com/avatar.jpg" comment:"评论者头像URL"`
}

// SubCommentListResp 子评论列表响应数据
// @Description 子评论列表的响应数据
type SubCommentListResp []BriefSubComment
