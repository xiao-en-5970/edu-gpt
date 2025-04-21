package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/logic"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/responce"
)

// @Summary 用户登录
// @Description 使用用户名密码登录系统
// @Tags 认证服务
// @Accept json
// @Produce json
// @Param login body model.LoginReq true "登录凭证"
// @Success 200 {object} model.LoginResp "登录成功"
// @Router /login [post]
func HandlerLogin(c *gin.Context){
	req := &types.LoginReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		responce.ErrorBadRequest(c,err)
		return
	}
	resp,code,err:=logic.LogicLogin(c,req)
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	if code == codes.CodeUserLoginSuccess{
		responce.SuccessWithCodeData(c,code,*resp)
	}else{
		responce.SuccessWithCode(c,code)
	}
}