package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/handler/post"
)

func RoutePostInit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/post")

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/:id",handler.HandlerPost)
		auth.POST("/create", handler.HandlerPostCreate)
		auth.POST("/edit", handler.HandlerPostEdit)
	}
}
