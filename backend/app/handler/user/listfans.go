package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
)

func HandlerListFans(c *gin.Context){
	services.ServiceHandlerWithJson(c,UserListFans{})
}

type UserListFans struct{}

func (UserListFans) NewReq() any {
	return &types.FansListReq{}
}

func (UserListFans) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicFansList(c, req.(*types.FansListReq))
}