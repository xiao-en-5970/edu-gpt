package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCourseGetInfo(c *gin.Context, req *types.GetCourseInfoReq) (resp types.GetCourseInfoResp, code int, err error) {
	course, err := models.FindCourseById(c, req.CourseID)
	
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if course == nil {
		return resp, codes.CodeCourseNotExist, nil
	}
	resp = types.GetCourseInfoResp(*course)
	return resp, codes.CodeAllSuccess, nil
}