package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/handler/comment"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
)

func RouteCommentInit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/comment")
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/create", handler.HandlerCommentCreate)
		auth.POST("/createreply", handler.HandlerSubCommentCreate)
		// TODO
		auth.POST("/list", handler.HandlerCommentList)
		// TODO
		auth.POST("/listreply", handler.HandlerSubCommentList)
		// TODO
		auth.POST("/like", handler.HandlerCommentLike)
		// TODO
		auth.POST("/likereply", handler.HandlerSubCommentLike)
		
	}
}
