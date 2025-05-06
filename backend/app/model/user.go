package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"

	"gorm.io/gorm"
)

type AccountStatus string

type User struct {
	ID                  uint      `gorm:"column:id;primaryKey;autoIncrement;comment:用户ID" json:"id"`
	Password            string    `gorm:"column:password;type:varchar(255)" json:"-"`
	UsernameEn          string    `gorm:"column:username_en;type:varchar(100)" json:"username_en"`
	UsernameZh          string    `gorm:"column:username_zh;type:varchar(100)" json:"username_zh"`
	Sex                 string    `gorm:"column:sex;type:enum('男','女','其他')" json:"sex"`
	CultivateType       string    `gorm:"column:cultivate_type;type:varchar(50)" json:"cultivate_type"`
	Department          string    `gorm:"column:department;type:varchar(100)" json:"department"`
	Grade               string    `gorm:"column:grade;type:varchar(20)" json:"grade"`
	Level               string    `gorm:"column:level;type:varchar(50)" json:"level"`
	StudentType         string    `gorm:"column:student_type;type:varchar(50)" json:"student_type"`
	Major               string    `gorm:"column:major;type:varchar(100)" json:"major"`
	Class               string    `gorm:"column:class;type:varchar(50)" json:"class"`
	Campus              string    `gorm:"column:campus;type:varchar(50)" json:"campus"`
	Status              string    `gorm:"column:status;type:varchar(50)" json:"status"`
	Length              string    `gorm:"column:length;type:decimal(3,1)" json:"length"`
	EnrollmentDate      string    `gorm:"column:enrollment_date;type:varchar(50)" json:"enrollment_date"`
	GraduateDate        string    `gorm:"column:graduate_date;type:varchar(50)" json:"graduate_date"`
	CreateAt            time.Time `gorm:"column:create_at;autoCreateTime;not null;comment:创建时间" json:"create_at"`
	UpdateAt            time.Time `gorm:"column:update_at;autoUpdateTime;not null;comment:更新时间" json:"update_at"`
	Username            string    `gorm:"column:username;size:50;not null;uniqueIndex;comment:登录用户名(唯一)" json:"username" validate:"required,min=3,max=50"`
	AccountStatus       string    `gorm:"column:account_status;type:ENUM('active', 'locked', 'disabled');not null;default:'active';comment:账号状态" json:"account_status" validate:"required,oneof=active locked disabled"`
	Nickname            string    `gorm:"column:nickname;size:50;not null;comment:用户昵称" json:"nickname" validate:"required,min=1,max=50"`
	AvatarPath          string    `gorm:"column:avatar_path;size:255;not null;default:'./static/avatar/default-avatar.png';comment:头像路径" json:"avatar_path"`
	BackgroundImagePath string    `gorm:"column:backimage_path;size:255;not null;default:'./static/backgrounds/default-image.png';comment:背景路径" json:"backimage_path"`
	Signature           string    `gorm:"column:signature;type:varchar(255);comment:'个性签名'" json:"signature"`
	Tags                string    `gorm:"column:tags;type:varchar(255);comment:'标签'" json:"tags"`
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}

// BeforeCreate 钩子 - 创建前设置默认值
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.AvatarPath == "" {
		u.AvatarPath = "default-avatar.png"
	}
	return nil
}

func FindUserByName(c *gin.Context,username string) (*User, error) {
	user := &User{}
	err := global.Db.WithContext(c).Model(user).Where("username=?", username).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err // 其他数据库错误
	}
	return user, nil // 用户存在
}
func FindUserById(c *gin.Context,id uint) (*User, error) {
	user := &User{}
	err := global.Db.WithContext(c).Model(user).Where("id=?", id).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err // 其他数据库错误
	}
	return user, nil // 用户存在
}

func InsertUser(c *gin.Context,user *User) (id uint, err error) {
	result := global.Db.WithContext(c).Model(user).Create(user) // 通过指针传递数据
	if result.Error != nil {
		// 处理错误
		global.Logger.Warnf("创建记录失败: %v", result.Error)
		return 0, result.Error
	}
	global.Logger.Infof("插入成功，ID: %d\n", user.ID)
	return user.ID, nil
}

func UpdateUser(c *gin.Context,newuser *User, id uint) error {
	global.Logger.Infof("Nickname:%v", newuser.Nickname)
	return global.Db.WithContext(c).Model(newuser).Where("id=?", id).Updates(*newuser).Error
}
