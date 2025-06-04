package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
)

func HandlerListFollow(c *gin.Context){
	services.ServiceHandlerWithJson(c,UserListFollow{})
}

type UserListFollow struct{}

func (UserListFollow) NewReq() any {
	return &types.FollowListReq{}
}

func (UserListFollow) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicFollowList(c, req.(*types.FollowListReq))
}