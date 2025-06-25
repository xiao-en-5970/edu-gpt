package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerCourseGetFile(c *gin.Context) {
	filepathstr := c.Query("filepath")
	if filepathstr == "" {
		responce.ErrorBadRequestWithCode(c, codes.CodeFileQueryFilePathEmpty)
		return
	}
	global.Logger.Info("get file path:", global.Cfg.Static.BookPath +"/"+ filepathstr)
	c.File(global.Cfg.Static.BookPath +"/"+ filepathstr)
}
