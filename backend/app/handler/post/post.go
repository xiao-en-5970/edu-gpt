package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)


func HandlerPost(c *gin.Context){
	
	idstr := c.Param("id")
	pidint, err := strconv.Atoi(idstr)
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
	pid := uint(pidint)
	post,err:=model.FindPostById(c,pid)
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
	user,err := model.FindUserById(c,uid)
	if user == nil{
		responce.ErrorInternalServerErrorWithCode(c,codes.CodeUserNotExist)
		return
	}
	if err!=nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	images,err := model.FindPostImageByPid(c,pid)
	if images == nil{
		responce.ErrorInternalServerErrorWithCode(c,codes.CodeImageNotExist)
		return
	}
	urls:=make([]string,0,1)
	for _,i:=range(images){
		urls = append(urls,GetUrl("postimage",i.ID))
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
		ImageUrls: urls,
	}
	responce.SuccessWithData(c,resp)
}
func GetUrl(prefix string, id uint) string {
	return fmt.Sprintf("https://%s/api/v1/post/auth/%s/%d", global.Cfg.Server.Address, prefix, id)
}