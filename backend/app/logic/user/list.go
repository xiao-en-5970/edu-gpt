package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/model"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

// 用户列表
func LogicFollowList(c *gin.Context, req *types.FollowListReq) (resp types.FollowListResp, code int, err error) {
	id, ex := c.Get("id")
	if !ex {
		return types.FollowListResp{}, codes.CodeAuthUnvalidToken, nil
	}
	if req.Order == "" {
		return types.FollowListResp{}, codes.CodeAllRequestFormatError, nil
	}
	uid := id.(uint)
	users, err := model.FollowFansList(c, uid, req.Page, req.Size, req.Desc, req.Order, true)
	if err != nil {
		return types.FollowListResp{}, codes.CodeAllIntervalError, err
	}

	briefffs := make([]types.BriefFollow, 0, req.Size)
	for _, u := range users {
		userinfo, err := model.FindUserById(c, u.Follow)
		if userinfo == nil {
			return types.FollowListResp{}, codes.CodeUserNotExist, nil
		}
		if err != nil {
			return types.FollowListResp{}, codes.CodeAllIntervalError, err
		}
		if userinfo.Active == global.Active.String() {
			briefffs = append(briefffs, types.BriefFollow{
				Nickname:     userinfo.Nickname,
				ID:           userinfo.ID,
				Signature:    userinfo.Signature,
				FollowStatus: u.Status,
				AvatarUrl:    global.GetUrl("user/auth/avatar", userinfo.ID),
			})
		}
	}
	return briefffs, codes.CodeAllSuccess, nil
}

// 用户列表
func LogicFansList(c *gin.Context, req *types.FansListReq) (resp types.FansListResp, code int, err error) {
	id, ex := c.Get("id")
	if !ex {
		return types.FansListResp{}, codes.CodeAuthUnvalidToken, nil
	}
	if req.Order == "" {
		return types.FansListResp{}, codes.CodeAllRequestFormatError, nil
	}
	uid := id.(uint)
	followfans, err := model.FollowFansList(c, uid, req.Page, req.Size, req.Desc, req.Order, false)
	if err != nil {
		return types.FansListResp{}, codes.CodeAllIntervalError, err
	}

	briefffs := make([]types.BriefFans, 0, req.Size)
	for _, u := range followfans {
		userinfo, err := model.FindUserById(c, u.UserID)
		if userinfo == nil {
			return types.FansListResp{}, codes.CodeUserNotExist, nil
		}
		if err != nil {
			return types.FansListResp{}, codes.CodeAllIntervalError, err
		}
		if userinfo.Active == global.Active.String() {
			briefffs = append(briefffs, types.BriefFans{
				Nickname:  userinfo.Nickname,
				ID:        userinfo.ID,
				Signature: userinfo.Signature,
				AvatarUrl: global.GetUrl("user/auth/avatar", userinfo.ID),
			})
		}
	}
	return briefffs, codes.CodeAllSuccess, nil
}
