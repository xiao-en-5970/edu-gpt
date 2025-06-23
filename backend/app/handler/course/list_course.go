package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
)

func HandlerListCourses(c *gin.Context) {
	services.ServiceHandlerWithJson(c,ListCourses{})
}

type ListCourses struct{}

func (ListCourses) NewReq() any {
	return &types.ListCoursesReq{}
}

func (ListCourses) Logic(c *gin.Context, req any) (resp any, code int, err error) {
	return logic.LogicListCourses(c, req.(*types.ListCoursesReq))
}