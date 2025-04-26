package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

// @Summary 用户登录
// @Description 使用用户名密码登录系统
// @Tags 认证服务
// @Accept json
// @Produce json
// @Param login body model.LoginReq true "登录凭证"
// @Success 200 {object} model.LoginResp "登录成功"
// @Router /login [post]
func HandlerUserLogin(c *gin.Context){
	req := &types.LoginReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		responce.ErrorBadRequest(c,err)
		return
	}
	resp,code,err:=logic.LogicUserLogin(c,req)
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	if code == codes.CodeAllSuccess{
		responce.SuccessWithData(c,*resp)
	}else{
		responce.ErrorInternalServerErrorWithCode(c,code)
	}
}