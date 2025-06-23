package logic

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserGetUserInfo(c *gin.Context, req *types.GetUserInfoReq) (resp types.GetUserInfoResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	//用户存在
	user, err := models.FindUserById(c, uid)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	var tag = make([]string, 0)
	err = json.Unmarshal([]byte(user.Tags), &tag)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	return types.GetUserInfoResp{
		ID:             user.ID,
		UsernameZh:     user.UsernameZh,
		Sex:            user.Sex,
		CultivateType:  user.CultivateType,
		Department:     user.Department,
		Grade:          user.Grade,
		Level:          user.Level,
		Major:          user.Major,
		Class:          user.Class,
		Campus:         user.Campus,
		EnrollmentDate: user.EnrollmentDate,
		GraduateDate:   user.GraduateDate,
		CreateAt:       user.CreateAt,
		Username:       user.Username,
		AccountStatus:  user.Active,
		Nickname:       user.Nickname,
		AvatarUrl:      global.GetUrl("user/auth/avatar", user.ID),
		BackImageUrl:   global.GetUrl("user/auth/backimage", user.ID),
		Signature:      user.Signature,
		Tags:           tag,
		Follows:        user.Follows,
		Fans:           user.Fans,
		AllPostLike:    user.AllPostLike,
	}, codes.CodeAllSuccess, nil
}
