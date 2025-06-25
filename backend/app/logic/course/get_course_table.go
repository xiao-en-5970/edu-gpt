package logic

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/HFUT"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicCourseGetTable(c *gin.Context, req *types.GetCourseTableReq) (resp types.GetCourseTableResp, code int, err error) {
	semesterId := global.YearTerm2SemesterId(req.Year, req.Term)
	if semesterId == 0 {
		return resp, codes.CodeCourseTableTermFault, nil
	}
	global.Logger.Infof("get course table, year: %d, term: %d, semesterId: %d", req.Year, req.Term, semesterId)
	if time.Now().Year() < req.Year {
		return resp, codes.CodeCourseTableYearFault, nil
	}
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	user, err := models.FindUserById(c, uid)
	if err != nil {
		return resp,codes.CodeAllIntervalError, err
	}
	if user == nil {
		return resp,codes.CodeUserNotExist, nil
	}
	courselistresp, code, err := logic.LogicHFUTCoursesList(c, user.Username, req.CampusId, semesterId)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	if code!=codes.CodeAllSuccess {
		return resp, code, nil
	}
	resp = types.GetCourseTableResp(courselistresp.Data)
	return resp, codes.CodeAllSuccess, nil
}
