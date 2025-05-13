package handler

import (
	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
)
// @Summary 更改帖子信息【TODO】
// @Description 更改帖子信息【TODO】
// @Tags Post模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param edit body types.EditPostReq true "请求体"
// @Success 200 {object} types.EditPostResp "成功响应"
// @Router /post/auth/edit [post]
func HandlerPostEdit(c* gin.Context){
	services.ServiceHandlerWithJson(c,PostEdit{})
}

type PostEdit struct{}
func(PostEdit)NewReq() any {
    return &types.EditPostReq{}
}

func(PostEdit)Logic(c *gin.Context,req any)(resp any,code int,err error ){
    return logic.LogicPostEdit(c,req.(*types.EditPostReq))
}

