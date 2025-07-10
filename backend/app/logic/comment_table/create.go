package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/comment_table"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCommentTableCreate(c *gin.Context,req types.CommentTableCreateReq)(resp types.CommentTableCreateResp,code int,err error){
	id,err :=models.CommentTableCreate(c)
	if err !=nil{
		return resp,codes.CodeAllIntervalError,err
	}
	resp.ID = id 
	return resp,codes.CodeAllSuccess,err
}