package redisprefix

const (
	//HFUTcookie前缀
	PrefixUserCookieKey = "cookie:username"
	//帖子点赞前缀
	PrefixPostLikeKey1 = "like:postid:"
	PrefixPostLikeKey2 = "userid:"
	//评论点赞前缀
	PrefixCommentLikeKey1 = "like:commentid:"
	PrefixCommentLikeKey2 = "userid:"
	//子评论点赞前缀
	PrefixSubCommentLikeKey1 = "like:subcommentid:"
	PrefixSubCommentLikeKey2 = "userid:"
)