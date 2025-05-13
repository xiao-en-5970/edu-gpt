package types

import "time"

// UpdateUserInfoReq 更新用户信息请求参数
// @Description 更新用户信息的请求参数
type UpdateUserInfoReq struct {
	AccountStatus string   `json:"account_status" validate:"required,oneof=active locked disabled" example:"active"`
	Nickname      string   `json:"nickname" validate:"required,min=1,max=50" example:"李华2333"`
	Signature     string   `json:"signature" example:"这个人说了什么"`
	Tags          []string `json:"tags" example:"唱,跳,rap,篮球"`
}

// UpdateUserInfoResp 更新用户信息响应数据
// @Description 更新用户信息后返回的响应数据
type UpdateUserInfoResp struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id" example:"1"`
	UsernameZh     string    `json:"username_zh" example:"傅益忠"`
	Sex            string    `json:"sex" example:"男"`
	CultivateType  string    `json:"cultivate_type" example:"主修"`
	Department     string    `json:"department" example:"数学学院"`
	Grade          string    `json:"grade" example:"2022"`
	Level          string    `json:"level" example:"本科"`
	Major          string    `json:"major" example:"信息与计算科学"`
	Class          string    `json:"class" example:"信息计22-2班"`
	Campus         string    `json:"campus" example:"翡翠湖校区"`
	EnrollmentDate string    `json:"enrollment_date" example:"2022-09-01"`
	GraduateDate   string    `json:"graduate_date" example:"2026-07-01"`
	CreateAt       time.Time `json:"create_at" example:"2025-04-27T16:10:08.5Z"`
	Username       string    `json:"username" validate:"required,min=3,max=50" example:"2022210826"`
	AccountStatus  string    `json:"account_status" validate:"required,oneof=active locked disabled" example:"active"`
	Nickname       string    `json:"nickname" validate:"required,min=1,max=50" example:"傅益忠"`
	AvatarUrl      string    `json:"avatar_url" example:"https://127.0.0.1:8080/api/v1/user/auth/avatar/1"`
	BackImageUrl   string    `json:"backimage_url" example:"https://127.0.0.1:8080/api/v1/user/auth/backimage/1"`
	Signature      string    `json:"signature" example:"这人啥也没说"`
	Tags           []string  `json:"tags" example:"学习,运动,音乐"` // 数组类型用逗号分隔示例
}
