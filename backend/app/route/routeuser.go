package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/handler"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
)

func UserRouteInit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/user")
	r.POST("/login", handler.HandlerLogin)
	
	auth:=r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/get_userinfo",handler.HandlerGetUserInfo)
		auth.POST("/update_userinfo",handler.HandlerUpdateUserInfo)
	}
	
}
