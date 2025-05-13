package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
)
// @Summary 帖子列表
// @Description 输入最后刷新的帖子id和返回数量进行帖子列表的获取
// @Tags Post模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param edit body types.PostListReq true "请求体"
// @Success 200 {object} types.PostListResp "成功响应"
// @Router /post/auth/list [post]
func HandlerPostList(c *gin.Context) {
	services.ServiceHandlerWithJson(c, PostList{})
}

type PostList struct {
}

func (PostList) NewReq() any {
	return &types.PostListReq{}
}

func (PostList) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicPostList(c, req.(*types.PostListReq))
}
