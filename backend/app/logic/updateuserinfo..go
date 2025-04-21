package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/model"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
)

func LogicUpdateUserInfo(c *gin.Context,userinfo *types.UserInfo)(resp *types.UserInfo,code int,err error){
	u,ex:=c.Get("username")
	if !ex{
		return &types.UserInfo{},codes.CodeAuthUnvalidToken,nil
	}
	username := u.(string)
	us,_:=model.FindUserByName(username)
	if us!=nil{
		//用户存在
		user,err:=model.ConvertUserInfoToUser(userinfo)
		if err!=nil{
			return &types.UserInfo{},codes.CodeAllIntervalError,err
		}
		if err:=model.UpdateUser(user);err!=nil{
			return &types.UserInfo{},codes.CodeUserInfoUpdateFail,err
		}
		return userinfo,codes.CodeUserLoginSuccess,nil
	}else{
		
		return &types.UserInfo{},codes.CodeUserNotExist,nil
	}
}