package types

type EditPostReq struct{
	ID     uint   `json:"id" validate:"required"`
	Title          string    `json:"title" comment:"标题"`
	Content        string    `json:"content" comment:"内容（除标题）"`
}

type EditPostResp CreatePostResp