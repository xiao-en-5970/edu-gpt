package logic

import (

	"net/http"
	"net/url"
	"time"
	"github.com/xiao-en-5970/Goodminton/backend/app/global"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
)

func LogicLoginHFUT(req *types.LoginHFUTReq)(code int,err error){
	c:=&http.Client{}
	params := url.Values{}
	params.Add("username", req.Username)
	params.Add("password", req.Password)
	r,_ := http.NewRequest("GET","http://127.0.0.1:8082/login?"+params.Encode(),nil)
	
	rsp,err:=c.Do(r)
	if err!=nil{
		return codes.CodeAllBadGateway,err
	}
	defer rsp.Body.Close()
	
	if rsp.StatusCode==200{
		//登录信息门户成功
		return codes.CodeUserLoginSchoolAuthSuccess,nil
	}else if rsp.StatusCode==400{
		//登录信息门户失败
		return codes.CodeUserLoginSchoolAuthError,nil
	}else if rsp.StatusCode==500{
		//登录信息门户限流
		for i := 0; i < global.Cfg.HfutAPI.Retry ; i++{
			time.Sleep(100*time.Millisecond)
			rsp,err:=c.Do(r)
			if err!=nil{
				return codes.CodeAllBadGateway,err
			}
			if rsp.StatusCode!=500{
				break
			}
		}
		return codes.CodeUserLoginSchoolIntervalError,nil
	}else{
		global.Logger.Infof("rsp.Body: %v\n", )
		//登录信息门户未知问题
		return codes.CodeUserLoginSchoolUnkonwnError,nil
	}
}