package route

import (
	"github.com/gin-gonic/gin"
	handler "github.com/xiao-en-5970/edu-gpt/backend/app/handler/community"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
)

func RouteCommunityinit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/community")

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/create",handler.HandlerCommunityCreate)
		// auth.POST("/edit",handler.Edit)
		// auth.POST("/delete",handler.Delete)
		// auth.POST("/list",handler.List)
	}
}