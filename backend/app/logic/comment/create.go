package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCommentCreate(c *gin.Context, req *types.CommentCreateReq) (resp types.CommentCreateResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	post, err := models.FindPostById(c, req.PostID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if post == nil {
		return resp, codes.CodePostNotExist, err
	}
	comment := &models.Comment{
		PostID:  req.PostID,
		UserID:  uid,
		Content: req.Content,
	}
	cid, err := models.CreateComment(c, comment)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.ID = cid
	return resp, codes.CodeAllSuccess, nil
}

func LogicSubCommentCreate(c *gin.Context, req *types.SubCommentCreateReq) (resp types.SubCommentCreateResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	post, err := models.FindCommentById(c, req.PostID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if post == nil {
		return resp, codes.CodePostNotExist, nil
	}
	comment, err := models.FindCommentById(c, req.ParentCommentID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if comment == nil {
		return resp, codes.CodeCommentNotExist, nil
	}
	if req.ReplyUserID != 0 {
		user, err := models.FindUserById(c, req.ReplyUserID)
		if err != nil {
			return resp, codes.CodeAllIntervalError, err
		}
		if user == nil {
			return resp, codes.CodeUserNotExist, nil
		}
	}
	subcomment := &models.SubComment{
		PostID:          req.PostID,
		UserID:          uid,
		ParentCommentID: req.ParentCommentID,
		ReplyUserID:     req.ReplyUserID,
		Content:         req.Content,
	}
	cid, err := models.CreateSubComment(c, subcomment)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.ID = cid
	return resp, codes.CodeAllSuccess, nil
}
