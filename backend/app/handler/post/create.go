package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	logic "github.com/xiao-en-5970/edu-gpt/backend/app/logic/post"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

func HandlerPostCreate(c *gin.Context) {
	req := &types.CreatePostReq{}
	form, err := c.MultipartForm()
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
	if form==nil{
		responce.ErrorBadRequest(c, err)
		return
	}
	files := form.File["postimage"]   // 获取文件
	jsonData := form.Value["json"][0] // 获取JSON字符串
	err = json.Unmarshal([]byte(jsonData), req)
	if err != nil {
		responce.ErrorBadRequest(c, err)
	}
	
	resp, code, err := logic.LogicPostCreate(c, req)
	if code != codes.CodeAllSuccess {
		responce.ErrorInternalServerErrorWithCode(c, code)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
	}
	_, code, err = logic.LogicPostUploadPostImage(c, &types.UploadManyImagesReq{Files: files,ID: resp.ID})
	if code != codes.CodeAllSuccess {
		responce.ErrorInternalServerErrorWithCode(c, code)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
	}
	responce.SuccessWithData(c, resp)

}
