package types

import "mime/multipart"

// UploadImageReq 上传图片请求参数
// @Description 上传图片的请求参数
type UploadImageReq struct {
	File *multipart.FileHeader `form:"file" binding:"required"` // 上传的文件
}

// UploadImageResp 上传图片响应数据
// @Description 上传图片成功后返回的响应数据
type UploadImageResp struct {
	Url string `json:"url" example:"https://127.0.0.1:8080/api/v1/user/auth/avatar/1"`
}
