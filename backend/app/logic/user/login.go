package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/auth"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/bcrypts"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserLogin(c *gin.Context, req *types.LoginReq) (resp *types.LoginResp, code int, err error) {
	user, _ := model.FindUserByName(c, req.Username)
	var id uint = 0
	if user == nil {
		// 用户不存在
		uid, code, err := LogicUserCreate(c, req)
		if err != nil {
			return &types.LoginResp{}, codes.CodeAllIntervalError, err
		}
		if code != codes.CodeAllSuccess {
			return &types.LoginResp{}, code, nil
		}
		id = uid
	} else {
		// 用户存在
		ok := bcrypts.CheckPasswordHash(req.Password, user.Password)
		if !ok {
			return &types.LoginResp{}, codes.CodeUserLoginPasswordError, nil
		}
		id = user.ID
		go LogicUserRefreshHFUTInfo(c, req.Username, req.Password, user.ID)
	}
	// 生成Token
	token, err := auth.GenerateToken(id)
	if err != nil {
		return &types.LoginResp{}, codes.CodeAllIntervalError, err
	}
	return &types.LoginResp{Token: token, ID: id}, codes.CodeAllSuccess, nil
}
