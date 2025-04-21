package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/xiao-en-5970/Goodminton/backend/app/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.SugaredLogger
var Cfg * conf.Config
var Db *gorm.DB
var RedisClient *redis.Client