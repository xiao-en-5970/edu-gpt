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

func GetCacheCountMiddleware(c *gin.Context, id uint, prefix string, getfunc FuncMysqlGetCount, updatefunc FuncMysqlUpdateCount) (count int, err error) {
	//先尝试从redis获取
	countstr, err := global.RedisClient.Get(c, GetPrefix(prefix, strconv.Itoa(int(id)))).Result()
	if err != nil {
		if err == redis.Nil {
			global.Logger.Debugf("缓存未命中 key: %s", GetPrefix(prefix, strconv.Itoa(int(id))))
			//获取失败则从mysql获取数量
			count, err = getfunc(c, id)
			if err != nil {
				global.Logger.Error("Failed to get count from MySQL", "error", err)
				return -1, err
			}
			//然后存入redis，然后设置不过期，
			global.Logger.Debugf("加入缓存 key %s: value: %d", GetPrefix(prefix, strconv.Itoa(int(id))),count)
			_, err := global.RedisClient.Set(c,
				GetPrefix(prefix, strconv.Itoa(int(id))),
				count,
				0,
			).Result()
			if err != nil {
				return -1, err
			}
			return count, nil
		} else {
			return -1, err
		}
	} else {
		global.Logger.Debugf("缓存命中！ key: %s", GetPrefix(prefix, strconv.Itoa(int(id))))
		count, err = strconv.Atoi(countstr)
		if err != nil {
			return -1, err
		}

	}
	return count, nil
}

// 缓存刷入盘
func Cache2Mysql(c *gin.Context, prefix string, id uint , newcount int, updatefunc FuncMysqlUpdateCount) {
	// 并且设置定时任务在redis中30s后刷入数据库
	//占有一定时间的锁，确保缓存充分命中
	global.RedisClient.Set(c, GetPrefix("lock:"+prefix, strconv.Itoa(int(id))), "1", time.Duration(global.Cfg.Cache.WriteExpiration)*time.Second).Result()
	// 删缓存 刷数据库 删缓存。双删保持一致性
	//1.删除redis中的缓存
	delresult := global.RedisClient.Del(c, GetPrefix(prefix, strconv.Itoa(int(id))))
	if delresult.Err() != nil {
		global.Logger.Error("Failed to delete count from redis1", "error", delresult.Err())
	} else {
		global.Logger.Debug("一次删除缓存\t", "key\t", GetPrefix(prefix, strconv.Itoa(int(id))))
	}
	//2.刷入数据库
	err := updatefunc(c, id, newcount)
	if err != nil {
		global.Logger.Error("Failed to update count in database", "error", err)
		return
	}
	global.Logger.Debug("写入数据库\t")
	//3.删除redis中的缓存
	delresult = global.RedisClient.Del(c, GetPrefix(prefix, strconv.Itoa(int(id))))
	if delresult.Err() != nil {
		global.Logger.Error("Failed to delete count from redis2", "error", delresult.Err())
	} else {
		global.Logger.Debug("二次删除缓存\t", "key\t", GetPrefix(prefix, strconv.Itoa(int(id))))
	}
}

func UpdateCacheCount(c *gin.Context, id uint, prefix string, addValue int, getfunc FuncMysqlGetCount, updatefunc FuncMysqlUpdateCount) (err error) {
	_, err = GetCacheCountMiddleware(c, id, prefix, getfunc, updatefunc)
	if err != nil {
		return err
	}
	count, err := global.RedisClient.IncrBy(c, GetPrefix(prefix, strconv.Itoa(int(id))), int64(addValue)).Result()
	if err != nil {
		return err
	}
	
	//获取刷数据库锁
	err = global.RedisClient.Get(c, GetPrefix("lock:"+prefix, strconv.Itoa(int(id)))).Err()
	if err != nil {
		//如果获取成功【锁未被占用】则刷入数据库
		if err == redis.Nil {
			global.Logger.Debugf("获取锁，开始刷入数据库")
			go Cache2Mysql(c, prefix, id, int(count),updatefunc)
			return nil
		} else {
			global.Logger.Error(err)
			return err
		}
	}
	global.Logger.Debugf("获取锁失败，不刷入数据库")
	return nil
}

// 删除读缓存中间件
func DelReadCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next() // 继续处理请求
	}
}
