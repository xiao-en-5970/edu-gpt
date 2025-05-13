package types

// EditPostReq 编辑帖子请求参数
// @Description 编辑帖子的请求参数
type EditPostReq struct {
	ID      uint   `json:"id" validate:"required" example:"123" comment:"帖子ID"`
	Title   string `json:"title" validate:"required,min=1,max=100" example:"修改后的标题" comment:"标题"`
	Content string `json:"content" validate:"required,min=1,max=5000" example:"修改后的内容..." comment:"内容（除标题）"`
}

// EditPostResp 编辑帖子响应数据
// @Description 编辑帖子后返回的响应数据
type EditPostResp CreatePostResp
