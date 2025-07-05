package codes

// 12345举例，1代表模块名称，2345代表模块下的错误码

const (
	//总体错误
	CodeAllSuccess            = 10000 //成功
	CodeAllIntervalError      = 10001 //服务器内部错误
	CodeAllRequestFormatError = 10002 //请求格式错误
	CodeAllUnknownError       = 10003 //未知错误
	CodeAllBadGateway         = 10004 //错误网关

	//用户逻辑错误
	CodeUserLoginPasswordError  = 20001 //用户存在，但是密码错误
	CodeUserNotExist            = 20002 //用户不存在
	CodeUserAlreadyExist        = 20003 //用户已经存在
	CodeUserInfoUpdateFail      = 20004 //用户信息更新失败
	CodeUserRefreshHFUTInfoFail = 20005 //用户HFUT信息无法刷新
	CodeUserLocked              = 20006 // 用户已锁定
	CodeUserLDisabled           = 20007 // 用户已注销
	//鉴权错误
	CodeAuthNotExistError = 30001 //未授权
	CodeAuthUnvalidToken  = 30002 //无效token
	//HFUT-api错误
	CodeHFUTLoginError    = 40001 //信息门户校登录失败，重新登录
	CodeHFUTIntervalError = 40002 //信息门户内部问题，请重试
	CodeHFUTUnkonwnError  = 40003 //信息门户未知错误
	CodeHFUTNotLogin      = 40004 //信息门户未登录
	CodeHFUTRetrySoMuch   = 40005 //信息门户重试次数过多，请稍后再试

	//图片错误
	CodeImageFormatError = 50001 //图片格式错误
	CodeImageNotExist    = 50002 //图片不存在
	//帖子错误
	CodePostNotExist       = 60001 //帖子不存在
	CodePostLikeFail       = 60002 //帖子点赞失败
	CodePostLikeStatusSame = 60003 //帖子点赞状态不变
	CodePostLocked         = 60004 // 帖子已锁定
	CodePostLDisabled      = 60005 // 帖子已删除
	//评论错误
	CodeCommentNotExist     = 70001 //评论不存在
	CodeSubCommentNotExist  = 70002 //子评论不存在
	CodeCommentLocked       = 70003 // 评论已锁定
	CodeCommentLDisabled    = 70004 // 评论已删除
	CodeSubCommentLocked    = 70005 // 回复已锁定
	CodeSubCommentLDisabled = 70006 // 回复已删除
	//课程错误
	CodeCourseNotExist     = 80001 //课程不存在
	CodeCourseAlreadyExist = 80002 //课程已存在
	CodeCourseNotFound     = 80003 //课程未找到
	//文件错误
	CodeFileQueryFail          = 90001 //文件查询失败
	CodeFileNotExist           = 90002 //文件不存在
	CodeFileUploadFail         = 90003 //文件上传失败
	CodeFileDownloadFail       = 90004 //文件下载失败
	CodeFileDeleteFail         = 90005 //文件删除失败
	CodeFileQueryFilePathEmpty = 90006 //文件查询路径不允许为空
	//课程表错误
	CodeCourseTableQueryFail   = 100001 //课程表查询失败
	CodeCourseTableCampusFault = 100002 //课程表校区错误
	CodeCourseTableYearFault   = 100003 //课程表学年错误
	CodeCourseTableTermFault   = 100004 //课程表学期错误
	CodeCourseTableEmpty       = 100005 //课程表为空
)

var (
	CodeMsg = map[int]string{
		//总体错误
		CodeAllSuccess:            "成功",
		CodeAllIntervalError:      "服务器内部错误",
		CodeAllRequestFormatError: "请求格式错误",
		CodeAllUnknownError:       "未知错误",
		CodeAllBadGateway:         "错误网关",
		//用户逻辑错误
		CodeUserLoginPasswordError:  "用户存在，但是密码错误",
		CodeUserNotExist:            "用户不存在",
		CodeUserAlreadyExist:        "用户已经存在",
		CodeUserInfoUpdateFail:      "用户信息更新失败",
		CodeUserRefreshHFUTInfoFail: "用户HFUT信息无法刷新",
		CodeUserLocked:              "用户已锁定",
		CodeUserLDisabled:           "用户已注销",
		//鉴权错误
		CodeAuthNotExistError: "未授权",
		CodeAuthUnvalidToken:  "无效token",
		//HFUT-api错误
		CodeHFUTLoginError:    "信息门户校登录失败",
		CodeHFUTIntervalError: "信息门户内部问题，请重试",
		CodeHFUTUnkonwnError:  "信息门户未知错误",
		CodeHFUTNotLogin:      "信息门户未登录",
		CodeHFUTRetrySoMuch:   "信息门户重试次数过多，请稍后再试",
		//图片错误
		CodeImageFormatError: "图片格式错误",
		CodeImageNotExist:    "图片不存在",
		//帖子错误
		CodePostNotExist:       "帖子不存在",
		CodePostLikeFail:       "帖子点赞失败",
		CodePostLikeStatusSame: "帖子点赞状态不变",
		CodePostLocked:         "帖子已锁定",
		CodePostLDisabled:      "帖子已删除",
		//评论错误
		CodeCommentNotExist:     "评论不存在",
		CodeSubCommentNotExist:  "子评论不存在",
		CodeCommentLocked:       "评论已锁定",
		CodeCommentLDisabled:    "评论已删除",
		CodeSubCommentLocked:    "回复已锁定",
		CodeSubCommentLDisabled: "回复已删除",
		//课程错误
		CodeCourseNotExist:     "课程不存在",
		CodeCourseAlreadyExist: "课程已存在",
		CodeCourseNotFound:     "课程未找到",
		//文件错误
		CodeFileQueryFail:          "文件查询失败",
		CodeFileNotExist:           "文件不存在",
		CodeFileUploadFail:         "文件上传失败",
		CodeFileDownloadFail:       "文件下载失败",
		CodeFileDeleteFail:         "文件删除失败",
		CodeFileQueryFilePathEmpty: "文件查询路径不允许为空",
		//课程表错误
		CodeCourseTableQueryFail:   "课程表查询失败。原因：hfutapi问题或信息门户未登录",
		CodeCourseTableCampusFault: "课程表校区错误。原因：校区2代表合肥校区，23代表宣城校区",
		CodeCourseTableYearFault:   "课程表学年错误。原因：学年不能超过当前年份",
		CodeCourseTableTermFault:   "课程表学期错误。原因：学期为1或2",
		CodeCourseTableEmpty:       "课程表为空。原因：当前学期无课或者当前学期未开始",
	}
)
