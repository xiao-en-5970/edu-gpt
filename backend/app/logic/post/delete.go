package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostDelete(c *gin.Context, req *types.PostDeleteReq) (resp types.PostDeleteResp, code int, err error) {
	if err = models.ChangePostStatus(c, req.PostID, 0); err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.OK = 1
	return resp, codes.CodeAllSuccess, nil
}
