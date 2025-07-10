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
	if req.PostID != 0 {
		post, err := models.FindPostById(c, req.PostID)
		if err != nil {
			return resp, codes.CodeAllIntervalError, err
		}
		if post == nil {
			return resp, codes.CodePostNotExist, err
		}
		req.CommentTableID = post.CommentTableID
	}
	ct, err := models.FindCommentTableById(c, req.CommentTableID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if ct == nil {
		return resp, codes.CodeCommentTableNotFound, nil
	}

	comment := &models.Comment{
		CommentTableID: req.CommentTableID,
		UserID:         uid,
		Content:        req.Content,
	}
	cid, err := models.CreateComment(c, comment)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.ID = cid
	return resp, codes.CodeAllSuccess, nil
}

func LogicSubCommentCreate(c *gin.Context, req *types.SubCommentCreateReq) (resp types.CommentCreateResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	if req.PostID != 0 {
		post, err := models.FindPostById(c, req.PostID)
		if err != nil {
			return resp, codes.CodeAllIntervalError, err
		}
		if post == nil {
			return resp, codes.CodePostNotExist, err
		}
		req.CommentTableID = post.CommentTableID
	}
	ct, err := models.FindCommentTableById(c, req.CommentTableID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if ct == nil {
		return resp, codes.CodeCommentTableNotFound, nil
	}
	cmt, err := models.FindCommentById(c, req.ParentCommentID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if cmt == nil {
		return resp, codes.CodeCommentParentNotExist, nil
	}
	if cmt.CommentTableID != req.CommentTableID {
		return resp, codes.CodeCommentParentNotExistThisTable, nil
	}
	if cmt.ParentCommentID != 0 {
		return resp, codes.CodeCommentParentCantBeReply, nil
	}
	if req.ReplyID == 0 {
		req.ReplyID = req.ParentCommentID
	}

	comment := &models.Comment{
		CommentTableID:  ct.ID,
		ParentCommentID: req.ParentCommentID,
		ReplyCommentID:  req.ReplyID,
		UserID:          uid,
		Content:         req.Content,
	}
	cid, err := models.CreateComment(c, comment)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.ID = cid
	return resp, codes.CodeAllSuccess, nil
}
