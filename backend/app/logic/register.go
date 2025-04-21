package logic

import (


	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/model"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
)

func LogicRegister(c *gin.Context, req *types.RegisterReq) (resp *types.RegisterResp, code int, err error) {
	u,_:=model.FindUserByName(req.Username)
	if u!=nil{
		//用户存在
		return &types.RegisterResp{}, codes.CodeUserAlreadyExist, nil
	}
	user, err := model.RegisterReqToUser(req)
	if err != nil {
		return &types.RegisterResp{}, codes.CodeAllIntervalError, err
	}

	err = model.InsertUser(user)
	if err != nil {
		return &types.RegisterResp{}, codes.CodeAllIntervalError, err
	}
	return &types.RegisterResp{}, codes.CodeUserRegisterSuccess, nil
}
