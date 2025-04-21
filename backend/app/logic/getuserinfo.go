package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicGetUserInfo(c *gin.Context,req *types.GetUserInfoReq)(resp *types.UserInfo,code int,err error){
	u,ex:=c.Get("username")
	if !ex{
		return &types.UserInfo{},codes.CodeAuthUnvalidToken,nil
	}
	username := u.(string)
	user,_:=model.FindUserByName(username)
	if user!=nil{
		//用户存在
		userinfo,err:=model.ConvertUserToUserInfo(user)
		if err!=nil{
			return &types.UserInfo{},codes.CodeAllIntervalError,err
		}
		return userinfo,codes.CodeUserLoginSuccess,nil
	}else{
		
		return &types.UserInfo{},codes.CodeUserNotExist,nil
	}
}