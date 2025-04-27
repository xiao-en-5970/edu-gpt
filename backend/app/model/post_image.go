package model

import (
	"errors"
	"time"

	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"gorm.io/gorm"
)

type PostImage struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement;comment:帖子图片ID"`
	PostID     uint      `gorm:"column:post_id;comment:发帖人id"`
	Number     int       `gorm:"column:number;comment:第几张图片"`
	ImagesPath string    `gorm:"column:images_path;type:varchar(255);default:'';comment:图片路径"`
	CreateAt   time.Time `gorm:"column:create_at;autoCreateTime;not null;comment:创建时间"`
	UpdateAt   time.Time `gorm:"column:update_at;autoUpdateTime;not null;comment:更新时间"`
}

func (PostImage) TableName() string {
	return "post_image"
}

func FindPostImageByPidNum(pid uint, num int) (*PostImage, error) {
	postimage := &PostImage{}
	err := global.Db.Model(postimage).Where("post_id=? AND number=?", pid, num).First(postimage).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err // 其他数据库错误
	}
	return postimage, nil // 用户存在
}
func FindPostImageByPid(pid uint) ([]PostImage, error) {
	postimage := make([]PostImage,0,1)
	err := global.Db.Model(postimage).Where("post_id=?", pid,).Find(&postimage).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err // 其他数据库错误
	}
	return postimage, nil // 用户存在
}

func FindPostImageById(id uint) (*PostImage, error) {
	postimage := &PostImage{}
	err := global.Db.Model(postimage).Where("id=?", id).First(postimage).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 图片不存在
		}
		return nil, err // 其他数据库错误
	}
	return postimage, nil // 用户存在
}

func InsertPostImage(postimage *PostImage) (id uint, err error) {
	result := global.Db.Model(postimage).Create(postimage) // 通过指针传递数据
	if result.Error != nil {
		// 处理错误
		global.Logger.Warnf("创建记录失败: %v", result.Error)
		return 0, result.Error
	}
	global.Logger.Infof("插入成功，ID: %d\n", postimage.ID)
	return postimage.ID, nil
}
