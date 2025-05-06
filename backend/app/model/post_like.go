package model

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"gorm.io/gorm"
)

type PostLike struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    uint      `gorm:"not null;index:uk_post_user,unique;comment:帖子ID" json:"post_id"`
	UserID    uint      `gorm:"not null;index:idx_user;comment:用户ID" json:"user_id"`
	Status    int       `gorm:"not null;default:1;comment:1-点赞 0-取消" json:"status"`
	CreatedAt time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdatedAt time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
}

// TableName 设置表名
func (PostLike) TableName() string {
	return "post_likes"
}


func AddLikeCount(c *gin.Context,postid uint,userid uint,expectlikestatus int)(newlikeCount int,newlikeStatus int,err error){
	err=global.Db.Transaction(func(tx *gorm.DB)(err error){
		oldlike := &PostLike{}
		post:=&Post{}
		oldlike.PostID = postid
		oldlike.UserID = userid
		err=tx.WithContext(c).Model(oldlike).Where("post_id=? and user_id=?",postid,userid).First(oldlike).Error
		if err!=nil{
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if expectlikestatus != 0{
					oldlike.Status = 1
					err=tx.WithContext(c).Model(oldlike).Where("post_id=? and user_id=?",postid,userid).Create(oldlike).Error
					if err!=nil{
						return err
					}
					err=tx.WithContext(c).Model(post).Where("id=?",postid).Update("like_count", gorm.Expr("like_count + 1")).Error
					if err!=nil{
						return err
					}
				}
			}else{
				return err
			}
		}else{
			oldstatus := oldlike.Status
			if oldstatus != expectlikestatus{
				err=tx.WithContext(c).Model(oldlike).Where("id=?",oldlike.ID).Update("status",expectlikestatus).Error
				if err !=nil{
					return err
				}
				err=tx.WithContext(c).Model(post).Where("id=?",postid).Update("like_count", gorm.Expr("like_count + ?",expectlikestatus-oldstatus)).Error
				if err!=nil{
					return err
				}
			}
		}
		err=tx.WithContext(c).Model(post).Where("id=?",postid).First(post).Error
		if err!=nil{
			return err
		}
		global.Logger.Infoln(post.LikeCount)
		newlikeCount = post.LikeCount
		return nil
	})
	if err !=nil{
		return -1,0,err
	}
	return newlikeCount,expectlikestatus,nil
}