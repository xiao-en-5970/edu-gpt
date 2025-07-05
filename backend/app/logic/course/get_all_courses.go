package logic

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/HFUT"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/HFUT"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicGetAllCourses(c *gin.Context,semester int) (code int, err error) {
	sem, err := models.CheckSemester(c, semester)
	if err != nil {
		return codes.CodeAllIntervalError, err
	}
	if sem != nil {
		return codes.CodeCourseAlreadyExist, nil
	}
	CourseInfos := make(chan types.CourseInfo, 100)
	Lookup := make(map[string]bool)
	u, ex := c.Get("id")
	if !ex {
		return codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	user, err := models.FindUserById(c, uid)
	if user == nil {
		return codes.CodeUserNotExist, nil
	}
	if err != nil {
		return codes.CodeAllIntervalError, err
	}
	wg_recv := sync.WaitGroup{}
	for i := 1; i <= 500; i++ {
		time.Sleep(50 * time.Millisecond) // 避免请求过快
		wg_recv.Add(1)
		go func(i int) {
			defer wg_recv.Done()
			hfutresp, code, err := logic.LogicHFUTCourses(c, user.Username, "", i, 10, semester)
			if err != nil {
				return
			}
			if code != codes.CodeAllSuccess {
				return 
			}
			for _, r := range hfutresp.Data.List {
				if strings.HasPrefix(r.CourseType, "公选") || strings.HasPrefix(r.CourseType, "慕课") {
					// 忽略公选课和慕课
					global.Logger.Debugf("忽略课程：%s %s", r.CourseCode, r.CourseName)
					continue
				}
				if !Lookup[r.CourseCode] {
					Lookup[r.CourseCode] = true
					CourseInfos <- types.CourseInfo{
						CourseName: r.CourseName,
						CourseCode: r.CourseCode,
						CourseType: r.CourseType,
						Credits:    r.Credits,
						OpenDepart: r.OpenDepart,
						ExamMod:    r.ExamMod,
						Campus:     r.Campus,
					}
				}
			}
			global.Logger.Infof("获取第%d页课程信息成功，共获取%d条课程信息", i, len(hfutresp.Data.List))
		}(i)
	}
	wg_chan := sync.WaitGroup{}
	wg_chan.Add(1)
	go func(){
		defer wg_chan.Done()
		for ci := range CourseInfos {
			
			id, err := models.InsertCourse(c, semester, &models.Course{
				CourseName: ci.CourseName,
				CourseCode: ci.CourseCode,
				CourseType: ci.CourseType,
				Credits:    ci.Credits,
				OpenDepart: ci.OpenDepart,
				ExamMod:    ci.ExamMod,
				Campus:     ci.Campus,
			})
			if id == 0 {
				return 
			}
			if err != nil {
				return
			}
			os.MkdirAll(global.Cfg.Static.BookPath+"/"+ci.CourseType+"/"+ci.CourseName, 0755)
		}
	}()
	wg_recv.Wait()
	close(CourseInfos)
	wg_chan.Wait()
	return codes.CodeAllSuccess, nil
}
