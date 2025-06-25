package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/services"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/HFUT"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/redisprefix"
)

func LogicUserHFUTLogin(req *types.HFUTLoginReq) (resp types.HFUTLoginResp, code int, err error) {
	c := &http.Client{}
	params := &url.Values{}
	params.Add("username", req.Username)
	params.Add("password", req.Password)
	r, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/vpn/login?%s", global.Cfg.HfutAPI.Host, global.Cfg.HfutAPI.Port, params.Encode()), nil)
	rsp, err := c.Do(r)
	if err != nil {
		return resp, codes.CodeAllBadGateway, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode == 200 {
		//登录信息门户成功
		body, _ := io.ReadAll(rsp.Body)
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return resp, codes.CodeHFUTIntervalError, nil
		}
		return resp, codes.CodeAllSuccess, nil
	} else if rsp.StatusCode == 400 {
		//登录信息门户失败
		return resp, codes.CodeHFUTLoginError, nil
	} else if rsp.StatusCode == 500 {
		//登录信息门户限流
		for i := 0; i < global.Cfg.HfutAPI.Retry; i++ {
			time.Sleep(1000 * time.Millisecond)
			rsp, err := c.Do(r)
			if err != nil {
				return resp, codes.CodeAllBadGateway, err
			}
			if rsp.StatusCode != 500 {
				break
			}
		}
		if rsp.StatusCode == 200 {
			//登录信息门户成功
			body, _ := io.ReadAll(rsp.Body)
			err = json.Unmarshal(body, &resp)
			if err != nil {
				return resp, codes.CodeHFUTIntervalError, nil
			}
			return resp, codes.CodeAllSuccess, nil
		} else if rsp.StatusCode == 400 {
			//登录信息门户失败
			return resp, codes.CodeHFUTLoginError, nil
		} else {
			return resp, codes.CodeHFUTIntervalError, nil
		}

	} else {
		global.Logger.Infof("rsp.Body: %v\n")
		//登录信息门户未知问题
		return resp, codes.CodeHFUTUnkonwnError, nil
	}
}

func LogicHFUTStudentInfo(c *gin.Context, username string) (resp types.HFUTStudentInfoResp, code int, err error) {
	result := global.RedisClient.Get(c, middleware.GetPrefix(redisprefix.PrefixUserCookieKey, username))
	if result.Err() != nil {
		return resp, codes.CodeHFUTIntervalError, nil
	}
	cookie := result.Val()
	global.Logger.Infof("cookie:%s", cookie)
	client := &http.Client{}
	r, _ := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/vpn/eam/studentinfo?", global.Cfg.HfutAPI.Host, global.Cfg.HfutAPI.Port), nil)
	r.Header.Set("cookie", cookie)
	rsp, err := client.Do(r)
	if err != nil {
		return resp, codes.CodeAllBadGateway, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode == 200 {
		//获取信息成功
		bytebody, _ := io.ReadAll(rsp.Body)
		err := json.Unmarshal(bytebody, &resp)
		if err != nil {
			return resp, codes.CodeHFUTIntervalError, nil
		}
		return resp, codes.CodeAllSuccess, nil
	} else if rsp.StatusCode == 401 || rsp.StatusCode == 400 {
		// 删除 Redis 中的无效 cookie
		global.RedisClient.Del(c, middleware.GetPrefix(redisprefix.PrefixUserCookieKey, username))
		//未登录
		return resp, codes.CodeHFUTNotLogin, nil
	} else if rsp.StatusCode == 500 {
		//信息门户限流
		for i := 0; i < global.Cfg.HfutAPI.Retry; i++ {
			time.Sleep(100 * time.Millisecond)
			rsp, err := client.Do(r)
			if err != nil {
				return resp, codes.CodeAllBadGateway, err
			}
			if rsp.StatusCode != 500 {
				break
			}
		}
		return resp, codes.CodeHFUTIntervalError, nil
	} else {
		global.Logger.Infof("rsp.Body: %v\n", rsp.Body)
		//登录信息门户未知问题
		return resp, codes.CodeHFUTUnkonwnError, nil
	}
}

func LogicHFUTCourses(c *gin.Context, username string, courseName string, page int, size int, semesterCode int) (resp types.HFUTCoursesResp, code int, err error) {
	result := services.NewHFUTApi().WithContext(c).
		WithPrefix("vpn/eam/course/search").
		WithParmas(
			"username", username,
			"courseName", courseName,
			"page", strconv.Itoa(page),
			"size", strconv.Itoa(size),
			"semesterCode", strconv.Itoa(semesterCode),
		).
		GenNewRequest().RequestSetCookieByUsername(username).
		Do().Result()
	if result.Err() != nil {
		return resp, codes.CodeAllIntervalError, result.Err()
	}
	if result.Val() == nil {
		return resp, result.Code(), nil
	}
	err = json.Unmarshal(result.Val(), &resp)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	return resp, result.Code(), nil
}

func LogicHFUTCoursesList(c *gin.Context, username string, bizTypeId int, semesterId int) (resp types.CourseListResponse, code int, err error) {
	if bizTypeId != 2 && bizTypeId != 23 {
		return resp, codes.CodeCourseTableCampusFault, nil
	}
	result := services.NewHFUTApi().WithContext(c).
		WithPrefix("vpn/eam/courseList").
		WithParmas(
			"bizTypeId", strconv.Itoa(bizTypeId),
			"semesterId", strconv.Itoa(semesterId),
		).
		GenNewRequest().RequestSetCookieByUsername(username).
		Do().Result()
	if err = result.Err(); err != nil {
		return resp, codes.CodeCourseTableQueryFail, result.Err()
	}
	if result.Val() == nil {
		return resp, result.Code(), nil
	}
	err = json.Unmarshal(result.Val(), &resp)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	return resp, result.Code(), nil
}
