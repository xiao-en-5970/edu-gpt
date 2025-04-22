package types


type HFUTLoginReq struct{
	
	Username string `json:"username"`
	Password string `json:"password"`
}
type HFUTLogin struct{
	Cookie 			string `json:"cookie"`
	OneLoginCookies string `json:"oneLoginCookie"`
}
type HFUTLoginResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data HFUTLogin `json:"data"`
}