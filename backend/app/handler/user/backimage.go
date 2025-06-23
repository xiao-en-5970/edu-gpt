package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

// @Summary      获取用户背景图
// @Description  根据用户ID返回背景图数据
// @Tags         User模块
// @Security     BearerAuth
// @Produce      octet-stream
// @Param        id  path  string  true  "用户ID"
// @Success      200  {file}  binary  "背景图片文件"
// @Router       /user/auth/backimage/{id} [get]
func HandlerUserBackImage(c *gin.Context) {
	idstr := c.Param("id")
	uid, err := strconv.Atoi(idstr)
	id := uint(uid)
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
	user, err := models.FindUserById(c, id)
	if user == nil {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodeUserNotExist)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	c.File(user.BackgroundImagePath)
}
