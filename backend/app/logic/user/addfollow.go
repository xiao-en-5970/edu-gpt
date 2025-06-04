package logic

import (

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)


func LogicUserAddFollow(c *gin.Context,req *types.UserAddFollowReq)(resp *types.UserAddFollowResp,code int,err error){
	u, ex := c.Get("id")
	if !ex {
		return &types.UserAddFollowResp{}, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	user,_ := model.FindUserById(c,req.FollowID)
	if user==nil || user.Active!=global.Active.String(){
		return &types.UserAddFollowResp{},codes.CodeUserNotExist,nil
	}
	go model.AddFollows(c,uid,req.FollowID,req.FollowStatus)
	return &types.UserAddFollowResp{OK: 1},codes.CodeAllSuccess,nil
}