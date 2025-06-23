package types

import "github.com/xiao-en-5970/edu-gpt/backend/app/models"

type GetCourseInfoReq struct {
	CourseID uint `json:"course_id" form:"course_id"` // 课程ID
}

type GetCourseInfoResp models.Course