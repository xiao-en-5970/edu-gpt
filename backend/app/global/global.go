package global

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/xiao-en-5970/edu-gpt/backend/app/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.SugaredLogger
var Cfg *conf.Config
var Db *gorm.DB
var RedisClient *redis.Client

type Status int

const (
	Disabled Status = iota // 0
	Active                 // 1
	Locked                 // 2
)

// 为 Color 实现 String() 方法
func (s Status) String() string {
	switch s {
	case Active:
		return "active"
	case Disabled:
		return "disabled"
	case Locked:
		return "locked"
	default:
		return "disabled"
	}
}

func GetUrl(prefix string, id uint) string {
	return fmt.Sprintf("https://%s/api/v1/%s/%d", Cfg.Server.Address, prefix, id)
}

func GetFileUrl(prefix string, path string) string {
	return fmt.Sprintf("https://%s/api/v1/%s?filepath=%s", Cfg.Server.Address, prefix, path)
}

// 2022 学年第一学期【2022 9-12月】作为基准，学期代码为194
// 每一个学期递增20,每学年两个学期递增40
// 例：2023年第一学期【2023 3-2024 7月】对应学期代码为214
// 例：2025年第一学期【2025 9-2026 1月】对应学期代码为314
const BaseSemester = 194 // 基础学期

func YearTerm2SemesterId(year int, term int) int {
	if term < 1 || term >= 3 {
		return 0
	}
	return BaseSemester + (year-2022)*40 + (term-1)*20
}

func GetIDByContext(c *gin.Context) (uint) {
	u, ex := c.Get("id")
	if !ex {
		return 0
	}
	uid := u.(uint)
	return uid
}


//锁

 var LockPostLikeCount sync.Mutex 