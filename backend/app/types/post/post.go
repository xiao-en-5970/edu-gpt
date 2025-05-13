package types

import "github.com/xiao-en-5970/edu-gpt/backend/app/model"

// PostResp 帖子详情响应数据
// @Description 帖子详情的响应数据
type PostResp struct {
	model.Post
	Nickname   string   `json:"poster_nickname" example:"张三" comment:"发帖人昵称"`
	Grade      string   `json:"poster_grade" example:"2022" comment:"发帖人年级"`
	Campus     string   `json:"poster_campus" example:"翡翠湖校区" comment:"发帖人校区"`
	Department string   `json:"poster_department" example:"计算机学院" comment:"发帖人院系"`
	ImageUrls  []string `json:"image_urls" example:"https://example.com/image1.jpg,https://example.com/image2.jpg" comment:"帖子图片URL列表"`
	LikeStatus int      `json:"like_status" example:"1" comment:"当前用户点赞状态(0:未点赞,1:已点赞)"`
	Avatar     string   `json:"avatar" example:"https://example.com/avatar.jpg" comment:"发帖人头像URL"`
}
