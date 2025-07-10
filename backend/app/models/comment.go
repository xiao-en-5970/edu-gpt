package models

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"gorm.io/gorm"
)

type Comment struct {
	ID              uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	CommentTableID  uint     `gorm:"not null" json:"comment_table_id"`                                       // 评论的评论区ID
	UserID          uint     `gorm:"not null" json:"user_id"`                                                // 发评论的用户ID
	ParentCommentID uint     `gorm:"not null;default:0" json:"parent_comment_id"`                            // 上层评论id，为0表示评论区顶层
	ReplyCommentID  uint     `gorm:"not null;default:0" json:"reply_comment_id"`                             // 回复的评论id，为0表示评论区顶层
	Content         string    `gorm:"type:text;not null" json:"content"`                                      // 内容
	LikeCount       int       `gorm:"not null;default:0" json:"like_count"`                                   // 点赞数
	ChildCount      int       `gorm:"not null;default:0" json:"child_count"`                                  // 子评论数量
	Active          string    `gorm:"type:enum('active','locked','disabled');default:'active'" json:"active"` // 状态
	CreateAt        time.Time `gorm:"autoCreateTime" json:"create_at"`                                        // 创建时间
	UpdateAt        time.Time `gorm:"autoUpdateTime" json:"update_at"`                                        // 更新时间
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

func CreateComment(c *gin.Context, comment *Comment) (id uint, err error) {
	err = global.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.WithContext(c).Model(comment).Create(comment).Error
		if err != nil {
			global.Logger.Warnf("创建记录失败: %v", err)
			return err
		}
		err = tx.WithContext(c).Model(&Post{}).Where("id=?", comment.CommentTableID).Update("comment_count", gorm.Expr("comment_count + 1")).Error
		if err != nil {
			global.Logger.Warnf("评论数增加失败: %v", err)
			return err
		}
		err = tx.WithContext(c).Model(&Comment{}).Where("id=?", comment.ParentCommentID).Update("child_count", gorm.Expr("child_count + 1")).Error
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
	return comment.ID, nil
}

func ListComment(c *gin.Context, ctid uint, page int, size int, desc int, order string) (comments []Comment, err error) {
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
	err = global.Db.WithContext(c).Model(&Comment{}).Where("active=? and comment_table_id=?", "active", ctid).Order(orderdesc).Offset((page - 1) * size).Limit(size).Find(&comments).Error
	return comments, err
}
