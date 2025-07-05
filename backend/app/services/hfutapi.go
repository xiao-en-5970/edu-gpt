package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/redisprefix"
)

type ServiceHFUTApi struct {
	ctx       *gin.Context
	err       error
	req       *http.Request
	client    http.Client
	resp      *http.Response
	retry     int
	body      []byte
	code      int
	username  string
	params    url.Values
	prefix    string
	hasCookie bool
}

type ServiceHFUTApiResult struct {
	err  error
	body []byte
	code int
}

func NewHFUTApi() *ServiceHFUTApi {
	h := new(ServiceHFUTApi)
	if h.err != nil {
		return h
	}
	h.client = http.Client{
		Timeout: time.Second * time.Duration(10), // 设置超时时间为10秒

	}
	h.retry = 0
	h.code = codes.CodeAllSuccess

	h.params = url.Values{}
	h.hasCookie = false
	return h
}
func (h *ServiceHFUTApi) WithPrefix(prefix string) *ServiceHFUTApi {
	if h.err != nil {
		return h
	}
	h.prefix = prefix
	return h
}

func (h *ServiceHFUTApi) WithContext(c *gin.Context) *ServiceHFUTApi {
	if h.err != nil {
		return h
	}
	h.ctx = c
	return h
}
func (h *ServiceHFUTApi) WithParmas(keyvalues ...string) *ServiceHFUTApi {
	if h.err != nil {
		return h
	}
	n := len(keyvalues)
	if n%2 != 0 {
		h.err = errors.New("key value数量不对等")
	}
	for i := 0; i < n-1; i += 2 {
		h.params.Add(keyvalues[i], keyvalues[i+1])
	}
	return h
}

func (h *ServiceHFUTApi) GenNewRequest() *ServiceHFUTApi {
	if h.err != nil {
		return h
	}
	if err := h.ctx.Err(); err != nil {
		h.err = err
		return h // 直接返回取消或超时错误
	}
	address := fmt.Sprintf("http://%s:%d/%s?%s", global.Cfg.HfutAPI.Host, global.Cfg.HfutAPI.Port, h.prefix, h.params.Encode())
	global.Logger.Debugf("hfut api address: %s", address)
	h.req, h.err = http.NewRequestWithContext(h.ctx, "GET", address, nil)
	return h
}

func (h *ServiceHFUTApi) RequestSetCookieByUsername(username string) *ServiceHFUTApi {
	if h.err != nil {
		return h
	}
	if err := h.ctx.Err(); err != nil {
		h.err = err
		return h // 直接返回取消或超时错误
	}
	h.username = username
	result := global.RedisClient.Get(h.ctx, middleware.GetPrefix(redisprefix.PrefixUserCookieKey, h.username))
	h.err = result.Err()
	cookie := result.Val()
	global.Logger.Debugf("username:%s cookie:%s", h.username, cookie)
	h.req.Header.Set("cookie", cookie)
	h.hasCookie = true
	return h
}
func (h *ServiceHFUTApi) Do() *ServiceHFUTApi {
	if h.err != nil {
		return h
	}
	if h.retry > global.Cfg.HfutAPI.Retry {
		h.code = codes.CodeHFUTRetrySoMuch
		return h
	}
	if err := h.ctx.Err(); err != nil {
		h.err = err
		return h // 直接返回取消或超时错误
	}
	h.resp, h.err = h.client.Do(h.req)
	return h
}

func (h *ServiceHFUTApi) Result() (res ServiceHFUTApiResult) {
	if h.err != nil {
		res.err = h.err
		return res
	}
	if err := h.ctx.Err(); err != nil {
		res.err = err
		return res // 直接返回取消或超时错误
	}
	if h.resp.StatusCode == 200 {
		//获取信息成功
		h.body, h.err = io.ReadAll(h.resp.Body) // 不再忽略错误
		if h.err != nil {
			res.err = h.err
			return res
		}
	} else if h.resp.StatusCode == 401 || h.resp.StatusCode == 400 {
		// 删除 Redis 中的无效 cookie
		if h.hasCookie {
			result := global.RedisClient.Del(h.ctx, middleware.GetPrefix(redisprefix.PrefixUserCookieKey, h.username))
			h.err = result.Err()
		}
		//未登录
		h.code = codes.CodeHFUTNotLogin
	} else if h.resp.StatusCode == 500 {
		//信息门户限流,重试
		h.retry++
		time.Sleep(100 * time.Millisecond)
		h.Do().Result()
	}
	res.body = h.body
	res.code = h.code
	res.err = h.err
	return res
}
func (h *ServiceHFUTApiResult) Err() error {
	return h.err
}
func (h *ServiceHFUTApiResult) Val() []byte {
	return h.body
}
func (h *ServiceHFUTApiResult) Code() int {
	return h.code
}
