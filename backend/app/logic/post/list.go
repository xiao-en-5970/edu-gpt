package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)


func LogicPostList(c *gin.Context,req *types.PostListReq)(resp types.PostListResp,code int,err error){
	_,ex:=c.Get("id")
	if !ex{
		return types.PostListResp{},codes.CodeAuthUnvalidToken,nil
	}
	
	posts,err:=model.ListPost(c,req.Offset,req.Limit)
	if err!=nil{
		return types.PostListResp{},codes.CodeAllIntervalError,err
	}
	briefposts:=make([]types.BriefPost,0,1)
	for _,p:=range(posts){
		poster,err := model.FindUserById(c,p.PosterID)
		if poster == nil{
			return types.PostListResp{},codes.CodeUserNotExist,nil
		}
		if err!=nil{
			return types.PostListResp{},codes.CodeAllIntervalError,err
		}
		briefposts = append(briefposts, types.BriefPost{
			Title: p.Title,
			Content: p.Content,
			Nickname: poster.Nickname,
			ID: p.ID,
			PosterID: p.PosterID,
			ViewCount: p.ViewCount,
			LikeCount: p.LikeCount,
			CollectCount: p.CollectCount,
			CommentCount: p.CommentCount,
			CreateAt: p.CreateAt,
			AvatarUrl: global.GetUrl("user/auth/avatar",p.PosterID),
		})
	}
	return briefposts,codes.CodeAllSuccess,nil
}