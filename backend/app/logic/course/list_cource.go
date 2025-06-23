package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicListCourses(c *gin.Context, req *types.ListCoursesReq) (resp types.ListCoursesResp, code int, err error) {
	page := req.Page
	size := req.Size
	if page <= 0 {
		page = 1
	}
	resp, err = models.ListCourses(c, 294, req.OpenDepart, req.CourseType,req.Campus, page, size)
	if err != nil {
		return types.ListCoursesResp{}, codes.CodeAllIntervalError, err
	}
	return resp, codes.CodeAllSuccess, nil
}
