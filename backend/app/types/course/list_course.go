package types

import "github.com/xiao-en-5970/edu-gpt/backend/app/models"

type ListCoursesReq struct {
	Page       int    `json:"page" form:"page"`               // 页码
	Size       int    `json:"size" form:"size"`               // 每页大小
	CourseType string `json:"course_type" form:"course_type"` // 课程类型
	// 各专业选修课
	// 学科基础和专业必修课
	// 实践环节
	// 通识必修课
	// 创新创业教育
	OpenDepart string `json:"open_depart" form:"open_depart"` // 开课学院
	// 仪器科学与光电工程学院
	// 机械工程学院
	// 材料科学与工程学院
	// 电气与自动化工程学院
	// 计算机与信息学院（人工智能学院）
	// 化学与化工学院
	// 土木与水利工程学院
	// 建筑与艺术学院
	// 资源与环境工程学院
	// 物理学院
	// 微电子学院
	// 管理学院
	// 马克思主义学院
	// 数学学院
	// 食品与生物工程学院
	// 外国语学院
	// 软件学院
	// 汽车与交通工程学院
	// 经济学院
	// 文法学院
	// 体育部
	// 本科生院工程素质教育中心
	// 党委学生工作部（处）
	// 创新创业教育处
	Campus string `json:"campus" form:"campus"` // 校区
	//翡翠湖校区
	//屯溪路校区
	//其他

}
type ListCoursesResp []models.Course
