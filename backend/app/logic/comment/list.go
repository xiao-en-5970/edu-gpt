package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCommentList(c *gin.Context, req *types.CommentListReq) (resp types.CommentListResp, code int, err error) {
	_, ex := c.Get("id")
	if !ex {
		return types.CommentListResp{}, codes.CodeAuthUnvalidToken, nil
	}
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
	comments, err := models.ListComment(c, req.CommentTableID, req.Page, req.Size, req.Desc, req.Order)
	if err != nil {
		return types.CommentListResp{}, codes.CodeAllIntervalError, err
	}
	global.Logger.Info(comments)
	briefcomments := make([]types.BriefComment, 0, 1)
	for _, co := range comments {
		poster, err := models.FindUserById(c, co.UserID)
		if poster == nil {
			return types.CommentListResp{}, codes.CodeUserNotExist, nil
		}
		if err != nil {
			return types.CommentListResp{}, codes.CodeAllIntervalError, err
		}
		briefcomments = append(briefcomments, types.BriefComment{
			Content:        co.Content,
			Nickname:       poster.Nickname,
			ID:             co.ID,
			CommentTableID: co.CommentTableID,
			PosterID:       co.UserID,
			Active:         co.Active,
			LikeCount:      co.LikeCount,
			ChildCount:     co.ChildCount,
			CreateAt:       co.CreateAt,
			AvatarUrl:      global.GetUrl("user/auth/avatar", co.UserID),
		})
	}
	return briefcomments, codes.CodeAllSuccess, nil
}
