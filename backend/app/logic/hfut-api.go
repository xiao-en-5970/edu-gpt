package logic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/types"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicHFUTLogin(req *types.HFUTLoginReq)(resp *types.HFUTLoginResp,code int,err error){
	c:=&http.Client{}
	params := url.Values{}
	params.Add("username", req.Username)
	params.Add("password", req.Password)
	r,_ := http.NewRequest("GET",fmt.Sprintf("http://%s:%d/login?",global.Cfg.HfutAPI.Host,global.Cfg.HfutAPI.Port)+params.Encode(),nil)
	
	rsp,err:=c.Do(r)
	if err!=nil{
		return &types.HFUTLoginResp{},codes.CodeAllBadGateway,err
	}
	defer rsp.Body.Close()
	
	if rsp.StatusCode==200{
		//登录信息门户成功
		body,_:=io.ReadAll(rsp.Body)
		resp=&types.HFUTLoginResp{}
		err=json.Unmarshal(body,resp)
		if err!=nil{
			return &types.HFUTLoginResp{},codes.CodeHFUTIntervalError,nil
		}
		return resp,codes.CodeAllSuccess,nil
	}else if rsp.StatusCode==400{
		//登录信息门户失败
		return &types.HFUTLoginResp{},codes.CodeHFUTLoginError,nil
	}else if rsp.StatusCode==500{
		//登录信息门户限流
		for i := 0; i < global.Cfg.HfutAPI.Retry ; i++{
			time.Sleep(1000*time.Millisecond)
			rsp,err:=c.Do(r)
			if err!=nil{
				return &types.HFUTLoginResp{},codes.CodeAllBadGateway,err
			}
			if rsp.StatusCode!=500{
				break
			}
		}
		if rsp.StatusCode==200{
			//登录信息门户成功
			body,_:=io.ReadAll(rsp.Body)
			resp=&types.HFUTLoginResp{}
			err=json.Unmarshal(body,resp)
			if err!=nil{
				return &types.HFUTLoginResp{},codes.CodeHFUTIntervalError,nil
			}
			return resp,codes.CodeAllSuccess,nil
		}else if rsp.StatusCode==400{
			//登录信息门户失败
			return &types.HFUTLoginResp{},codes.CodeHFUTLoginError,nil
		}else{
			return &types.HFUTLoginResp{},codes.CodeHFUTIntervalError,nil
		}
		
	}else{
		global.Logger.Infof("rsp.Body: %v\n", )
		//登录信息门户未知问题
		return &types.HFUTLoginResp{},codes.CodeHFUTUnkonwnError,nil
	}
}

func LogicHFUTStudentInfo(c *gin.Context,username string)(resp *types.HFUTStudentInfoResp,code int,err error){
	client:=&http.Client{}
	r,_ := http.NewRequest("GET",fmt.Sprintf("http://%s:%d/eam/studentinfo?",global.Cfg.HfutAPI.Host,global.Cfg.HfutAPI.Port),nil)
	result :=global.RedisClient.Get(c,username)
	if result.Err()!=nil{
		return &types.HFUTStudentInfoResp{},codes.CodeHFUTIntervalError,nil
	}
	cookie:=result.Val()
	global.Logger.Infof("cookie:%s",cookie)
	r.Header.Set("cookie",cookie)
	rsp,err:=client.Do(r)
	if err!=nil{
		return &types.HFUTStudentInfoResp{},codes.CodeAllBadGateway,err
	}
	defer rsp.Body.Close()
	
	if rsp.StatusCode==200{
		//获取信息成功
		bytebody,_:=io.ReadAll(rsp.Body)
		hfutrsp:=&types.HFUTStudentInfoResp{}
		err:=json.Unmarshal(bytebody,hfutrsp)
		if err!=nil{
			return &types.HFUTStudentInfoResp{},codes.CodeHFUTIntervalError,nil
		}
		return hfutrsp,codes.CodeAllSuccess,nil
	}else if rsp.StatusCode==401||rsp.StatusCode==400{
		// 删除 Redis 中的无效 cookie
		global.RedisClient.Del(c, username)
		//未登录
		return &types.HFUTStudentInfoResp{},codes.CodeHFUTNotLogin,nil
	}else if rsp.StatusCode==500{
		//信息门户限流
		for i := 0; i < global.Cfg.HfutAPI.Retry ; i++{
			time.Sleep(100*time.Millisecond)
			rsp,err:=client.Do(r)
			if err!=nil{
				return &types.HFUTStudentInfoResp{},codes.CodeAllBadGateway,err
			}
			if rsp.StatusCode!=500{
				break
			}
		}
		return &types.HFUTStudentInfoResp{},codes.CodeHFUTIntervalError,nil
	}else{
		global.Logger.Infof("rsp.Body: %v\n", rsp.Body)
		//登录信息门户未知问题
		return &types.HFUTStudentInfoResp{},codes.CodeHFUTUnkonwnError,nil
	}
}