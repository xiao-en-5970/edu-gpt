package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
)

// @Summary 关注列表
// @Description 粉关注列表
// @Tags User模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param list_fans body types.FollowListReq true "请求体"
// @Success 200 {object} types.FollowListResp "成功响应"
// @Router /user/auth/list_follow [post]
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