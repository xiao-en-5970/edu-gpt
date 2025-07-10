package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/middleware"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/redisprefix"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
)

// @Summary 帖子查看
// @Description 通过帖子id直接查看帖子
// @Tags Post模块
// @Security     BearerAuth
// @Produce json
// @Param        id  path  string  true  "帖子ID"
// @Success 200 {object} types.PostResp "成功响应"
// @Router /post/auth/{id} [post]
func HandlerPost(c *gin.Context) {

	idstr := c.Param("id")
	pidint, err := strconv.Atoi(idstr)
	if err != nil {
		responce.ErrorBadRequest(c, err)
		return
	}
	pid := uint(pidint)
	post, err := models.FindPostById(c, pid)
	if post == nil {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodePostNotExist)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	post.LikeCount,err = middleware.GetCacheCountMiddleware(
		c,
		pid,
		redisprefix.PrefixPostLikeCountKey,
		models.GetPostLikeCountFromMysql,
		models.UpdatePostLikeCountFromMysql,
	)
	if err !=nil{
		responce.ErrorInternalServerError(c,err)
		return 
	}
	switch post.Active {
	case global.Disabled.String():
		responce.ErrorInternalServerErrorWithCode(c, codes.CodePostLDisabled)
		return
	case global.Locked.String():
		responce.ErrorInternalServerErrorWithCode(c, codes.CodePostLocked)
		return
	}

	u, ex := c.Get("id")
	if !ex {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodeAuthNotExistError)
		return
	}
	uid := u.(uint)
	user, err := models.FindUserById(c, uid)
	if user == nil {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodeUserNotExist)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	images, err := models.FindPostImageByPid(c, pid)
	if images == nil {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodeImageNotExist)
		return
	}
	urls := make([]string, 0, 1)
	for _, i := range images {
		urls = append(urls, global.GetUrl("post/auth/postimage", i.ID))
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	likestatus, err := models.GetLikeStatus(c, pid, uid)
	if err != nil {
		responce.ErrorInternalServerError(c, err)
	}
	poster, err := models.FindUserById(c, post.PosterID)
	if poster == nil {
		responce.ErrorInternalServerErrorWithCode(c, codes.CodeUserNotExist)
		return
	}
	if err != nil {
		responce.ErrorInternalServerError(c, err)
		return
	}
	resp := types.PostResp{
		Post:       *post,
		Nickname:   poster.Nickname,
		Campus:     poster.Campus,
		Grade:      poster.Grade,
		LikeStatus: likestatus,
		Department: poster.Department,
		ImageUrls:  urls,
		Avatar:     global.GetUrl("user/auth/avatar", poster.ID),
	}
	responce.SuccessWithData(c, resp)
}
