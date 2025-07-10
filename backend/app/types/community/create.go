package types

type CommunityCreateReq struct {
	Name        string `json:"name"`         // 社区名称
	Description string `json:"description"`  // 社区描述
}

type CommunityCreateResp struct {
	ID uint `json:"id"` // 创建的社区ID
}