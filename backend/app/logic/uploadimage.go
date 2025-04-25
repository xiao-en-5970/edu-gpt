package logic

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUploadImage(c *gin.Context,req *types.UploadImageReq)(resp *types.UploadImageResp,code int,err error){
	u,ex:=c.Get("username")
	if !ex{
		return &types.UploadImageResp{},codes.CodeAuthUnvalidToken,nil
	}
	username := u.(string)
	user,_:=model.FindUserByName(username)
	if user!=nil{
		//用户存在
		//生成存储路径
		absPath := fmt.Sprintf("%s/%s",global.Cfg.Image.RootPath, req.File.Filename)
		//删除旧头像（如果存在）
		if user.AvatarPath == absPath{
			os.RemoveAll(absPath)
		}else{
			user.AvatarPath = absPath
			model.UpdateUser(user,user.ID)
		}
		// 保存新文件
		if err := c.SaveUploadedFile(req.File, absPath); err != nil {
			return &types.UploadImageResp{},codes.CodeAllIntervalError,err
		}
		url := fmt.Sprintf("%s:%d/api/v1/user/imageurl/%d",global.Cfg.Server.Address,global.Cfg.Server.Port,user.ID)
		return &types.UploadImageResp{
			Url:url,
		},codes.CodeAllSuccess,nil
	}else{
		return &types.UploadImageResp{},codes.CodeUserNotExist,nil
	}
}