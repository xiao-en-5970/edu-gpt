package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerCourseGetAllCourses(c *gin.Context) {
	code,err:=logic.LogicGetAllCourses(c,274)
	if err !=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	if code != codes.CodeAllSuccess{
		responce.ErrorInternalServerErrorWithCode(c,code)
		return
	}
	responce.Success(c)
}
