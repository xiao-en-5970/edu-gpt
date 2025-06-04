package types


type UserAddFollowReq struct{
	FollowID     uint `json:"follow_id" validate:"required" example:"123" comment:"帖子ID"`
	FollowStatus int  `json:"follow_status" validate:"required,oneof=0 1" example:"1" comment:"关注状态(0:取消关注,1:关注)"`
}

type UserAddFollowResp struct{
	OK  int `json:"ok" example:"1" comment:"是否成功"`
}