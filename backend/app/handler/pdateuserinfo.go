package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/logic"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/responce"
)

func HandlerUpdateUserInfo(c *gin.Context){
	resp,code,err:=logic.LogicGetUserInfo(c,&types.GetUserInfoReq{})
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