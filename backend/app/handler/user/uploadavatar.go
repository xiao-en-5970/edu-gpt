package handler

import (
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

// @Summary      更新头像
// @Description  通过form-data上传头像
// @Tags         User模块
// @Security     BearerAuth 
// @Accept       multipart/form-data
// @Produce      json
// @Param        avatar formData file true "图片文件"
// @Success      200  {object}  types.UploadImageResp  "成功响应"
// @Router       /user/auth/upload_avatar [post]
func HandlerUserUploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
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
	resp, code, err := logic.LogicUserUploadAvatar(c, &types.UploadImageReq{File: file})
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

// 辅助函数：检查是否为图片
func isImageFile(file *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/jpg":  true,
	}
	return allowedTypes[file.Header.Get("Content-Type")]
}
