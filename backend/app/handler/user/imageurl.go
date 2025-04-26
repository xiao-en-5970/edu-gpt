package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerUserImageUrl(c *gin.Context){
	idstr := c.Param("id")
	uid,err := strconv.Atoi(idstr)
	id := uint(uid)
	if err != nil{
		responce.ErrorBadRequest(c,err)
		return
	}
	user,err:=model.FindUserById(id)
	if user == nil{
		responce.ErrorInternalServerErrorWithCode(c,codes.CodeUserNotExist)
		return
	}
	if err != nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	c.File(user.AvatarPath)
}