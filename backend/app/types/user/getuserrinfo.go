package types

import (
	"time"
)

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id"`
	UsernameZh     string    `json:"username_zh"`
	Sex            string    `json:"sex"`
	CultivateType  string    `json:"cultivate_type"`
	Department     string    `json:"department"`
	Grade          string    `json:"grade"`
	Level          string    `json:"level"`
	Major          string    `json:"major"`
	Class          string    `json:"class"`
	Campus         string    `json:"campus"`
	EnrollmentDate string    `json:"enrollment_date"`
	GraduateDate   string    `json:"graduate_date"`
	CreateAt       time.Time `json:"create_at"`
	Username       string    `json:"username" validate:"required,min=3,max=50"`
	AccountStatus  string    `json:"account_status" validate:"required,oneof=active locked disabled"`
	Nickname       string    `json:"nickname" validate:"required,min=1,max=50"`
	AvatarUrl      string    `json:"avatar_url"`
	BackImageUrl   string    `json:"backimage_url"`
	Signature      string    `json:"signature"`
	Tags           []string  `json:"tags"`
}
