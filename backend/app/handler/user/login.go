package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
)

// @Summary 用户登录
// @Description 使用信息门户用户名密码登录软件
// @Tags User模块
// @Accept json
// @Produce json
// @Param login body types.LoginReq true "请求体"
// @Success 200 {object} types.LoginResp "成功响应"
// @Router /user/login [post]
func HandlerUserLogin(c *gin.Context) {
	services.ServiceHandlerWithJson(c, UserLogin{})
}

type UserLogin struct{}

func (UserLogin) NewReq() any {
	return &types.LoginReq{}
}

func (UserLogin) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicUserLogin(c, req.(*types.LoginReq))
}
