package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostEdit(c *gin.Context, req *types.EditPostReq) (resp types.EditPostResp, code int, err error) {
	if req.ID == 0 {
		return resp, codes.CodeAllRequestFormatError, nil
	}
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	user, err := models.FindUserById(c, uid)
	if user == nil {
		return resp, codes.CodeUserNotExist, nil
	}
	post := &models.Post{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
	}
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	err = models.UpdatePost(c, post, req.ID)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	rsppost, err := models.FindPostById(c, req.ID)
	if rsppost == nil {
		return resp, codes.CodePostNotExist, err
	}
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	resp.ID = req.ID
	return resp, codes.CodeAllSuccess, nil
}
