package models

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"gorm.io/gorm"
)

type Comment struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement;column:id" json:"id"`
	PostID     uint      `gorm:"column:post_id;not null;index:idx_post" json:"post_id"`
	UserID     uint      `gorm:"column:user_id;not null" json:"user_id"`
	Content    string    `gorm:"column:content;type:text" json:"content"`
	LikeCount  int       `gorm:"column:like_count;default:0" json:"like_count"`
	ChildCount int       `gorm:"column:child_count;default:0" json:"child_count"`
	Active     string    `gorm:"column:active;type:ENUM('active','locked','disabled');not null;default:'active'" json:"active"`
	CreateAt   time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt   time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
}

type SubComment struct {
	ID              uint      `gorm:"column:id;primaryKey;autoIncrement;column:id" json:"id"`
	PostID          uint      `gorm:"column:post_id;not null;index:idx_post" json:"post_id"`
	UserID          uint      `gorm:"column:user_id;not null" json:"user_id"`
	ParentCommentID uint      `gorm:"column:parent_comment_id;not null;default:0;index:idx_parent" json:"parent_comment_id"`
	ReplyUserID     uint      `gorm:"column:reply_user_id;not null;default:0" json:"reply_user_id"`
	Content         string    `gorm:"column:content;type:text" json:"content"`
	LikeCount       int       `gorm:"column:like_count;default:0" json:"like_count"`
	Active          string    `gorm:"column:active;type:ENUM('active','locked','disabled');not null;default:'active'" json:"active"`
	CreateAt        time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt        time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
}

// TableName 设置表名
func (SubComment) TableName() string {
	return "sub_comment"
}

// TableName 设置表名
func (Comment) TableName() string {
	return "comment"
}

func FindCommentById(c *gin.Context, id uint) (comment *Comment, err error) {
	comment = &Comment{}
	err = global.Db.WithContext(c).Model(comment).Where("id=?", id).First(comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err // 其他数据库错误
	}
	return comment, nil // 用户存在
}

func FindSubCommentById(c *gin.Context, id uint) (subcomment *SubComment, err error) {
	subcomment = &SubComment{}
	err = global.Db.WithContext(c).Model(subcomment).Where("id=?", id).First(subcomment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err // 其他数据库错误
	}
	return subcomment, nil // 用户存在
}

func CreateComment(c *gin.Context, comment *Comment) (id uint, err error) {
	err = global.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.WithContext(c).Model(comment).Create(comment).Error
		if err != nil {
			global.Logger.Warnf("创建记录失败: %v", err)
			return err
		}
		err = tx.WithContext(c).Model(&Post{}).Where("id=?", comment.PostID).Update("comment_count", gorm.Expr("comment_count + 1")).Error
		if err != nil {
			global.Logger.Warnf("评论数增加失败: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Logger.Warnf("创建失败: %v", err)
		return 0, err
	}
	return comment.ID, nil
}

func CreateSubComment(c *gin.Context, subcomment *SubComment) (id uint, err error) {
	err = global.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.WithContext(c).Model(subcomment).Create(subcomment).Error
		if err != nil {
			global.Logger.Warnf("创建记录失败: %v", err)
			return err
		}
		err = tx.WithContext(c).Model(&Post{}).Where("id=?", subcomment.PostID).Update("comment_count", gorm.Expr("comment_count + 1")).Error
		if err != nil {
			global.Logger.Warnf("评论数增加失败: %v", err)
			return err
		}
		err = tx.WithContext(c).Model(&Comment{}).Where("id=?", subcomment.ParentCommentID).Update("child_count", gorm.Expr("child_count + 1")).Error
		if err != nil {
			global.Logger.Warnf("回复数增加失败: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Logger.Warnf("创建失败: %v", err)
		return 0, err
	}
	return subcomment.ID, nil
}

func ListComment(c *gin.Context, pcid uint, page int, size int, desc int, order string) (comments []Comment, err error) {
	comments = make([]Comment, 0, size)
	orderdesc := ""
	if order == "time" {
		orderdesc += "id "
	} else if order == "like" {
		orderdesc += "like_count "
	}
	if desc == 0 {
		orderdesc += "ASC"
	} else {
		orderdesc += "DESC"
	}
	err = global.Db.WithContext(c).Model(&Comment{}).Where("active=? and post_id=?", "active",pcid).Order(orderdesc).Offset((page - 1) * size).Limit(size).Find(&comments).Error
	return comments, err
}

func ListSubComment(c *gin.Context, pcid uint, page int, size int, desc int, order string) (comments []SubComment, err error){
	comments = make([]SubComment, 0, size)
	orderdesc := ""
	if order == "time" {
		orderdesc += "id "
	} else if order == "like" {
		orderdesc += "like_count "
	}
	if desc == 0 {
		orderdesc += "ASC"
	} else {
		orderdesc += "DESC"
	}
	err = global.Db.WithContext(c).Model(&SubComment{}).Where("active=? and parent_comment_id=?", "active",pcid).Order(orderdesc).Offset((page - 1) * size).Limit(size).Find(&comments).Error
	return comments, err
}
