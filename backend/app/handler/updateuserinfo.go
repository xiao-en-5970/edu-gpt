package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/logic"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerUpdateUserInfo(c *gin.Context) {
	req := &model.User{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		responce.ErrorBadRequest(c,err)
		return
	}
	resp, code, err := logic.LogicUpdateUserInfo(c, req)
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	if code == codes.CodeAllSuccess {

		responce.SuccessWithCodeData(c, code, *resp)
	} else {
		responce.SuccessWithCode(c, code)
	}
}
