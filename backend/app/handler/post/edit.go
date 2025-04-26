package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/post"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerPostEdit(c* gin.Context){
	req := &types.EditPostReq{}
	err := c.ShouldBindJSON(req)
	if err !=nil{
		responce.ErrorBadRequest(c,err)
	}
	resp,code,err:=logic.LogicPostEdit(c,req)
	if code != codes.CodeAllSuccess{
		responce.ErrorInternalServerErrorWithCode(c,code)
		return
	}
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
	}
	responce.SuccessWithData(c,resp)
}