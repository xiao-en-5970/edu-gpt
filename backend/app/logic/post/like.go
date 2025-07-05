package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostLike(c *gin.Context, req *types.PostLikeReq) (resp types.PostLikeResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	post, _ := models.FindPostById(c, req.PostID)
	if post == nil || post.Active != global.Active.String() {
		return resp, codes.CodePostNotExist, nil
	}
	go models.UpdateUserLikePost(c, req.PostID, uid, req.LikeStatus)
	resp.OK = 1
	return resp, codes.CodeAllSuccess, nil
}
