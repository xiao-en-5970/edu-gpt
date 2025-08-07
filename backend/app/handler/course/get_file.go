package handler

import (
	"os"

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
	filePath := global.Cfg.Static.BookPath +"/"+ filepathstr
	global.Logger.Info("get file path:", filePath)
	if _, err := os.Stat(filePath); err == nil {
		c.File(global.Cfg.Static.BookPath +"/"+ filepathstr)
    } else if os.IsNotExist(err) {
		responce.ErrorBadRequestWithCode(c,codes.CodeFileNotExist)
    } else {
        responce.ErrorBadRequest(c,err)
    }
	
}
