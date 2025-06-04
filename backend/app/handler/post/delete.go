package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
)

func HandlerPostDelete(c *gin.Context){
	services.ServiceHandlerWithJson(c,PostDelete{})
}

type PostDelete struct {
}

func (PostDelete) NewReq() any {
	return &types.PostDeleteReq{}
}

func (PostDelete) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicPostDelete(c, req.(*types.PostDeleteReq))
}