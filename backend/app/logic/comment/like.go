package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCommentLike(c *gin.Context,req *types.CommentLikeReq)(resp *types.CommentLikeResp,code int,err error){
	u, ex := c.Get("id")
	if !ex {
		return &types.CommentLikeResp{}, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	comment,_ := model.FindCommentById(c,req.CommentID)
	if comment==nil || comment.Active!=global.Active.String(){
		return &types.CommentLikeResp{},codes.CodeCommentNotExist,nil
	}
	go model.AddCommentLikeCount(c,req.CommentID,uid,req.LikeStatus)
	return &types.CommentLikeResp{OK: 1}, codes.CodeAllSuccess, nil
}

func LogicSubCommentLike(c *gin.Context,req *types.SubCommentLikeReq)(resp *types.SubCommentLikeResp,code int,err error){
	u, ex := c.Get("id")
	if !ex {
		return &types.SubCommentLikeResp{}, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	subcomment,_ := model.FindSubCommentById(c,req.SubCommentID)
	if subcomment==nil || subcomment.Active!=global.Active.String(){
		return &types.SubCommentLikeResp{},codes.CodeSubCommentNotExist,nil
	}
	go model.AddSubCommentLikeCount(c,req.SubCommentID,uid,req.LikeStatus)
	return &types.SubCommentLikeResp{OK: 1}, codes.CodeAllSuccess, nil
}