package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

// @Summary      获取用户数据
// @Description  根据用户ID返回用户
// @Tags         User模块
// @Security     BearerAuth
// @Produce      json
// @Param        id  path  string  true  "用户ID"
// @Success 200 {object} types.BriefUser "成功响应"
// @Router       /user/auth/{id} [post]
func HandlerUser(c *gin.Context) {
	idstr := c.Param("id")
	uid, err := strconv.Atoi(idstr)
	id := uint(uid)
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
	user, err := model.FindUserById(c, id)
	if user == nil {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodeUserNotExist)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	var tag = make([]string, 0)
	err = json.Unmarshal([]byte(user.Tags), &tag)
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	rsp := &types.BriefUser{
		ID:           user.ID,
		CreateAt:     user.CreateAt,
		Department:   user.Department,
		Nickname:     user.Nickname,
		AvatarUrl:    global.GetUrl("user/auth/avatar", user.ID),
		BackImageUrl: global.GetUrl("user/auth/backimage", user.ID),
		Sex:          user.Sex,
		Grade:        user.Grade,
		Campus:       user.Campus,
		Signature:    user.Signature,
		Tags:         tag,
		Follows:      user.Follows,
		Fans:         user.Fans,
		AllPostLike:  user.AllPostLike,
	}
	responce.SuccessWithData(c, rsp)
}
