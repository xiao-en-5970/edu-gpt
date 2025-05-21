package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCommentCreate(c *gin.Context, req *types.CommentCreateReq) (resp *types.CommentCreateResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return &types.CommentCreateResp{}, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	post, err := model.FindPostById(c, req.PostID)
	if err != nil {
		return &types.CommentCreateResp{}, codes.CodeAllIntervalError, err
	}
	if post == nil {
		return &types.CommentCreateResp{}, codes.CodePostNotExist, err
	}
	comment := &model.Comment{
		PostID:  req.PostID,
		UserID:  uid,
		Content: req.Content,
	}
	cid, err := model.CreateComment(c, comment)
	if err != nil {
		return &types.CommentCreateResp{}, codes.CodeAllIntervalError, err
	}
	return &types.CommentCreateResp{ID: cid}, codes.CodeAllSuccess, nil
}

func LogicSubCommentCreate(c *gin.Context, req *types.SubCommentCreateReq) (resp *types.SubCommentCreateResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return &types.SubCommentCreateResp{}, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	post, err := model.FindCommentById(c, req.PostID)
	if err != nil {
		return &types.SubCommentCreateResp{}, codes.CodeAllIntervalError, err
	}
	if post == nil {
		return &types.SubCommentCreateResp{}, codes.CodePostNotExist, nil
	}
	comment, err := model.FindCommentById(c, req.ParentCommentID)
	if err != nil {
		return &types.SubCommentCreateResp{}, codes.CodeAllIntervalError, err
	}
	if comment == nil {
		return &types.SubCommentCreateResp{}, codes.CodeCommentNotExist, nil
	}
	if req.ReplyUserID != 0 {
		user, err := model.FindUserById(c, req.ReplyUserID)
		if err != nil {
			return &types.SubCommentCreateResp{}, codes.CodeAllIntervalError, err
		}
		if user == nil {
			return &types.SubCommentCreateResp{}, codes.CodeUserNotExist, nil
		}
	}
	subcomment := &model.SubComment{
		PostID:          req.PostID,
		UserID:          uid,
		ParentCommentID: req.ParentCommentID,
		ReplyUserID:     req.ReplyUserID,
		Content:         req.Content,
	}
	cid, err := model.CreateSubComment(c, subcomment)
	if err != nil {
		return &types.SubCommentCreateResp{}, codes.CodeAllIntervalError, err
	}
	return &types.SubCommentCreateResp{ID: cid}, codes.CodeAllSuccess, nil
}
