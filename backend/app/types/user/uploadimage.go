package types

import "mime/multipart"

type UploadImageReq struct {
	File *multipart.FileHeader
}

type UploadImageResp struct {
	Url string `json:"url"`
}
