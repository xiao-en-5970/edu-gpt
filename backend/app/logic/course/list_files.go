package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/fileop"
)

func LogicCourseFiles(c *gin.Context, req *types.CourseFilesReq) (resp types.CourseFilesResp, code int, err error) {
	course, err := models.FindCourseById(c, req.CourseID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if course == nil {
		return resp, codes.CodeCourseNotExist, nil
	}
	finfo, err := fileop.GetAllFilesInfo(global.Cfg.Static.BookPath + "/" + course.CourseName)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.FileList = make([]types.File, 0, len(finfo))
	for _, file := range finfo {
		resp.FileList = append(resp.FileList, types.File{
			FileName: file.Name,
			FileSize: file.Size,
			FileType: file.FileType,
			ModTime:  file.ModTime,
			FileURL:  global.GetFileUrl("course/auth/getfile", course.CourseName+"/"+file.Name),
		})
	}
	return resp, codes.CodeAllSuccess, nil
}
