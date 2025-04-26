package codes




// 12345举例，1代表模块名称，2345代表模块下的错误码

const(
	//总体错误
	CodeAllSuccess = 10000 //成功
	CodeAllIntervalError = 10001 //服务器内部错误
	CodeAllRequestFormatError = 10002 //请求格式错误
	CodeAllUnknownError = 10003 //未知错误
	CodeAllBadGateway = 10004 //错误网关
	

	//用户逻辑错误
	CodeUserLoginPasswordError = 20001 //用户存在，但是密码错误
	CodeUserNotExist = 20002 //用户不存在
	CodeUserAlreadyExist = 20003 //用户已经存在
	CodeUserInfoUpdateFail = 20004 //用户信息更新失败
	//鉴权错误
	CodeAuthNotExistError = 30001 //未授权
	CodeAuthUnvalidToken = 30002 //无效token
	//HFUT-api错误
	CodeHFUTLoginError 			= 40001 //信息门户校登录失败，重新登录
	CodeHFUTIntervalError  		= 40002 //信息门户内部问题，请重试
	CodeHFUTUnkonwnError  		= 40003 //信息门户未知错误
	CodeHFUTNotLogin 			= 40004 //信息门户未登录
	//图片错误
	CodeImageFormatError  		= 50001 //图片格式错误
	//帖子错误
	CodePostNotExist = 60001 //帖子不存在
)
var(
	CodeMsg = map[int]string{
		//总体错误
		CodeAllSuccess :"成功",
		CodeAllIntervalError : "服务器内部错误",
		CodeAllRequestFormatError : "请求格式错误",
		CodeAllUnknownError : "未知错误",
		CodeAllBadGateway : "错误网关",
		//用户逻辑错误
		CodeUserLoginPasswordError : "用户存在，但是密码错误",
		CodeUserNotExist :"用户不存在",
		CodeUserAlreadyExist :"用户已经存在",
		CodeUserInfoUpdateFail :"用户信息更新失败",
		//鉴权错误
		CodeAuthNotExistError :"未授权",
		CodeAuthUnvalidToken :"无效token",
		//HFUT-api错误
		CodeHFUTLoginError 	:"信息门户校登录失败",
		CodeHFUTIntervalError :"信息门户内部问题，请重试",
		CodeHFUTUnkonwnError :"信息门户未知错误",
		CodeHFUTNotLogin :"信息门户未登录",
		//图片错误
		CodeImageFormatError :"图片格式错误",
		//帖子错误
		CodePostNotExist :"帖子不存在",
	}
)