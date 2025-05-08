package logic

import (
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserUploadBackImage(c *gin.Context, req *types.UploadImageReq) (resp *types.UploadImageResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return &types.UploadImageResp{}, codes.CodeAuthUnvalidToken, nil
	}
	id := u.(uint)
	user, _ := model.FindUserById(c,id)
	if user != nil {
		//用户存在
		//生成存储路径
		absPath := fmt.Sprintf("%s/%d%s", global.Cfg.Image.BackImagePath, user.ID, path.Ext(req.File.Filename))
		//删除旧头像（如果存在）
		if user.BackgroundImagePath == absPath {
			os.RemoveAll(absPath)
		} else {
			user.BackgroundImagePath = absPath
			model.UpdateUser(c,user, user.ID)
		}
		// 保存新文件
		if err := c.SaveUploadedFile(req.File, absPath); err != nil {
			return &types.UploadImageResp{}, codes.CodeAllIntervalError, err
		}
		url := global.GetUrl("user/auth/backimage", user.ID)
		return &types.UploadImageResp{
			Url: url,
		}, codes.CodeAllSuccess, nil
	} else {
		return &types.UploadImageResp{}, codes.CodeUserNotExist, nil
	}
}