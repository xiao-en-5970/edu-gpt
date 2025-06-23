package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
)

type Course struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`                                                     // 课程ID
	CourseName   string    `gorm:"type:varchar(200);not null" json:"course_name"`                                          // 课程名字
	CourseCode   string    `gorm:"type:varchar(32);not null;index:idx_course_code" json:"course_code"`                     // 课程代码（索引）
	CourseType   string    `gorm:"type:varchar(32);not null" json:"course_type"`                                           // 课程类型
	Credits      float64   `gorm:"type:float;not null;" json:"credits"`                                                    // 学分
	OpenDepart   string    `gorm:"type:varchar(32);not null" json:"open_depart"`                                           // 开课学院
	ExamMod      string    `gorm:"type:varchar(32);not null" json:"exam_mod"`                                              // 考核方式
	Campus       string    `gorm:"type:varchar(32);not null" json:"campus"`                                                // 校区
	Description  string    `gorm:"type:varchar(256);not null;default:'尚无描述'" json:"description"`                           // 课程描述
	ViewCount    int       `gorm:"type:int;not null;default:0" json:"view_count"`                                          // 浏览数
	LikeCount    int       `gorm:"type:int;not null;default:0" json:"like_count"`                                          // 点赞数
	CommentCount int       `gorm:"type:int;not null;default:0" json:"comment_count"`                                       // 评论数
	SemesterCode int       `gorm:"type:int;not null;default:0" json:"semester_code"`                                       // 学期代码
	CreatedAt    time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_at"`                      // 创建时间
	UpdatedAt    time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;autoUpdateTime:milli" json:"update_at"` // 更新时间
}

// TableName 设置表名
func (Course) TableName() string {
	return "course"
}

func InsertCourse(c *gin.Context, semester int, course *Course) (id uint, err error) {
	course.SemesterCode = semester
	err = global.Db.WithContext(c).Model(course).Create(course).Error
	return course.ID, err
}

func CheckSemester(c *gin.Context, semester int) (*Course, error) {
	// 检查该学期课程是否已存在
	s := &Course{}
	err:=global.Db.WithContext(c).Model(&Course{}).Where("semester_code = ?", semester).First(s).Error
	if err != nil {
		if err.Error() == "record not found" {
			// 如果没有找到记录，则返回一个新的 Course 实例
			return nil, nil
		}
		return nil, err // 其他数据库错误
	}
	return s, nil // 返回找到的 Course 实例
}


func ListCourses(c *gin.Context, semester int,open_depart string, courseType string,campus string, page int, size int) ([]Course, error) {
	var courses []Course
	query := global.Db.WithContext(c).Model(&Course{}).Where("semester_code = ?", semester)

	if open_depart != "" {
		query = query.Where("open_depart = ?", open_depart)
	}

	if courseType != "" {
		query = query.Where("course_type = ?", courseType)
	}

	if campus != "" {
		query = query.Where("campus = ?", campus)
	}

	if err := query.Offset((page - 1) * size).Limit(size).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func FindCourseById(c *gin.Context, id uint) (*Course, error) {
	course := &Course{}
	if err := global.Db.WithContext(c).Model(&Course{}).Where("id = ?", id).First(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}