package types

import "mime/multipart"

// CreatePostReq 创建帖子请求参数
// @Description 创建帖子的请求参数
type CreatePostReq struct {
	ImageCount  int    `json:"image_count" validate:"required,min=0,max=9" example:"3" comment:"图片数量"`
	Title       string `json:"title" validate:"required,min=1,max=100" example:"这是一个帖子标题" comment:"帖子标题"`
	Content     string `json:"content" validate:"required,min=1,max=5000" example:"这是帖子的详细内容..." comment:"帖子内容"`
	CommunityID uint `json:"community_id" example:"社区id" comment:"社区ID"`
}

// CreatePostResp 创建帖子响应数据
// @Description 创建帖子成功后返回的响应数据
type CreatePostResp struct {
	ID uint `json:"id" example:"123" comment:"帖子ID"`
}

// UploadManyImagesReq 批量上传图片请求参数
// @Description 批量上传帖子图片的请求参数
type UploadManyImagesReq struct {
	ID    uint                  `form:"id" binding:"required" example:"123" comment:"帖子ID"`
	Files []*multipart.FileHeader `form:"files" binding:"required,min=1,max=9" comment:"图片文件列表"`
}

// UploadManyImagesResp 批量上传图片响应数据
// @Description 批量上传图片成功后返回的响应数据
type UploadManyImagesResp struct {
	Urls []string `json:"url" example:"https://example.com/image1.jpg,https://example.com/image2.jpg" comment:"图片URL列表"`
}
