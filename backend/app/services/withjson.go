package services

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

type Service interface {
	NewReq() any
	Logic(c *gin.Context, req any) (resp any, code int, err error)
}

func ServiceHandlerWithJson(c *gin.Context, service Service) {
	req := service.NewReq()
	raw, ex := c.Get("raw_req")
	var err error
	if ex {
		err = json.Unmarshal(raw.([]byte), req)
	} else {
		err = c.ShouldBindJSON(req)
	}
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
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
