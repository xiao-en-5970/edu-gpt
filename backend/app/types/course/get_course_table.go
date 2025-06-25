package types

import types "github.com/xiao-en-5970/edu-gpt/backend/app/types/HFUT"

type GetCourseTableReq struct {
	Year     int `json:"year" binding:"required" example:"例:2026"`         //学年
	Term     int `json:"term" binding:"required" example:"1上2下"`           //学期
	CampusId int `json:"campus" binding:"required" example:"2表示合肥,23表示宣城"` //校区 【合肥校区/宣城校区】
}

type GetCourseTableResp []types.CourseTable
