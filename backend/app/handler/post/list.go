package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/post"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)


func HandlerPostList(c * gin.Context){
	req:=&types.PostListReq{}
	err:=c.ShouldBindJSON(req)
	if err!=nil{
		responce.ErrorBadRequest(c,err)
		return
	}
	resp,code,err:=logic.LogicPostList(c,req)
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	if code!=codes.CodeAllSuccess{
		responce.ErrorInternalServerErrorWithCode(c,code)
		return
	}
	responce.SuccessWithData(c,resp)
}