package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/logic"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/responce"
)

func HandlerRegister(c *gin.Context){
	req := &types.RegisterReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		responce.ErrorBadRequest(c,err)
		return
	}
	_,code,err:=logic.LogicRegister(c,req)
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	if code == codes.CodeUserRegisterSuccess{
		responce.SuccessWithCode(c,code)
	}else{
		responce.SuccessWithCode(c,code)
	}
}