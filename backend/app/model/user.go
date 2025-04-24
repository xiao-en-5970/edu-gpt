package model

import (
	"errors"
	"time"

	"github.com/xiao-en-5970/edu-gpt/backend/app/global"

	"gorm.io/gorm"
)

type AccountStatus string



type User struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id"`
	Password       string    `gorm:"type:varchar(255)" json:"-"`
    UsernameEn     string    `gorm:"type:varchar(100)" json:"usernameEn"`
    UsernameZh     string    `gorm:"type:varchar(100)" json:"usernameZh"`
    Sex            string    `gorm:"type:enum('男','女','其他')" json:"sex"`
    CultivateType  string    `gorm:"type:varchar(50)" json:"cultivateType"`
    Department     string    `gorm:"type:varchar(100)" json:"department"`
    Grade          string    `gorm:"type:varchar(20)" json:"grade"`
    Level          string    `gorm:"type:varchar(50)" json:"level"`
    StudentType    string    `gorm:"type:varchar(50)" json:"studentType"`
    Major          string    `gorm:"type:varchar(100)" json:"major"`
    Class          string    `gorm:"type:varchar(50)" json:"class"`
    Campus         string    `gorm:"type:varchar(50)" json:"campus"`
    Status         string    `gorm:"type:varchar(50)" json:"status"`
    Length         string   `gorm:"type:decimal(3,1)" json:"length"`
    EnrollmentDate string `gorm:"type:date" json:"enrollmentDate"`
    GraduateDate   string `gorm:"type:date" json:"graduateDate"`
	CreatedAt     time.Time `gorm:"autoCreateTime;not null;comment:创建时间" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime;not null;comment:更新时间" json:"updated_at"`
	Username      string    `gorm:"size:50;not null;uniqueIndex;comment:登录用户名(唯一)" json:"username" validate:"required,min=3,max=50"`
	AccountStatus string    `gorm:"type:ENUM('active', 'locked', 'disabled');not null;default:'active';comment:账号状态" json:"account_status" validate:"required,oneof=active locked disabled"`
	Nickname      string    `gorm:"size:50;not null;comment:用户昵称" json:"nickname" validate:"required,min=1,max=50"`
	AvatarPath    string    `gorm:"size:255;not null;default:'default-avatar.png';comment:头像路径" json:"avatar_path"`
	Signature     string    `gorm:"type:varchar(255);comment:'个性签名'" json:"signature"`
	Tags 		  string    `gorm:"type:varchar(255);comment:'标签'" json:"tags"`
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


func FindUserByName(username string)(*User,error){
	user:=&User{}
	err:=global.Db.Model(user).Where("username=?",username).First(user).Error
	if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // 用户不存在
        }
        return nil, err // 其他数据库错误
    }
    return user, nil // 用户存在
}

func InsertUser(user *User)(err error){
	result := global.Db.Model(user).Create(user) // 通过指针传递数据
	if result.Error != nil {
		// 处理错误
		global.Logger.Warnf("创建记录失败: %v", result.Error)
		return result.Error
	}
	global.Logger.Infof("插入成功，ID: %d\n", user.ID)
	return nil
}

func UpdateUser(newuser * User,id uint)(error){
	global.Logger.Infof("Nickname:%v",newuser.Nickname)
    return global.Db.Model(newuser).Where("id=?",id).Updates(*newuser).Error
}
