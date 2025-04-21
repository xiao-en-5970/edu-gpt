package codes




// 12345举例，1代表模块名称，2345代表模块下的错误码

const(
	//总体错误
	CodeAllSuccess = 10001 //成功
	CodeAllIntervalError = 10002 //服务器内部错误
	CodeAllRequestFormatError = 10003 //请求格式错误
	CodeAllUnknownError = 10004 //未知错误
	CodeAllBadGateway = 10005 //错误网关
	//用户逻辑错误
	CodeUserLoginSuccess = 20001 //登录成功 ,返回用户数据
	CodeUserLoginSchoolAuthError = 20002 //信息门户校验失败
	CodeUserLoginSchoolAuthSuccess = 20003 //信息门户校验成功，跳转注册，输入一堆信息（nick....)之后访问/register
	CodeUserLoginPasswordError = 20004 //用户存在，但是密码错误
	CodeUserLoginSchoolIntervalError  = 20005 //信息门户内部问题，请重试
	CodeUserLoginSchoolUnkonwnError  = 20006 //信息门户未知错误
	CodeUserRegisterSuccess = 20007 //注册成功
	CodeUserNotExist = 20008 //用户不存在
	CodeUserAlreadyExist = 20009 //用户已经存在
	CodeUserInfoUpdateFail = 20010 //用户信息更新失败
	//鉴权错误
	CodeAuthNotExistError = 30001 //未授权
	CodeAuthUnvalidToken = 30002 //无效token
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
		CodeUserLoginSuccess : "登录成功",
		CodeUserLoginSchoolAuthError : "信息门户校验失败",
		CodeUserLoginSchoolAuthSuccess : "信息门户校验成功，跳转注册",
		CodeUserLoginPasswordError : "用户存在，但是密码错误",
		CodeUserLoginSchoolIntervalError : "信息门户内部问题，请重试",
		CodeUserLoginSchoolUnkonwnError  : "信息门户未知错误",
		CodeUserRegisterSuccess : "注册成功",
		CodeUserNotExist :"用户不存在",
		CodeUserAlreadyExist :"用户已经存在",
		CodeUserInfoUpdateFail :"用户信息更新失败",

		//鉴权错误
		CodeAuthNotExistError :"未授权",
		CodeAuthUnvalidToken :"无效token",
	}
)