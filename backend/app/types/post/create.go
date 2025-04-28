package types

import "mime/multipart"

type CreatePostReq struct {
	ImageCount int `json:"image_count"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostResp struct {
	ID uint `json:"id"`
}

type UploadManyImagesReq struct{
	ID uint
	Files []*multipart.FileHeader 
}

type UploadManyImagesResp struct{
	Urls []string `json:"url"`
}