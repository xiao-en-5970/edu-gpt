package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicGetUserInfo(c *gin.Context,req *types.GetUserInfoReq)(resp *types.GetUserInfoResp,code int,err error){
	u,ex:=c.Get("username")
	if !ex{
		return &types.GetUserInfoResp{},codes.CodeAuthUnvalidToken,nil
	}
	username := u.(string)
	user,_:=model.FindUserByName(username)
	if user!=nil{
		//用户存在
		hfutrsp,code,err:=LogicHFUTStudentInfo(c,username)
		if code!=codes.CodeAllSuccess{
			return &types.GetUserInfoResp{},code,nil
		}
		if err!=nil{
			return &types.GetUserInfoResp{},code,err
		}

		return &types.GetUserInfoResp{
			HFUTStudentInfo: hfutrsp.Data,
			User:*user,
		},codes.CodeAllSuccess,nil
	}else{
		
		return &types.GetUserInfoResp{},codes.CodeUserNotExist,nil
	}
}