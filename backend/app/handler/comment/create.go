package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
)

func HandlerCommentCreate(c *gin.Context) {
	services.ServiceHandlerWithJson(c, CommentCreate{})
}

type CommentCreate struct{}

func (CommentCreate) NewReq() any {
	return &types.CommentCreateReq{}
}

func (CommentCreate) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicCommentCreate(c, req.(*types.CommentCreateReq))
}

func HandlerSubCommentCreate(c *gin.Context) {
	services.ServiceHandlerWithJson(c, SubCommentCreate{})
}

type SubCommentCreate struct{}

func (SubCommentCreate) NewReq() any {
	return &types.SubCommentCreateReq{}
}

func (SubCommentCreate) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicSubCommentCreate(c, req.(*types.SubCommentCreateReq))
}
