package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostEdit(c *gin.Context,req *types.EditPostReq)(resp * types.EditPostResp,code int,err error){
	if req.ID==0{
		return &types.EditPostResp{},codes.CodeAllRequestFormatError,nil
	}
	u,ex:=c.Get("id")
	if !ex{
		return &types.EditPostResp{},codes.CodeAuthUnvalidToken,nil
	}
	uid := u.(uint)
	user,err := model.FindUserById(c,uid)
	if user == nil{
		return &types.EditPostResp{},codes.CodeUserNotExist,nil
	}
	post := &model.Post{
		ID: req.ID,
		Title: req.Title,
		Content: req.Content,
	}
	if err!=nil{
		return &types.EditPostResp{},codes.CodeAllIntervalError,err
	}
	err=model.UpdatePost(c,post,req.ID)
	if err!=nil{
		return &types.EditPostResp{},codes.CodeAllIntervalError,err
	}
	rsppost,err:=model.FindPostById(c,req.ID)
	if rsppost==nil{
		return &types.EditPostResp{},codes.CodePostNotExist,err
	}
	if err !=nil{
		return &types.EditPostResp{},codes.CodeAllIntervalError,err
	}
	return &types.EditPostResp{
		ID: req.ID,
	},codes.CodeAllSuccess,nil
}
