package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
)

func HandlerFollow(c *gin.Context){
	services.ServiceHandlerWithJson(c,UserFollow{})
}

type UserFollow struct{}

func (UserFollow) NewReq() any {
	return &types.UserAddFollowReq{}
}

func (UserFollow) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicUserAddFollow(c, req.(*types.UserAddFollowReq))
}