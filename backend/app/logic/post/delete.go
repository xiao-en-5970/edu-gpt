package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostDelete(c *gin.Context,req *types.PostDeleteReq)(resp *types.PostDeleteResp,code int,err error){
	if err = model.ChangePostStatus(c,req.PostID,0);err!=nil{
		return &types.PostDeleteResp{},codes.CodeAllIntervalError,err
	}
	return &types.PostDeleteResp{OK: 1},codes.CodeAllSuccess,nil
}