package logic

import (

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostCreate(c *gin.Context,req *types.CreatePostReq)(resp * types.CreatePostResp,code int,err error){
	u,ex:=c.Get("id")
	if !ex{
		return &types.CreatePostResp{},codes.CodeAuthUnvalidToken,nil
	}
	uid := u.(uint)
	user,err := model.FindUserById(uid)
	if user == nil{
		return &types.CreatePostResp{},codes.CodeUserNotExist,nil
	}
	post := &model.Post{
		PosterID: user.ID,
		Title: req.Title,
		Content: req.Content,
		ViewCount: 0,
		LikeCount: 0,
		CollectCount: 0,
		CommentCount: 0,
	}
	if err!=nil{
		return &types.CreatePostResp{},codes.CodeAllIntervalError,err
	}
	pid,err:=model.InsertPost(post)
	if err!=nil{
		return &types.CreatePostResp{},codes.CodeAllIntervalError,err
	}
	return &types.CreatePostResp{ID: pid},codes.CodeAllSuccess,nil

}