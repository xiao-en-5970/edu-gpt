package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/conf"
	"github.com/xiao-en-5970/edu-gpt/backend/app/db"
	docs "github.com/xiao-en-5970/edu-gpt/backend/app/docs"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/route"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/zaplog"
)

// @title           接口文档
// @version         1.0
// @BasePath        /api/v1
// @securityDefinitions.apikey  BearerAuth
// @in              header
// @name            Authorization
// @description     Bearer Token 格式：`Bearer {token}`（注意空格）
// @contact.name    API 支持
// @host            localhost:8080
func main() {
	cfg, err := conf.ConfInit("./config.yaml")
	if err != nil {
		panic(err)
	}
	global.Cfg = cfg
	zaplog.InitZapLogger()
	db.InitDB()
	docs.SwaggerInfo.BasePath = "/api/v1"  // 与注释保持一致
	docs.SwaggerInfo.Host = global.Cfg.Server.Address // 确保Host更新
	r := gin.New()
	route.RouteInit(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", global.Cfg.Server.Port)) // 默认监听8080端口
	
}
