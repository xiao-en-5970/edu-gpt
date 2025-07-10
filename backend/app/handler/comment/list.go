package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
)


// @Summary 评论列表
// @Description 输入帖子id返回评论列表
// @Tags Comment模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param edit body types.CommentListReq true "请求体"
// @Success 200 {object} types.CommentListResp "成功响应"
// @Router /comment/auth/list [post]
func HandlerCommentList(c *gin.Context) {
	services.ServiceHandlerWithJson(c,CommentList{})
}


type CommentList struct{}

func (CommentList) NewReq() any {
	return &types.CommentListReq{}
}

func (CommentList) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicCommentList(c, req.(*types.CommentListReq))
}
