package logic

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserUpdateUserInfo(c *gin.Context,user *model.User)(resp *types.GetUserInfoResp,code int,err error){
	u,ex:=c.Get("id")
	if !ex{
		return &types.GetUserInfoResp{},codes.CodeAuthUnvalidToken,nil
	}
	id := u.(uint)
	us,_:=model.FindUserById(id)
	if us!=nil{
		//用户存在
		global.Logger.Debug(user)
		if err=model.UpdateUser(user,us.ID);err!=nil{
			return &types.GetUserInfoResp{},codes.CodeUserInfoUpdateFail,err
		}
		us,_=model.FindUserById(id)
		var tag = make([]string,0)
		err:=json.Unmarshal([]byte(us.Tags),&tag)
		if err !=nil{
			return &types.GetUserInfoResp{},codes.CodeAllIntervalError,err
		}
		return &types.GetUserInfoResp{
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
			CreatedAt:      us.CreatedAt,
			Username:       us.Username,
			AccountStatus:  us.AccountStatus,
			Nickname:       us.Nickname,
			AvatarUrl:     fmt.Sprintf("%s/api/v1/user/auth/imageurl/%d",global.Cfg.Server.Address,user.ID),
			Signature:      us.Signature,
			Tags:           tag,
		},codes.CodeAllSuccess,nil
	}else{
		return &types.GetUserInfoResp{},codes.CodeUserNotExist,nil
	}
}