package types

// LoginReq 登录请求参数
// swagger:parameters login
type LoginReq struct {

    Username string `json:"username" binding:"required,alphanum,min=4,max=20"`

    Password string `json:"password" binding:"required,min=6"`
}

// LoginResp 登录响应数据
// swagger:model LoginResp
type LoginResp struct{
	Token string `json:"token"`
    ID uint `json:"id"`
}