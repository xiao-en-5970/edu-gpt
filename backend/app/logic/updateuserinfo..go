package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUpdateUserInfo(c *gin.Context,user *model.User)(resp *types.GetUserInfoResp,code int,err error){
	u,ex:=c.Get("username")
	if !ex{
		return &types.GetUserInfoResp{},codes.CodeAuthUnvalidToken,nil
	}
	if user.ID==0{
		return &types.GetUserInfoResp{},codes.CodeAllRequestFormatError,err
	}
	username := u.(string)
	us,_:=model.FindUserByName(username)
	if us!=nil{
		//用户存在
		global.Logger.Debug(user)
		if err=model.UpdateUser(user,us.ID);err!=nil{
			return &types.GetUserInfoResp{},codes.CodeUserInfoUpdateFail,err
		}
		hfutrsp,code,err:=LogicHFUTStudentInfo(c,username)
		if err!=nil{
			return &types.GetUserInfoResp{},code,err
		}
		us,_=model.FindUserByName(username)
		return &types.GetUserInfoResp{HFUTStudentInfo: hfutrsp.Data,User: *us},codes.CodeAllSuccess,nil
	}else{
		return &types.GetUserInfoResp{},codes.CodeUserNotExist,nil
	}
}