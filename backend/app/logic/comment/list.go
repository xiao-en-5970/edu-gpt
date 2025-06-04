package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCommentList(c *gin.Context, req *types.CommentListReq) (resp types.CommentListResp, code int, err error) {
	_, ex := c.Get("id")
	if !ex {
		return types.CommentListResp{}, codes.CodeAuthUnvalidToken, nil
	}
	comments, err := model.ListComment(c, req.PostID, req.Page, req.Size,req.Desc,req.Order)
	if err != nil {
		return types.CommentListResp{}, codes.CodeAllIntervalError, err
	}
	global.Logger.Info(comments)
	briefcomments := make([]types.BriefComment, 0, 1)
	for _, co := range comments {
		poster, err := model.FindUserById(c, co.UserID)
		if poster == nil {
			return types.CommentListResp{}, codes.CodeUserNotExist, nil
		}
		if err != nil {
			return types.CommentListResp{}, codes.CodeAllIntervalError, err
		}
		briefcomments = append(briefcomments, types.BriefComment{
			Content:    co.Content,
			Nickname:   poster.Nickname,
			ID:         co.ID,
			PostID:     co.PostID,
			PosterID:   co.UserID,
			Active:     co.Active,
			LikeCount:  co.LikeCount,
			ChildCount: co.ChildCount,
			CreateAt:   co.CreateAt,
			AvatarUrl:  global.GetUrl("user/auth/avatar", co.UserID),
		})
	}
	return briefcomments, codes.CodeAllSuccess, nil
}
func LogicSubCommentList(c *gin.Context, req *types.SubCommentListReq) (resp types.SubCommentListResp, code int, err error) {
	_, ex := c.Get("id")
	if !ex {
		return types.SubCommentListResp{}, codes.CodeAuthUnvalidToken, nil
	}
	comments, err := model.ListSubComment(c, req.ParentCommentID, req.Page, req.Size,req.Desc,req.Order)
	if err != nil {
		return types.SubCommentListResp{}, codes.CodeAllIntervalError, err
	}
	briefcomments := make([]types.BriefSubComment, 0, 1)
	for _, co := range comments {
		poster, err := model.FindUserById(c, co.UserID)
		if poster == nil {
			return types.SubCommentListResp{}, codes.CodeUserNotExist, nil
		}
		if err != nil {
			return types.SubCommentListResp{}, codes.CodeAllIntervalError, err
		}
		briefcomments = append(briefcomments, types.BriefSubComment{
			Content:         co.Content,
			Nickname:        poster.Nickname,
			ID:              co.ID,
			Post_id:         co.PostID,
			PosterID:        co.UserID,
			Active:          co.Active,
			LikeCount:       co.LikeCount,
			ParentCommentID: co.ParentCommentID,
			ReplyUserID:     co.ReplyUserID,
			CreateAt:        co.CreateAt,
			AvatarUrl:       global.GetUrl("user/auth/avatar", co.UserID),
		})
	}
	return briefcomments, codes.CodeAllSuccess, nil
}
