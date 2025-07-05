package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func GetPrefix(prefix string, key string) string {
	return prefix + ":" + key
}

// 自定义ResponseWriter用于捕获响应数据
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 增加读缓存中间件
func AddReadCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		rawdata, err := c.GetRawData()
		if err != nil {
			responce.ErrorBadRequest(c, err)
			return
		}
		c.Set("raw_req", string(rawdata)) // 将原始数据存入上下文
		path_prefix := GetPrefix("path", path)
		json_prefix := GetPrefix("json", string(rawdata))
		suf := GetPrefix(path_prefix, json_prefix)
		key := GetPrefix("cache:read", suf)
		result, rediserr := global.RedisClient.Get(c, key).Result()

		if rediserr == redis.Nil {
			global.Logger.Debug("Cache miss!   ", "key:", key)
			writer := &responseWriter{ResponseWriter: c.Writer, body: bytes.NewBufferString("")}
			c.Writer = writer
			c.Next() // 继续处理请求
			if c.Writer.Status() == http.StatusOK {
				// 获取响应数据
				responseData := writer.body.String()
				// global.Logger.Debug("Response data:  ", responseData)
				// 将数据写入Redis
				if err := global.RedisClient.SetEx(c, key, responseData, time.Duration(global.Cfg.Cache.ReadExpiration)*time.Second).Err(); err != nil {
					global.Logger.Error("Failed to set cache    ", "key:", key, "error", err)
					responce.ErrorInternalServerError(c, err)
				} else {
					global.Logger.Debug("Cache set successfully   ", "key:", key)
				}
				return
			} else {
				global.Logger.Error("Request processing failed", "status", c.Writer.Status)
				return
			}

		} else if rediserr != nil {
			global.Logger.Error("Failed to get cache", "key", key, "error", rediserr)
			c.Next() // 继续处理请求
			return
		} else {
			global.Logger.Debug("Cache hit!   ", "key:", key)
			var cachedMap map[string]interface{}
			if err := json.Unmarshal([]byte(result), &cachedMap); err != nil {
				global.Logger.Error("Failed to unmarshal cache data", "key:", key, "error", err)
				c.Next() // 解析失败，继续后续处理
				return
			}
			c.JSON(http.StatusOK, cachedMap) // 返回缓存数据
			c.Abort()
			// 4. 续期缓存TTL
			if err := global.RedisClient.Expire(c, key, time.Duration(global.Cfg.Cache.ReadExpiration)*time.Second).Err(); err != nil {
				global.Logger.Error("Failed to extend cache TTL", "key:", key, "error", err)
			} else {
				global.Logger.Debug("Cache TTL extended successfully", "key:", key)
			}
			return
		}

	}
}

type FuncMysqlGetCount func(c *gin.Context, id uint) (count int, err error)
type FuncMysqlUpdateCount func(c *gin.Context, id uint, newcount int) (err error)

func GetCacheCountMiddleware(c *gin.Context, id uint,prefix string, getfunc FuncMysqlGetCount, updatefunc FuncMysqlUpdateCount) (count int, err error) {
	//先尝试从redis获取
	result := global.RedisClient.Get(c, prefix+strconv.Itoa(int(id)))
	if result.Err() != nil {
		if result.Err() == redis.Nil {
			//获取失败则从mysql获取数量
			count, err = getfunc(c, id)
			if err != nil {
				global.Logger.Error("Failed to get count from MySQL", "error", err)
				return -1 ,err
			}
			//然后存入redis，然后设置time*2过期，
			result := global.RedisClient.SetEx(c,
				prefix+strconv.Itoa(int(id)),
				strconv.Itoa(count),
				//最大存活时间为两倍预设时间，预设时间到了则刷入mysql
				time.Duration(global.Cfg.Cache.WriteExpiration*2)*time.Second)
			if result.Err() != nil {
				return -1, result.Err()
			}
			// 并且设置定时任务在redis中60s后刷入数据库
			go func() {
				//等待一段时间，确保缓存充分命中
				time.Sleep(time.Duration(global.Cfg.Cache.WriteExpiration) * time.Second)
				//从redis中获取最新的count
				newcountstr ,err := global.RedisClient.Get(c, prefix+strconv.Itoa(int(id))).Result()
				if err != nil {
					global.Logger.Error("Failed to get count from redis", "error", err)
					return
				}
				newcount,err:= strconv.Atoi(newcountstr)
				if err != nil {
					global.Logger.Error("Failed to convert count from string to int", "error", err)
					return
				}
			// 删缓存 刷数据库 删缓存。双删保持一致性
				//1.删除redis中的缓存
				delresult := global.RedisClient.Del(c, prefix+strconv.Itoa(int(id)))
				if delresult.Err() != nil {
					global.Logger.Error("Failed to delete count from redis1", "error", delresult.Err())
				} else {
					global.Logger.Debug("Count deleted from redis successfully1", "key", prefix+strconv.Itoa(int(id)))
				}
				//2.刷入数据库
				err = updatefunc(c, id, newcount)
				if err != nil {
					global.Logger.Error("Failed to update count in database", "error", err)
					return
				}
				//3.删除redis中的缓存
				delresult = global.RedisClient.Del(c, prefix+strconv.Itoa(int(id)))
				if delresult.Err() != nil {
					global.Logger.Error("Failed to delete count from redis2", "error", delresult.Err())
				} else {
					global.Logger.Debug("Count deleted from redis successfully2", "key", prefix+strconv.Itoa(int(id)))
				}
			}()
			return count, nil
		}
		return -1, result.Err()
	} else {
		count, err = strconv.Atoi(result.Val())
		if err != nil {
			return -1, err
		}
		return count, nil
	}
}

func UpdateCacheCount(c *gin.Context, id uint,prefix string,addValue int, getfunc FuncMysqlGetCount, updatefunc FuncMysqlUpdateCount)(err error){
	_,err = GetCacheCountMiddleware(c,id,prefix,getfunc,updatefunc)
	if err !=nil{
		return err
	}
	_,err= global.RedisClient.IncrBy(c,prefix,int64(addValue)).Result()
	if err !=nil{
		return err
	}
	return nil
}

// 删除读缓存中间件
func DelReadCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next() // 继续处理请求
	}
}

