package route

import (
	// "os"
	// "os/exec"

	// "github.com/fsnotify/fsnotify"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func RouteInit(r *gin.Engine) {
	
	apiGroup := r.Group("api/v1")
	apiGroup.Use(middleware.ZapLogger(global.Logger))
	apiGroup.Use(middleware.ErrorMiddleware(global.Logger))
	apiGroup.GET("/", func(c *gin.Context) {
		responce.SuccessWithMsg(c, "测试成功!")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pprof.Register(r)
	//监听文件变化（示例用 fsnotify）
	watcher, _ := fsnotify.NewWatcher()
    watcher.Add("./docs")       // Swagger生成的文档目录
    watcher.Add("./handlers")   // 接口代码目录
    watcher.Add("./models")     // 模型定义目录
    watcher.Add("./routers")    // 路由定义目录
	go func() {
		for event := range watcher.Events {
			if event.Op&fsnotify.Write == fsnotify.Write {
				global.Logger.Debug("重新生成接口文档")
				cmd := exec.Command("swag", "init") // 重新生成文档
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
		}
	}()
	RouteUserInit(apiGroup)
	RoutePostInit(apiGroup)
	RouteCommentInit(apiGroup)
	RouteCourseInit(apiGroup)
	RouteCommunityinit(apiGroup)
}
