package types

import "time"

type BriefUser struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"create_at"`
	Department string    `json:"department"`
	Nickname   string    `json:"nickname"`
	AvatarPath string    `json:"avatar_path"`
	Sex        string    `json:"sex"`
	Grade      string    `json:"grade"`
	Campus     string    `json:"campus"`
	Signature  string    `json:"signature"`
	Tags       string    `json:"tags"`
}
