package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
)
// @Summary 帖子点赞
// @Description 输入帖子id预期点赞状态(1/0)给帖子点赞
// @Tags Post模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param edit body types.PostLikeReq true "请求体"
// @Success 200 {object} types.PostLikeResp "成功响应"
// @Router /post/auth/like [post]
func HandlerPostLike(c *gin.Context) {
	services.ServiceHandlerWithJson(c,PostLike{})
}

type PostLike struct{}

func (PostLike) NewReq() any {
	return &types.PostLikeReq{}
}
func (PostLike) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicPostLike(c, req.(*types.PostLikeReq))
}
