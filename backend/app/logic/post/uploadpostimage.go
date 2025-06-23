package logic

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostUploadPostImage(c *gin.Context, req *types.UploadManyImagesReq) (resp types.UploadManyImagesResp, code int, err error) {
	post, _ := models.FindPostById(c, req.ID)
	if post == nil {
		//帖子不存在
		return resp, codes.CodePostNotExist, nil
	}
	urls := make([]string, 0)
	for index, file := range req.Files {
		//生成存储路径
		absPath := fmt.Sprintf("%s/%d_%d%s", global.Cfg.Static.PostPath, post.ID, index+1, path.Ext(file.Filename))
		imageid, err := models.InsertPostImage(c, &models.PostImage{
			PostID:     post.ID,
			Number:     index + 1,
			ImagesPath: absPath,
		})
		if err != nil {
			return resp, codes.CodeAllIntervalError, err
		}
		// 保存新文件
		if err := c.SaveUploadedFile(file, absPath); err != nil {
			return resp, codes.CodeAllIntervalError, err
		}
		url := global.GetUrl("post/auth/postimage", imageid)
		urls = append(urls, url)
	}
	resp.Urls = urls
	return resp, codes.CodeAllSuccess, nil
}
