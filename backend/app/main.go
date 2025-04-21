package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/conf"
	"github.com/xiao-en-5970/Goodminton/backend/app/db"
	"github.com/xiao-en-5970/Goodminton/backend/app/global"
	"github.com/xiao-en-5970/Goodminton/backend/app/route"
	"go.uber.org/zap"
)

// @title 用户认证服务API文档
// @version 1.0
// @description 用户登录认证接口文档
// @contact.name API支持
// @contact.email 1076763695@qq.com
// @license.name MIT
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg,err := conf.ConfInit("./config.yaml")
	if err!=nil{
		panic(err)
	}
	global.Cfg = cfg
    // 高性能场景先用zap.Logger
	rawLogger, _ := zap.NewProduction()
	// 需要易用性时临时转换
	sugar := rawLogger.Sugar()
	global.Logger = sugar
	defer global.Logger.Sync() // 注意: Sugar的Sync开销比Logger大
	db.InitDB()

    r := gin.New()
    route.RouteInit(r)
    r.Run(fmt.Sprintf("0.0.0.0:%d",global.Cfg.Server.Port)) // 默认监听8080端口
}