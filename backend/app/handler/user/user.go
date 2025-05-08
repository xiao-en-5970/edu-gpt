package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerUser(c *gin.Context) {
	idstr := c.Param("id")
	uid, err := strconv.Atoi(idstr)
	id := uint(uid)
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
	user, err := model.FindUserById(c,id)
	if user == nil {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodeUserNotExist)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	rsp := &types.BriefUser{
		ID:         user.ID,
		CreateAt:  user.CreateAt,
		Department: user.Department,
		Nickname:   user.Nickname,
		AvatarUrl: global.GetUrl("user/auth/avatar",user.ID),
		BackImageUrl: global.GetUrl("user/auth/backimage",user.ID),
		Sex:        user.Sex,
		Grade:      user.Grade,
		Campus:     user.Campus,
		Signature:  user.Signature,
		Tags:       user.Tags,
	}
	responce.SuccessWithData(c, rsp)
}
