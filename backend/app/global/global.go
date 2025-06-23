package global

import (
	"fmt"

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
	return fmt.Sprintf("https://%s/api/v1/%s/%s", Cfg.Server.Address, prefix, path)
}