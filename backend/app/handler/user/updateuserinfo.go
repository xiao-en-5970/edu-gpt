package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/logic/user"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

// @Summary 更改用户信息
// @Description 更改用户信息
// @Tags User模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param update_userinfo body types.UpdateUserInfoReq true "请求体"
// @Success 200 {object} types.UpdateUserInfoResp "成功响应"
// @Router /user/auth/update_userinfo [post]
func HandlerUserUpdateUserInfo(c *gin.Context) {
	req := &types.UpdateUserInfoReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		responce.ErrorBadRequest(c,err)
		return
	}
	resp, code, err := logic.LogicUserUpdateUserInfo(c, req)
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	if code == codes.CodeAllSuccess {
		responce.SuccessWithData(c, *resp)
	} else {
		responce.ErrorInternalServerErrorWithCode(c, code)
	}
}


