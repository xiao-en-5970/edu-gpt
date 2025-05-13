package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)
// @Summary 帖子图片查看
// @Description 通过帖子id直接查看帖子
// @Tags Post模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param        id  path  string  true  "帖子图片ID"  
// @Success      200  {file}  binary  "帖子图片文件"
// @Router /post/auth/postimage/{id} [get,post]
func HandlerPostPostImage(c*gin.Context){
	idstr := c.Param("id")
	uid,err := strconv.Atoi(idstr)
	id := uint(uid)
	if err != nil{
		responce.ErrorBadRequest(c,err)
		return
	}
	image,err:=model.FindPostImageById(c,id)
	if image == nil{
		responce.ErrorInternalServerErrorWithCode(c,codes.CodeImageNotExist)
		return
	}
	if err != nil{
		responce.ErrorInternalServerError(c,err)
		return
	}
	c.File(image.ImagesPath)
}