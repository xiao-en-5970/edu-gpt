package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/xiao-en-5970/Goodminton/backend/app/global"
	"github.com/xiao-en-5970/Goodminton/backend/app/model"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/auth"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/bcrypts"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
)

func LogicLogin(c *gin.Context,req *types.LoginReq)(resp *types.LoginResp,code int,err error){
	user,_:=model.FindUserByName(req.Username)
	if user!=nil{
		//用户存在
		if (bcrypts.CheckPasswordHash(req.Password,user.PasswordHash)){
			//账密正确
			// 生成Token
			token, err := auth.GenerateToken(req.Username)
			if err != nil {
				return &types.LoginResp{},codes.CodeAllIntervalError,err
			}

			// 设置HTTP-Only Cookie
			c.SetCookie("auth_token", token, int(global.Cfg.Auth.MaxAge), "/", "", false, true) // 24小时过期

			return &types.LoginResp{Token: token},codes.CodeUserLoginSuccess,nil
		}else{
			//账密错误
			global.Logger.Errorf("%#v %#v",user.PasswordHash,req.Password)
			return &types.LoginResp{},codes.CodeUserLoginPasswordError,nil
		}
	}else{
		// 用户不存在
		code,err:=LogicLoginHFUT(&types.LoginHFUTReq{
			Username: req.Username,
			Password: req.Password,
		})
		return &types.LoginResp{},code,err
	}
}
