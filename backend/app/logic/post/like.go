package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
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
	post,_ := model.FindPostById(c,req.PostID)
	if post==nil || post.Active!=global.Active.String(){
		return &types.PostLikeResp{},codes.CodePostNotExist,nil
	}
	go model.AddLikeCount(c, req.PostID, uid, req.LikeStatus)
	return &types.PostLikeResp{OK: 1}, codes.CodeAllSuccess, nil
}
