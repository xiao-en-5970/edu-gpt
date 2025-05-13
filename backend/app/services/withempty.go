package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func ServiceHandlerWithEmpty(c *gin.Context, service Service) {
	req := service.NewReq()
	resp, code, err := service.Logic(c, req)
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	if code == codes.CodeAllSuccess {
		responce.SuccessWithData(c, resp)
	} else {
		responce.ErrorInternalServerErrorWithCode(c, code)
	}
}