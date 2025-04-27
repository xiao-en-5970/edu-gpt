package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerUserUploadBackImage(c *gin.Context) {
	file, err := c.FormFile("backimage")
	if err != nil {
		global.Logger.Error(err)
		responce.ErrorBadRequest(c, err)
		return
	}
	if !isImageFile(file) {

		responce.ErrorBadRequest(c, errors.New(codes.CodeMsg[codes.CodeImageFormatError]))
		return
	}
	global.Logger.Infof("接收图片成功")
	resp, code, err := logic.LogicUserUploadBackImage(c, &types.UploadImageReq{File: file})
	if err != nil {
		global.Logger.Error(err)
		responce.ErrorInternalServerError(c, err)
		return
	}
	if code == codes.CodeAllSuccess {
		responce.SuccessWithData(c, resp)
	} else {
		responce.ErrorInternalServerErrorWithCode(c, code)
	}
}
