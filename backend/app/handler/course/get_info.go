package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
)

func HandlerCourseGetInfo(c *gin.Context) {
	services.ServiceHandlerWithJson(c, CourseGetInfo{})
}
type CourseGetInfo struct{}
func (CourseGetInfo) NewReq() any {
	return &types.GetCourseInfoReq{}
}

func (CourseGetInfo) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicCourseGetInfo(c, req.(*types.GetCourseInfoReq))
}
