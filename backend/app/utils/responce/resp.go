package responce

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
)



func Success(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":codes.CodeAllSuccess,
		"msg":codes.CodeMsg[codes.CodeAllSuccess],
		"data":"null",
	})
}
func SuccessWithCode(c *gin.Context,code int){
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":codes.CodeMsg[code],
		"data":"null",
	})
}



func SuccessWithMsg(c *gin.Context,msg string){
	c.JSON(http.StatusOK,gin.H{
		"code":codes.CodeAllSuccess,
		"msg":msg,
		"data":"null",
	})
}

func SuccessWithData(c *gin.Context,data interface{}){
	c.JSON(http.StatusOK,gin.H{
		"code":codes.CodeAllSuccess,
		"msg":codes.CodeMsg[codes.CodeAllSuccess],
		"data":data,
	})
}

func SuccessWithCodeData(c *gin.Context,code int,data interface{}){
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":codes.CodeMsg[code],
		"data":data,
	})
}

func ErrorBadRequest(c *gin.Context,err error){
	c.JSON(http.StatusBadRequest,gin.H{
		"code":codes.CodeAllRequestFormatError,
		"msg":codes.CodeMsg[codes.CodeAllRequestFormatError],
		"err":err.Error(),
	})
}

func ErrorBadRequestWithCode(c *gin.Context,code int){
	c.JSON(http.StatusBadRequest,gin.H{
		"code":code,
		"msg":codes.CodeMsg[code],
		"data":"null",
	})
}

func ErrorBadGateway(c *gin.Context,err error){
	c.JSON(http.StatusBadGateway,gin.H{
		"code":codes.CodeAllBadGateway,
		"msg":codes.CodeMsg[codes.CodeAllBadGateway],
		"err":err.Error(),
	})
}

func ErrorInternalServerError(c *gin.Context,err error){
	c.JSON(http.StatusInternalServerError,gin.H{
		"code":codes.CodeAllIntervalError,
		"msg":codes.CodeMsg[codes.CodeAllIntervalError],
		"err":err.Error(),
	})
}
