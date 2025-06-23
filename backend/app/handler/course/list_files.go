package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
)

func HandlerCourseFiles(c *gin.Context) {
	services.ServiceHandlerWithJson(c, CourseFiles{})
}
type CourseFiles struct{}

func (CourseFiles) NewReq() any {
	return &types.CourseFilesReq{}
}

func (CourseFiles) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicCourseFiles(c, req.(*types.CourseFilesReq))
}