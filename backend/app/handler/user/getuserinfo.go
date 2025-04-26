package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerUserGetUserInfo(c* gin.Context){
	
	resp,code,err:=logic.LogicUserGetUserInfo(c,&types.GetUserInfoReq{})
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