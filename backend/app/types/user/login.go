package types


// LoginReq 用户登录请求参数
// @Description 用户登录的请求参数
type LoginReq struct {
    Username string `json:"username" binding:"required,alphanum,min=4,max=20" example:"2022210857"`
    Password string `json:"password" binding:"required,min=6" example:"QQQQQAQWERqwer2."`
}

// LoginResp 用户登录响应数据
// @Description 用户登录成功后返回的响应数据
type LoginResp struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc5ODgxMjgsImlkIjoxfQ.-h4HvQ9PB15dxE-gDXQnSIY0h9ambb1AbMc0VFjZLTE"`
    ID    uint   `json:"id" example:"1"`
}



