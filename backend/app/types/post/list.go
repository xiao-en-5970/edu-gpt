package types

import "time"

// PostListReq 帖子列表请求参数
// @Description 获取帖子列表的请求参数
type PostListReq struct {
	LastPid uint `json:"last_pid" example:"100" comment:"上一页最后一条帖子ID，用于分页"`
	Limit   int  `json:"size" validate:"required,min=1,max=50" example:"10" comment:"每页数量"`
}

// BriefPost 简要帖子信息
// @Description 简要的帖子信息展示
type BriefPost struct {
	Nickname     string    `json:"poster_nickname" example:"张三" comment:"发帖人昵称"`
	ID           uint      `json:"id" example:"123" comment:"帖子ID"`
	PosterID     uint      `json:"poster_id" example:"456" comment:"发帖人ID"`
	Title        string    `json:"title" example:"这是一个帖子标题" comment:"帖子标题"`
	Content      string    `json:"content" example:"这是帖子的部分内容..." comment:"帖子内容摘要"`
	ViewCount    int       `json:"view_count" example:"100" comment:"浏览数"`
	LikeCount    int       `json:"like_count" example:"42" comment:"点赞数"`
	CollectCount int       `json:"collect_count" example:"5" comment:"收藏数"`
	CommentCount int       `json:"comment_count" example:"8" comment:"评论数"`
	CreateAt     time.Time `json:"create_at" example:"2025-04-27T16:10:08.5Z" comment:"创建时间"`
	Active       string    `json:"active" example:"active"`
	AvatarUrl    string    `json:"avatar_url" example:"https://example.com/avatar.jpg" comment:"发帖人头像URL"`
}

// PostListResp 帖子列表响应数据
// @Description 帖子列表的响应数据
type PostListResp []BriefPost
