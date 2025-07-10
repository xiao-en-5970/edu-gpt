package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
)

// Community 社区表
type Community struct {
	ID          uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"type:varchar(200);not null" json:"name"`
	Description string         `gorm:"type:text;not null" json:"description"`
	Active      string         `gorm:"type:enum('active','locked','disabled');default:'active'" json:"active"`
	CreateAt    time.Time     `gorm:"autoCreateTime" json:"create_at"`
	UpdateAt    time.Time     `gorm:"autoUpdateTime" json:"update_at"`
}

// TableName 设置表名
func (Community) TableName() string {
	return "community"
}


func CummunityCreate(c *gin.Context,community *Community)(id uint,err error){
	
	err = global.Db.WithContext(c).Model(community).Create(community).Error
	if err !=nil{
		return 0,err
	}
	return community.ID,nil
}