package logic

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserUpdateUserInfo(c *gin.Context, req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return &types.UpdateUserInfoResp{}, codes.CodeAuthUnvalidToken, nil
	}
	id := u.(uint)
	us, _ := model.FindUserById(c,id)
	if us != nil {
		//用户存在
		global.Logger.Debug(req)
		us.AccountStatus = req.AccountStatus
		us.Nickname = req.Nickname
		us.Signature = req.Signature
		tagstr, err := json.Marshal(req.Tags)
		if err!=nil{
			return &types.UpdateUserInfoResp{}, codes.CodeAllIntervalError, err
		}
		us.Tags = string(tagstr)
		if err = model.UpdateUser(c,us, us.ID); err != nil {
			return &types.UpdateUserInfoResp{}, codes.CodeUserInfoUpdateFail, err
		}
		us, _ = model.FindUserById(c,id)
		var tag = make([]string, 0)
		err = json.Unmarshal([]byte(us.Tags), &tag)
		if err != nil {
			return &types.UpdateUserInfoResp{}, codes.CodeAllIntervalError, err
		}
		return &types.UpdateUserInfoResp{
			ID:             us.ID,
			UsernameZh:     us.UsernameZh,
			Sex:            us.Sex,
			CultivateType:  us.CultivateType,
			Department:     us.Department,
			Grade:          us.Grade,
			Level:          us.Level,
			Major:          us.Major,
			Class:          us.Class,
			Campus:         us.Campus,
			EnrollmentDate: us.EnrollmentDate,
			GraduateDate:   us.GraduateDate,
			CreateAt:       us.CreateAt,
			Username:       us.Username,
			AccountStatus:  us.AccountStatus,
			Nickname:       us.Nickname,
			AvatarUrl:      global.GetUrl("user/auth/avatar", us.ID),
			BackImageUrl:   global.GetUrl("user/auth/backimage", us.ID),
			Signature:      us.Signature,
			Tags:           tag,
		}, codes.CodeAllSuccess, nil
	} else {
		return &types.UpdateUserInfoResp{}, codes.CodeUserNotExist, nil
	}
}
