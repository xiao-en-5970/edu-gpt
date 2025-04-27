package model

import (
	"errors"
	"time"

	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"gorm.io/gorm"
)

// Post 对应数据库中的 post 表
type Post struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"帖子ID"`
	PosterID     uint      `gorm:"column:poster_id" json:"poster_id" comment:"发帖人id"`
	Title        string    `gorm:"column:title;type:varchar(200);index:idx_title_prefix(10)" json:"title" comment:"标题"`
	Content      string    `gorm:"column:content;type:text" json:"content" comment:"内容（除标题）"`
	ViewCount    int       `gorm:"column:view_count;default:0" json:"view_count" comment:"浏览数"`
	LikeCount    int       `gorm:"column:like_count;default:0" json:"like_count" comment:"点赞数"`
	CollectCount int       `gorm:"column:collect_count;default:0" json:"collect_count" comment:"收藏数"`
	CommentCount int       `gorm:"column:comment_count;default:0" json:"comment_count" comment:"(被）评论数"`
	CreateAt     time.Time `gorm:"column:create_at;autoCreateTime;not null" json:"create_at" comment:"创建时间"`
	UpdateAt     time.Time `gorm:"column:update_at;autoUpdateTime;not null" json:"update_at" comment:"更新时间"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "post"
}

func FindPostById(id uint) (*Post, error) {
	post := &Post{}
	err := global.Db.Model(post).Where("id=?", id).First(post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err // 其他数据库错误
	}
	return post, nil // 用户存在
}

func InsertPost(post *Post) (id uint, err error) {
	result := global.Db.Model(post).Create(post) // 通过指针传递数据
	if result.Error != nil {
		// 处理错误
		global.Logger.Warnf("创建记录失败: %v", result.Error)
		return 0, result.Error
	}
	global.Logger.Infof("插入成功，ID: %d\n", post.ID)
	return post.ID, nil
}

func UpdatePost(newpost *Post, id uint) error {
	return global.Db.Model(newpost).Where("id=?", id).Updates(*newpost).Error
}
