package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
)

// @Summary 评论点赞
// @Description 给评论点赞
// @Tags Comment模块
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param edit body types.CommentLikeReq true "请求体"
// @Success 200 {object} types.CommentLikeResp "成功响应"
// @Router /comment/auth/like [post]
func HandlerCommentLike(c * gin.Context){
	services.ServiceHandlerWithJson(c,CommentLike{})
}

type CommentLike struct{}

func (CommentLike) NewReq() any {
	return &types.CommentLikeReq{}
}

func (CommentLike) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicCommentLike(c, req.(*types.CommentLikeReq))
}

