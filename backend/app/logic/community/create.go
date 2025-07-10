package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/community"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)


func LogicCommunityCreate(c *gin.Context,req *types.CommunityCreateReq)(resp types.CommunityCreateResp,code int,err error){
	coid, err :=models.CummunityCreate(c,&models.Community{
		Name: req.Name,
		Description: req.Description,
	})
	if err !=nil{
		return resp,codes.CodeAllIntervalError,err
	}
	resp.ID = coid
	return resp,codes.CodeAllSuccess,nil
}