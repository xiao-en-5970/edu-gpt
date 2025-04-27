package types

import "time"

type BriefUser struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"create_at"`
	Department   string    `json:"department"`
	Nickname     string    `json:"nickname"`
	AvatarUrl    string    `json:"avatar_url"`
	BackImageUrl string    `json:"backimage_url"`
	Sex          string    `json:"sex"`
	Grade        string    `json:"grade"`
	Campus       string    `json:"campus"`
	Signature    string    `json:"signature"`
	Tags         string    `json:"tags"`
}
