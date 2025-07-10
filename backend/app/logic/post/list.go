package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicPostList(c *gin.Context, req *types.PostListReq) (resp types.PostListResp, code int, err error) {
	_, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	
	if req.Order == "" {
		return resp, codes.CodeAllRequestFormatError, nil
	}
	posts, err := models.ListPost(c, req.Page, req.Size, req.Desc, req.Order)
	if err != nil {
		return resp, codes.CodeAllIntervalError, err
	}
	briefposts := make([]types.BriefPost, 0, 1)
	for _, p := range posts {
		poster, err := models.FindUserById(c, p.PosterID)
		if poster == nil {
			return resp, codes.CodeUserNotExist, nil
		}
		if err != nil {
			return resp, codes.CodeAllIntervalError, err
		}
		if p.Active == global.Active.String() {
			briefposts = append(briefposts, types.BriefPost{
				Title:        p.Title,
				Content:      p.Content,
				Nickname:     poster.Nickname,
				ID:           p.ID,
				PosterID:     p.PosterID,
				ViewCount:    p.ViewCount,
				Active:       p.Active,
				LikeCount:    p.LikeCount,
				CollectCount: p.CollectCount,
				CommentCount: p.CommentCount,
				CreateAt:     p.CreateAt,
				AvatarUrl:    global.GetUrl("user/auth/avatar", p.PosterID),
			})
		}
	}
	return briefposts, codes.CodeAllSuccess, nil
}
