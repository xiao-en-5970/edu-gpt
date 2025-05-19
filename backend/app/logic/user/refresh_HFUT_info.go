package logic

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserRefreshHFUTInfo(c *gin.Context, username string, password string, uid uint) (code int, err error) {
	hfutloginresp, code, err := LogicUserHFUTLogin(&types.HFUTLoginReq{
		Username: username,
		Password: password,
	})
	if code != codes.CodeAllSuccess {
		return  code, err
	}
	cookie := hfutloginresp.Data.Cookie
	global.RedisClient.SetEx(c, middleware.GetPrefix("username", username), cookie, time.Duration(global.Cfg.Redis.CookieExpire)*time.Hour)
	hfutrsp, code, _ := LogicHFUTStudentInfo(c, username)
	if code == codes.CodeAllSuccess {
		u := &model.User{
			UsernameEn:     hfutrsp.Data.UsernameEn,
			UsernameZh:     hfutrsp.Data.UsernameZh,
			Sex:            hfutrsp.Data.Sex,
			CultivateType:  hfutrsp.Data.CultivateType,
			Department:     hfutrsp.Data.Department,
			Grade:          hfutrsp.Data.Grade,
			Level:          hfutrsp.Data.Level,
			StudentType:    hfutrsp.Data.StudentType,
			Major:          hfutrsp.Data.Major,
			Class:          hfutrsp.Data.Class,
			Campus:         hfutrsp.Data.Campus,
			Status:         hfutrsp.Data.Status,
			Length:         hfutrsp.Data.Length,
			EnrollmentDate: hfutrsp.Data.EnrollmentDate,
			GraduateDate:   hfutrsp.Data.GraduateDate,
		}
		err = model.UpdateUser(c, u, uid)
		if err != nil {
			return codes.CodeAllIntervalError, err
		}
	} else {
		return codes.CodeUserRefreshHFUTInfoFail, nil
	}
	return codes.CodeAllSuccess, nil
}
