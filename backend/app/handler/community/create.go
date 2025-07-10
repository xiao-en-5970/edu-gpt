package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/community"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/community"
)




func HandlerCommunityCreate(c *gin.Context){
	services.ServiceHandlerWithJson(c,CommunityCreate{})
}

type CommunityCreate struct {
}

func (CommunityCreate) NewReq() any {
	return &types.CommunityCreateReq{}
}

func (CommunityCreate) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicCommunityCreate(c, req.(*types.CommunityCreateReq))
}