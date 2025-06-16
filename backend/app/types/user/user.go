package types

import "time"

// BriefUser 简要用户信息
// @Description 简要的用户信息展示
type BriefUser struct {
	ID           uint      `json:"id" example:"1"`
	CreateAt     time.Time `json:"create_at" example:"2025-04-27T16:10:08.5Z"`
	Department   string    `json:"department" example:"数学学院"`
	Nickname     string    `json:"nickname" example:"傅益忠"`
	AvatarUrl    string    `json:"avatar_url" example:"https://127.0.0.1:8080/api/v1/user/auth/avatar/1"`
	BackImageUrl string    `json:"backimage_url" example:"https://127.0.0.1:8080/api/v1/user/auth/backimage/1"`
	Sex          string    `json:"sex" example:"男"`
	Grade        string    `json:"grade" example:"2022"`
	Campus       string    `json:"campus" example:"翡翠湖校区"`
	Signature    string    `json:"signature" example:"这人啥也没说"`
	Tags         []string  `json:"tags" example:"学习,运动,音乐"`
	Follows      int64     `json:"follows" example:"365"`
	Fans         int64     `json:"fans" example:"365"`
	AllPostLike  int64     `json:"allpost_like" example:"365"`
	FollowStatus int       `json:"follow_status" example:"1"`
}
