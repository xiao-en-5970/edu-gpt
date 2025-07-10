package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCommentLike(c *gin.Context, req *types.CommentLikeReq) (resp types.CommentLikeResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	comment, _ := models.FindCommentById(c, req.CommentID)
	if comment == nil || comment.Active != global.Active.String() {
		return resp, codes.CodeCommentNotExist, nil
	}
	go models.AddCommentLikeCount(c, req.CommentID, uid, req.LikeStatus)
	resp.OK = 1
	return resp, codes.CodeAllSuccess, nil
}



