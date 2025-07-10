package models

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"gorm.io/gorm"
)

// CommentTable 评论区表
type CommentTable struct {
	ID      uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Active  string         `gorm:"type:enum('active','locked','disabled');default:'active'" json:"active"`
	CreateAt time.Time     `gorm:"autoCreateTime" json:"create_at"`
	UpdateAt time.Time     `gorm:"autoUpdateTime" json:"update_at"`
}

// TableName 设置表名
func (CommentTable) TableName() string {
	return "comment_table"
}

func CommentTableCreate(c *gin.Context)(id uint,err error){
	ct := &CommentTable{}
	err = global.Db.WithContext(c).Model(ct).Create(ct).Error
	if err !=nil{
		return 0,err
	}
	return ct.ID,nil
}

func FindCommentTableById(c *gin.Context,id uint)(*CommentTable,error){
	ct := &CommentTable{}
	err := global.Db.WithContext(c).Model(ct).Where("id=?", id).First(ct).Error
	if err !=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,err
		}
		return nil,err
	}
	return ct,nil
}