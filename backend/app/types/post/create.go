package types

type CreatePostReq struct {
	
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostResp struct {
	ID uint `json:"id"`
}
