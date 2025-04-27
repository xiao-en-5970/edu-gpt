package types

import "github.com/xiao-en-5970/edu-gpt/backend/app/model"

type PostResp struct {
	model.Post
	Nickname   string   `json:"poster_nickname"`
	Grade      string   `json:"poster_grade"`
	Campus     string   `json:"poster_campus"`
	Department string   `json:"poster_department"`
	ImageUrls  []string `json:"image_urls"`
}
