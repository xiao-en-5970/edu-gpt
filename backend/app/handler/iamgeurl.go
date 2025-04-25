package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerImageUrl(c *gin.Context){
	idstr := c.Param("id")
	id,err := strconv.Atoi(idstr)
	if err != nil{
		responce.ErrorBadRequest(c,err)
		return
	}
	user,err:=model.FindUserById(id)
	if user == nil{
		responce.SuccessWithCode(c,codes.CodeUserNotExist)
		return
	}
	if err != nil{
		responce.ErrorBadRequest(c,err)
		return
	}
	c.File(user.AvatarPath)
}