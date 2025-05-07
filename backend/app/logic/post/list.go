package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/post"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)


func LogicPostList(c *gin.Context,req *types.PostListReq)(resp types.PostListResp,code int,err error){
	u,ex:=c.Get("id")
	if !ex{
		return types.PostListResp{},codes.CodeAuthUnvalidToken,nil
	}
	uid := u.(uint)
	user,err := model.FindUserById(c,uid)
	if user == nil{
		return types.PostListResp{},codes.CodeUserNotExist,nil
	}
	if err!=nil{
		return types.PostListResp{},codes.CodeAllIntervalError,err
	}
	posts,err:=model.ListPost(c,req.Offset,req.Limit)
	if err!=nil{
		return types.PostListResp{},codes.CodeAllIntervalError,err
	}
	briefposts:=make([]types.BriefPost,0,1)
	for _,p:=range(posts){
		briefposts = append(briefposts, types.BriefPost{
			Title: p.Title,
			Content: p.Content,
			Nickname: user.Nickname,
			ID: p.ID,
			PosterID: p.PosterID,
			ViewCount: p.ViewCount,
			LikeCount: p.LikeCount,
			CollectCount: p.CollectCount,
			CommentCount: p.CommentCount,
			CreateAt: p.CreateAt,
			AvatarUrl: GetUrl("avatar",p.PosterID),
		})
	}
	return briefposts,codes.CodeAllSuccess,nil
}