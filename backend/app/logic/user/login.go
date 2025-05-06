package logic

import (
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/auth"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/crypt"
)

func LogicUserLogin(c *gin.Context, req *types.LoginReq) (resp *types.LoginResp, code int, err error) {
	hfutresp, code, err := LogicUserHFUTLogin(&types.HFUTLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	cookie := hfutresp.Data.Cookie
	global.RedisClient.SetEx(c, req.Username, cookie, time.Duration(global.Cfg.Redis.CookieExpire)*time.Hour)
	var id uint = 0
	switch code {
	case codes.CodeAllSuccess:
		user, _ := model.FindUserByName(c,req.Username)
		if user == nil {
			// 用户不存在
			hfutrsp, code, err := LogicHFUTStudentInfo(c, req.Username)
			if code != codes.CodeAllSuccess {
				return &types.LoginResp{}, code, nil
			}
			if err != nil {
				return &types.LoginResp{}, codes.CodeAllIntervalError, err
			}
			hashpass, err := crypt.HashPassword(req.Password, req.Username, "salt")
			if err != nil {
				return &types.LoginResp{}, codes.CodeAllIntervalError, err
			}
			u := &model.User{
				Username:            req.Username,
				Nickname:            hfutrsp.Data.UsernameZh,
				Password:            string(base64.StdEncoding.EncodeToString(hashpass)),
				UsernameEn:          hfutrsp.Data.UsernameEn,
				UsernameZh:          hfutrsp.Data.UsernameZh,
				Sex:                 hfutrsp.Data.Sex,
				CultivateType:       hfutrsp.Data.CultivateType,
				Department:          hfutrsp.Data.Department,
				Grade:               hfutrsp.Data.Grade,
				Level:               hfutrsp.Data.Level,
				StudentType:         hfutrsp.Data.StudentType,
				Major:               hfutrsp.Data.Major,
				Class:               hfutrsp.Data.Class,
				Campus:              hfutrsp.Data.Campus,
				Status:              hfutrsp.Data.Status,
				Length:              hfutrsp.Data.Length,
				EnrollmentDate:      hfutrsp.Data.EnrollmentDate,
				GraduateDate:        hfutrsp.Data.GraduateDate,
				Tags:                "[]",
				Signature:           "这人啥也没说",
				AvatarPath:          "./static/avatars/default-avatar.png",
				BackgroundImagePath: "./static/backgrounds/default-iamge.png",
			}
			id, err = model.InsertUser(c,u)
			if err != nil {
				return &types.LoginResp{}, codes.CodeAllIntervalError, err
			}
		} else {
			hfutrsp, code, _ := LogicHFUTStudentInfo(c, req.Username)
			if code != codes.CodeAllSuccess {
				return &types.LoginResp{}, code, nil
			}
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
			id = user.ID
			err = model.UpdateUser(c,u, user.ID)
			if err != nil {
				return &types.LoginResp{}, codes.CodeAllIntervalError, err
			}
		}
		// 生成Token
		token, err := auth.GenerateToken(id)
		if err != nil {
			return &types.LoginResp{}, codes.CodeAllIntervalError, err
		}
		return &types.LoginResp{Token: token, ID: id}, codes.CodeAllSuccess, nil
	default:
		return &types.LoginResp{}, code, err
	}
}
