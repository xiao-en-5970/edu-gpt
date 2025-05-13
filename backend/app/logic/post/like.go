package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostLike(c *gin.Context, req *types.PostLikeReq) (resp *types.PostLikeResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return &types.PostLikeResp{}, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	post, err := model.FindPostById(c, req.PostID)
	if err != nil {
		return &types.PostLikeResp{}, codes.CodeAllIntervalError, err
	}
	if post == nil {
		return &types.PostLikeResp{}, codes.CodePostNotExist, nil
	}
	count, status, err := model.AddLikeCount(c, req.PostID, uid, req.LikeStatus)
	if err != nil {
		return &types.PostLikeResp{}, codes.CodeAllIntervalError, err
	}
	return &types.PostLikeResp{LikeCount: count, LikeStatus: status}, codes.CodeAllSuccess, nil

}
