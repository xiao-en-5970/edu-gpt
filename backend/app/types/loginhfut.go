package types


type LoginHFUTReq struct{
	
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginHFUTResp struct{
	Cookie 			string `json:"cookie"`
	OneLoginCookies string `json:"oneLoginCookie"`
}
