package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
)

//课程表接口，赶紧完成
//还需要对接hfutapi接口
func HandlerCourseGetTable(c *gin.Context) {
	services.ServiceHandlerWithJson(c, CourseGetTable{})
}

type CourseGetTable struct{}
func (CourseGetTable) NewReq() any {
	return &types.GetCourseTableReq{}
}
func (CourseGetTable) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicCourseGetTable(c, req.(*types.GetCourseTableReq))
}