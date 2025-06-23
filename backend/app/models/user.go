package models

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
	Active              string    `gorm:"column:active;type:ENUM('active', 'locked', 'disabled');not null;default:'active';comment:状态" json:"active" validate:"required,oneof=active locked disabled"`
	Nickname            string    `gorm:"column:nickname;size:50;not null;comment:用户昵称" json:"nickname" validate:"required,min=1,max=50"`
	AvatarPath          string    `gorm:"column:avatar_path;size:255;not null;default:'./static/avatar/default-avatar.png';comment:头像路径" json:"avatar_path"`
	BackgroundImagePath string    `gorm:"column:backimage_path;size:255;not null;default:'./static/backgrounds/default-image.png';comment:背景路径" json:"backimage_path"`
	Signature           string    `gorm:"column:signature;type:varchar(255);comment:'个性签名'" json:"signature"`
	Tags                string    `gorm:"column:tags;type:varchar(255);comment:'标签'" json:"tags"`
	Follows             int64     `gorm:"column:follow;not null;comment:用户关注数量"` // BIGINT 对应 int64
	Fans                int64     `gorm:"column:fans;not null;comment:用户粉丝数量"`
	AllPostLike         int64     `gorm:"column:allpost_like;not null;comment:用户点赞数量"`
}

type UserFollow struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   uint      `gorm:"column:user_id;not null" json:"user_id"`
	Follow   uint      `gorm:"column:follow;not null" json:"follow"`
	Status   int       `gorm:"column:status;not null;default:0" json:"status"`
	CreateAt time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_at"`
}

// TableName sets the table name for the UserFollow model
func (UserFollow) TableName() string {
	return "user_follow"
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

func FindUserByName(c *gin.Context, username string) (*User, error) {
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
func FindUserById(c *gin.Context, id uint) (*User, error) {
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

func InsertUser(c *gin.Context, user *User) (id uint, err error) {
	result := global.Db.WithContext(c).Model(user).Create(user) // 通过指针传递数据
	if result.Error != nil {
		// 处理错误
		global.Logger.Warnf("创建记录失败: %v", result.Error)
		return 0, result.Error
	}
	global.Logger.Infof("插入成功，ID: %d\n", user.ID)
	return user.ID, nil
}

func UpdateUser(c *gin.Context, newuser *User, id uint) error {
	global.Logger.Infof("Nickname:%v", newuser.Nickname)
	return global.Db.WithContext(c).Model(newuser).Where("id=?", id).Updates(*newuser).Error
}

func ChangeUserStatus(c *gin.Context, uid uint, statusCode int) (err error) {
	var s = global.Status(statusCode)
	user := &User{Active: s.String()}
	return global.Db.WithContext(c).Model(user).Where("id=?", uid).Updates(*user).Error
}

func AddFollows(c *gin.Context, uid uint, followuid uint, expectlikestatus int) (err error) {
	err = global.Db.Transaction(func(tx *gorm.DB) (err error) {
		oldflo := &UserFollow{}
		user := &User{}
		oldflo.Follow = followuid
		oldflo.UserID = uid
		err = tx.WithContext(c).Model(oldflo).Where("follow=? and user_id=?", followuid, uid).First(oldflo).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if expectlikestatus != 0 {
					oldflo.Status = 1
					err = tx.WithContext(c).Model(oldflo).Where("follow=? and user_id=?", followuid, uid).Create(oldflo).Error
					if err != nil {
						return err
					}
					err = tx.WithContext(c).Model(user).Where("id=?", uid).Update("follow", gorm.Expr("follow + 1")).Error
					if err != nil {
						return err
					}
					err = tx.WithContext(c).Model(user).Where("id=?", followuid).Update("fans", gorm.Expr("fans + 1")).Error
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			oldstatus := oldflo.Status
			if oldstatus != expectlikestatus {
				err = tx.WithContext(c).Model(oldflo).Where("id=?", oldflo.ID).Update("status", expectlikestatus).Error
				if err != nil {
					return err
				}
				err = tx.WithContext(c).Model(user).Where("id=?", uid).Update("follow", gorm.Expr("follow + ?", expectlikestatus-oldstatus)).Error
				if err != nil {
					return err
				}
				err = tx.WithContext(c).Model(user).Where("id=?", followuid).Update("fans", gorm.Expr("fans + ?", expectlikestatus-oldstatus)).Error
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

func FollowFansList(c *gin.Context, uid uint, page int, size int, desc int, order string, isfollow bool) (usersf []UserFollow, err error) {
	usersf = make([]UserFollow, 0, size)
	orderdesc := ""
	if order == "time" {
		orderdesc += "id "
	} else if order == "fans" {
		orderdesc += "fans "
	}
	if desc == 0 {
		orderdesc += "ASC"
	} else {
		orderdesc += "DESC"
	}
	if isfollow {
		err = global.Db.WithContext(c).Model(&UserFollow{}).Where("user_id=?", uid).Order(orderdesc).Offset((page - 1) * size).Limit(size).Find(&usersf).Error
	} else {
		err = global.Db.WithContext(c).Model(&UserFollow{}).Where("follow=?", uid).Order(orderdesc).Offset((page - 1) * size).Limit(size).Find(&usersf).Error
	}
	return usersf, err
}

func FindFollowStatusById(c *gin.Context,uid uint,follow uint)(status int,err error){
	uf:=&UserFollow{}
	err = global.Db.WithContext(c).Model(uf).Where("user_id=? and follow=?", uid,follow).First(uf).Error
	if err !=nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil // 用户不存在
		}
		return -1,err
	}
	return uf.Status,nil
}