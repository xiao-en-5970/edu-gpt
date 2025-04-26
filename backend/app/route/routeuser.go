package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/handler/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
)

func RouteUserInit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/user")
	r.POST("/login", handler.HandlerUserLogin)
	
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/:id", handler.HandlerUser)
		auth.POST("/imageurl/:id", handler.HandlerUserImageUrl)
		auth.POST("/get_userinfo", handler.HandlerUserGetUserInfo)
		auth.POST("/update_userinfo", handler.HandlerUserUpdateUserInfo)
		auth.POST("/upload_image", handler.HandlerUserUploadImage)
	}
}
