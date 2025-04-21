package logic

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/auth"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicLogin(c *gin.Context,req *types.LoginReq)(resp *types.LoginResp,code int,err error){
	hfutresp,code,err:=LogicHFUTLogin(&types.HFUTLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	cookie:=hfutresp.Data.Cookie
	global.RedisClient.SetEx(c,req.Username,cookie,time.Duration(global.Cfg.Redis.CookieExpire)*time.Hour)
	switch code {
	case codes.CodeAllSuccess:
		user,_:=model.FindUserByName(req.Username)
		if user==nil{
			// 用户不存在
			hfutrsp,code,err:=LogicHFUTStudentInfo(c,req.Username)
			if err!=nil{
				return &types.LoginResp{},code,err
			}
			model.InsertUser(&model.User{
				Username: req.Username,
				Nickname: hfutrsp.Data.UsernameZh,
			})
		}
		// 生成Token
		token, err := auth.GenerateToken(req.Username)
		if err != nil {
			return &types.LoginResp{},codes.CodeAllIntervalError,err
		}
		return &types.LoginResp{Token: token},codes.CodeAllSuccess,nil
	default:
		return &types.LoginResp{},code,err
	}
}
