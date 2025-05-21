package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/redisprefix"
	"gorm.io/gorm"
)

type CommentLike struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CommentID uint      `gorm:"column:comment_id;not null;index:uk_comment_user,unique;comment:评论ID" json:"comment_id"`
	UserID    uint      `gorm:"column:user_id;not null;index:idx_user;comment:用户ID" json:"user_id"`
	Status    int       `gorm:"column:status;not null;default:1;comment:1-点赞 0-取消" json:"status"`
	CreatedAt time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdatedAt time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
}

type SubCommentLike struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubCommentID uint      `gorm:"column:subcomment_id;not null;index:uk_subcomment_user,unique;comment:子评论ID" json:"subcomment_id"`
	UserID       uint      `gorm:"column:user_id;not null;index:idx_user;comment:用户ID" json:"user_id"`
	Status       int       `gorm:"column:status;not null;default:1;comment:1-点赞 0-取消" json:"status"`
	CreatedAt    time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdatedAt    time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
}

// TableName 设置表名
func (SubCommentLike) TableName() string {
	return "subcomment_likes"
}

// TableName 设置表名
func (CommentLike) TableName() string {
	return "comment_likes"
}
func AddCommentLikeCount(c *gin.Context, commentid uint, userid uint, expectlikestatus int) (err error) {
	err = global.Db.Transaction(func(tx *gorm.DB) (err error) {
		oldlike := &CommentLike{}
		comment := &Comment{}
		oldlike.CommentID = commentid
		oldlike.UserID = userid
		err = tx.WithContext(c).Model(oldlike).Where("comment_id=? and user_id=?", commentid, userid).First(oldlike).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if expectlikestatus != 0 {
					oldlike.Status = 1
					err = tx.WithContext(c).Model(oldlike).Where("comment_id=? and user_id=?", commentid, userid).Create(oldlike).Error
					if err != nil {
						return err
					}
					err = tx.WithContext(c).Model(comment).Where("id=?", commentid).Update("like_count", gorm.Expr("like_count + 1")).Error
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			oldstatus := oldlike.Status
			if oldstatus != expectlikestatus {
				err = tx.WithContext(c).Model(oldlike).Where("id=?", oldlike.ID).Update("status", expectlikestatus).Error
				if err != nil {
					return err
				}
				err = tx.WithContext(c).Model(comment).Where("id=?", commentid).Update("like_count", gorm.Expr("like_count + ?", expectlikestatus-oldstatus)).Error
				if err != nil {
					return err
				}
			}
		}
		err = tx.WithContext(c).Model(comment).Where("id=?", commentid).First(comment).Error
		if err != nil {
			return err
		}
		global.Logger.Infoln(comment.LikeCount)
		return nil
	})
	if err != nil {
		return err
	}
	result := global.RedisClient.SetEx(c,
		redisprefix.PrefixCommentLikeKey1+strconv.Itoa(int(commentid))+redisprefix.PrefixCommentLikeKey2+strconv.Itoa(int(userid)),
		strconv.Itoa(int(expectlikestatus)),
		time.Duration(global.Cfg.Redis.LikeExpire)*time.Hour)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}


func AddSubCommentLikeCount(c *gin.Context, subcommentid uint, userid uint, expectlikestatus int) (err error) {
	err = global.Db.Transaction(func(tx *gorm.DB) (err error) {
		oldlike := &SubCommentLike{}
		subcomment := &SubComment{}
		oldlike.SubCommentID = subcommentid
		oldlike.UserID = userid
		err = tx.WithContext(c).Model(oldlike).Where("subcomment_id=? and user_id=?", subcommentid, userid).First(oldlike).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if expectlikestatus != 0 {
					oldlike.Status = 1
					err = tx.WithContext(c).Model(oldlike).Where("subcomment_id=? and user_id=?", subcommentid, userid).Create(oldlike).Error
					if err != nil {
						return err
					}
					err = tx.WithContext(c).Model(subcomment).Where("id=?", subcommentid).Update("like_count", gorm.Expr("like_count + 1")).Error
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		} else {
			oldstatus := oldlike.Status
			if oldstatus != expectlikestatus {
				err = tx.WithContext(c).Model(oldlike).Where("id=?", oldlike.ID).Update("status", expectlikestatus).Error
				if err != nil {
					return err
				}
				err = tx.WithContext(c).Model(subcomment).Where("id=?", subcommentid).Update("like_count", gorm.Expr("like_count + ?", expectlikestatus-oldstatus)).Error
				if err != nil {
					return err
				}
			}
		}
		err = tx.WithContext(c).Model(subcomment).Where("id=?", subcommentid).First(subcomment).Error
		if err != nil {
			return err
		}
		global.Logger.Infoln(subcomment.LikeCount)
		return nil
	})
	if err != nil {
		return err
	}
	result := global.RedisClient.SetEx(c,
		redisprefix.PrefixSubCommentLikeKey1+strconv.Itoa(int(subcommentid))+redisprefix.PrefixSubCommentLikeKey2+strconv.Itoa(int(userid)),
		strconv.Itoa(int(expectlikestatus)),
		time.Duration(global.Cfg.Redis.LikeExpire)*time.Hour)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}