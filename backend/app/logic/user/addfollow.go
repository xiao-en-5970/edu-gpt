package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"github.com/xiao-en-5970/edu-gpt/backend/app/models"
	types "github.com/xiao-en-5970/edu-gpt/backend/app/types/user"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/codes"
)

func LogicUserAddFollow(c *gin.Context, req *types.UserAddFollowReq) (resp types.UserAddFollowResp, code int, err error) {
	u, ex := c.Get("id")
	if !ex {
		return resp, codes.CodeAuthUnvalidToken, nil
	}
	uid := u.(uint)
	user, _ := models.FindUserById(c, req.FollowID)
	if user == nil || user.Active != global.Active.String() {
		return resp, codes.CodeUserNotExist, nil
	}
	resp.OK = 1
	go models.AddFollows(c, uid, req.FollowID, req.FollowStatus)
	return resp, codes.CodeAllSuccess, nil
}
