package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
)
// @Summary 用户个人信息
// @Description 获取用户个人信息
// @Tags User模块
// @Security BearerAuth 
// @Accept json
// @Produce json
// @Param get_userinfo body types.GetUserInfoReq true "请求体"
// @Success 200 {object} types.GetUserInfoResp "成功响应"
// @Router /user/auth/get_userinfo [post]
func HandlerUserGetUserInfo(c *gin.Context) {
	services.ServiceHandlerWithEmpty(c, UserGetUserInfo{})
}

type UserGetUserInfo struct{}

func (UserGetUserInfo) NewReq() any {
	return &types.GetUserInfoReq{}
}

func (UserGetUserInfo) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	global.Logger.Infof("logicuserinfo success!")
	return logic.LogicUserGetUserInfo(c, req.(*types.GetUserInfoReq))
}
