package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/global"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/responce"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/zaplog"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)


func RouteInit(r * gin.Engine){
	
	apiGroup := r.Group("api/v1")
	apiGroup.Use(zaplog.ZapLogger(global.Logger))
	apiGroup.GET("/",func(c *gin.Context) {
		responce.SuccessWithMsg(c,"测试成功!")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	UserRouteInit(apiGroup)
}