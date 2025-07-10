package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/redisprefix"
	"gorm.io/gorm"
)

type PostLike struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PostID    uint      `gorm:"column:post_id;not null;index:uk_post_user,unique;comment:帖子ID" json:"post_id"`
	UserID    uint      `gorm:"column:user_id;not null;index:idx_user;comment:用户ID" json:"user_id"`
	Status    int       `gorm:"column:status;not null;default:1;comment:1-点赞 0-取消" json:"status"`
	CreatedAt time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdatedAt time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_at"`
}

// TableName 设置表名
func (PostLike) TableName() string {
	return "post_likes"
}

func GetLikeStatus(c *gin.Context, postid uint, userid uint) (likeStatus int, err error) {
	pre1 := middleware.GetPrefix(redisprefix.PrefixPostLikeKey1,strconv.Itoa(int(postid)))
	pre2 := middleware.GetPrefix(redisprefix.PrefixPostLikeKey2,strconv.Itoa(int(userid)))
	pre:=middleware.GetPrefix(pre1,pre2)
	result := global.RedisClient.Get(c,pre)
	if result.Err() != nil {
		if result.Err() == redis.Nil {
			return 0, nil
		}
		return 0, result.Err()
	}
	likeStatus, err = strconv.Atoi(result.Val())
	if err != nil {
		return 0, err
	}
	return likeStatus, nil
}

func UpdateUserLikePost(c *gin.Context, postid uint, userid uint, expectlikestatus int) (err error) {
	err = global.Db.Transaction(func(tx *gorm.DB) (err error) {
		oldlike := &PostLike{}
		oldlike.PostID = postid
		oldlike.UserID = userid
		err = tx.WithContext(c).Model(oldlike).Where("post_id=? and user_id=?", postid, userid).First(oldlike).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if expectlikestatus != 0 {
					oldlike.Status = 1
					err = tx.WithContext(c).Model(oldlike).Where("post_id=? and user_id=?", postid, userid).Create(oldlike).Error
					if err != nil {
						return err
					}
					err = middleware.UpdateCacheCount(c,
						postid,
						redisprefix.PrefixPostLikeCountKey,
						1,
						GetPostLikeCountFromMysql,
						UpdatePostLikeCountFromMysql,
					)
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
				err := middleware.UpdateCacheCount(c,
					postid, 
					redisprefix.PrefixPostLikeCountKey,
					expectlikestatus-oldstatus,
					GetPostLikeCountFromMysql,
					UpdatePostLikeCountFromMysql)
				if err != nil {
					return err
				}
			}else{
				global.Logger.Debugf("点赞状态不变 %d",expectlikestatus)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	pre1 := middleware.GetPrefix(redisprefix.PrefixPostLikeKey1,strconv.Itoa(int(postid)))
	pre2 := middleware.GetPrefix(redisprefix.PrefixPostLikeKey2,strconv.Itoa(int(userid)))
	pre := middleware.GetPrefix(pre1,pre2)
	result := global.RedisClient.SetEx(c,
		pre,
		expectlikestatus,
		time.Duration(global.Cfg.Redis.LikeExpire)*time.Hour)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func GetPostLikeCountFromMysql(c *gin.Context, id uint) (count int, err error) {
	post := &Post{}
	err = global.Db.WithContext(c).Model(post).Where("id=?", id).First(post).Error
	if err != nil {
		return -1, err
	}
	return post.LikeCount, nil
}

func UpdatePostLikeCountFromMysql(c *gin.Context, id uint, newcount int) (err error) {
	post := &Post{}
	err = global.Db.WithContext(c).Model(post).Where("id=?", id).Update("like_count", newcount).Error
	if err != nil {
		return err
	}
	return nil
}
