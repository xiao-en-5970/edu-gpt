package types

type FollowListReq struct {
	Page  int    `json:"page" example:"1" comment:"页数"`
	Size  int    `json:"size" validate:"required,min=1,max=50" example:"10" comment:"每页数量"`
	Order string `json:"order" validate:"required,oneof=time fans" example:"time" comment:"排序依据"`
	Desc  int    `json:"desc" example:"0" comment:"是否倒序"`
}

type BriefFollow struct {
	ID           uint   `json:"id" example:"1"`
	Nickname     string `json:"nickname" example:"傅益忠"`
	AvatarUrl    string `json:"avatar_url" example:"https://127.0.0.1:8080/api/v1/user/auth/avatar/1"`
	Signature    string `json:"signature" example:"这人啥也没说"`
	FollowStatus int    `json:"follow_status" example:"1"`
}

type FollowListResp []BriefFollow

type FansListReq struct {
	Page  int    `json:"page" example:"1" comment:"页数"`
	Size  int    `json:"size" validate:"required,min=1,max=50" example:"10" comment:"每页数量"`
	Order string `json:"order" validate:"required,oneof=time fans" example:"time" comment:"排序依据"`
	Desc  int    `json:"desc" example:"0" comment:"是否倒序"`
}

type BriefFans struct {
	ID        uint   `json:"id" example:"1"`
	Nickname  string `json:"nickname" example:"傅益忠"`
	AvatarUrl string `json:"avatar_url" example:"https://127.0.0.1:8080/api/v1/user/auth/avatar/1"`
	Signature string `json:"signature" example:"这人啥也没说"`
}

type FansListResp []BriefFans
