package logic

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	hfuttypes "github.com/xiao-en-5970/edu-gpt/backend/app/types/HFUT"
	logichfut "github.com/xiao-en-5970/edu-gpt/backend/app/logic/HFUT"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/bcrypts"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserCreate(c *gin.Context, req *types.LoginReq)(id uint,code int,err error){
	hfutloginresp, code, err := logichfut.LogicUserHFUTLogin(&hfuttypes.HFUTLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if code != codes.CodeAllSuccess {
		return 0, code, err
	}
	cookie := hfutloginresp.Data.Cookie
	global.RedisClient.SetEx(c, middleware.GetPrefix("username", req.Username), cookie, time.Duration(global.Cfg.Redis.CookieExpire)*time.Hour)
	hfutinforesp, code, err := logichfut.LogicHFUTStudentInfo(c, req.Username)
	if err != nil {
		return 0, codes.CodeAllIntervalError, err
	}
	if code != codes.CodeAllSuccess {
		return 0, code, nil
	}
	hashpass, err := bcrypts.HashPassword(req.Password)
	if err != nil {
		return 0, codes.CodeAllIntervalError, err
	}
	u := &model.User{
		Username:            req.Username,
		Nickname:            hfutinforesp.Data.UsernameZh,
		Password:            hashpass,
		UsernameEn:          hfutinforesp.Data.UsernameEn,
		UsernameZh:          hfutinforesp.Data.UsernameZh,
		Sex:                 hfutinforesp.Data.Sex,
		CultivateType:       hfutinforesp.Data.CultivateType,
		Department:          hfutinforesp.Data.Department,
		Grade:               hfutinforesp.Data.Grade,
		Level:               hfutinforesp.Data.Level,
		StudentType:         hfutinforesp.Data.StudentType,
		Major:               hfutinforesp.Data.Major,
		Class:               hfutinforesp.Data.Class,
		Campus:              hfutinforesp.Data.Campus,
		Status:              hfutinforesp.Data.Status,
		Length:              hfutinforesp.Data.Length,
		EnrollmentDate:      hfutinforesp.Data.EnrollmentDate,
		GraduateDate:        hfutinforesp.Data.GraduateDate,
		Tags:                "[]",
		Signature:           "这人啥也没说",
		AvatarPath:          "./static/avatars/default-avatar.png",
		BackgroundImagePath: "./static/backgrounds/default-image.png",
	}
	id, err = model.InsertUser(c, u)
	if err != nil {
		return 0, codes.CodeAllIntervalError, err
	}
	return id,codes.CodeAllSuccess,nil
}
