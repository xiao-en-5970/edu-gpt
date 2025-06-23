package route

import (
	"github.com/gin-gonic/gin"
	handler "github.com/xiao-en-5970/edu-gpt/backend/app/handler/course"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
)

func RouteCourseInit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/course")

	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/getallcourses", handler.HandlerCourseGetAllCourses)
		auth.POST("/list", handler.HandlerListCourses)
		auth.POST("/list_files", handler.HandlerCourseFiles)
		auth.POST("/get_info", handler.HandlerCourseGetInfo)
		// auth.POST("/getfile", handler.HandlerCourseGetFile)
	}
}
