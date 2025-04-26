package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)


func HandlerPost(c *gin.Context){
	
	idstr := c.Param("id")
	pid, err := strconv.Atoi(idstr)
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
	id := uint(pid)
	post,err:=model.FindPostById(id)
	if post == nil{
		responce.ErrorInternalServerErrorWithCode(c,codes.CodePostNotExist)
		return
	}
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	u,ex:=c.Get("id")
	if !ex{
		responce.ErrorInternalServerErrorWithCode(c,codes.CodeAuthNotExistError)
		return 
	}
	uid := u.(uint)
	user,err := model.FindUserById(uid)
	if user == nil{
		responce.ErrorInternalServerErrorWithCode(c,codes.CodeUserNotExist)
		return
	}
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	resp := types.PostResp{
		Post: *post,
		Nickname: user.Nickname,
		Campus: user.Campus,
		Grade: user.Grade,
		Department: user.Department,
	}
	responce.SuccessWithData(c,resp)
}