package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostCreate(c *gin.Context, req *types.CreatePostReq) (resp types.CreatePostResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	post := &models.Post{
		PosterID:     uid,
		Title:        req.Title,
		Content:      req.Content,
		ViewCount:    0,
		LikeCount:    0,
		CollectCount: 0,
		CommentCount: 0,
	}
	pid, err := models.CreatePost(c, post)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.ID = pid
	return resp, codes.CodeAllSuccess, nil

}
