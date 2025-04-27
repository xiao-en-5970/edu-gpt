package logic

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostUploadPostImage(c *gin.Context, req *types.UploadManyImagesReq) (resp *types.UploadManyImagesResp, code int, err error) {
	post, _ := model.FindPostById(req.ID)
	if post == nil {
		//帖子不存在
		return &types.UploadManyImagesResp{}, codes.CodePostNotExist, nil
	}
	urls:=make([]string,0)
	for index, file := range req.Files {
		//生成存储路径
		absPath := fmt.Sprintf("%s/%d_%d%s", global.Cfg.Image.PostPath, post.ID, index+1,path.Ext(file.Filename))
		imageid,err:=model.InsertPostImage(&model.PostImage{
			PostID: post.ID,
			Number: index+1,
			ImagesPath: absPath,
		})
		if err!=nil{
			return &types.UploadManyImagesResp{}, codes.CodeAllIntervalError, err
		}
		// 保存新文件
		if err := c.SaveUploadedFile(file, absPath); err != nil {
			return &types.UploadManyImagesResp{}, codes.CodeAllIntervalError, err
		}
		url := GetUrl("postimage", imageid)
		urls = append(urls, url)
	}
	return &types.UploadManyImagesResp{
		Urls: urls,
	}, codes.CodeAllSuccess, nil
}
func GetUrl(prefix string, id uint) string {
	return fmt.Sprintf("https://%s/api/v1/post/auth/%s/%d", global.Cfg.Server.Address, prefix, id)
}