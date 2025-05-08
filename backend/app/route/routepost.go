package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/handler/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
)

func RoutePostInit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/post")

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/:id", handler.HandlerPost)
		auth.POST("/create", handler.HandlerPostCreate)
		auth.POST("/edit", handler.HandlerPostEdit)
		auth.GET("/postimage/:id", handler.HandlerPostPostImage)
		auth.POST("/postimage/:id", handler.HandlerPostPostImage)
		auth.POST("/list", handler.HandlerPostList)
		auth.POST("/like", handler.HandlerPostLike)
	}
}
