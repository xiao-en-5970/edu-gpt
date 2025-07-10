package redisprefix

const (
	//HFUTcookie前缀
	PrefixUserCookieKey = "cookie:username"
	//帖子点赞前缀
	PrefixPostLikeKey1 = "userlikepost:postid"
	PrefixPostLikeKey2 = "userid"
	//帖子点赞总数
	PrefixPostLikeCountKey = "likecount:postid"

	//评论点赞前缀
	PrefixCommentLikeKey1 = "userlikecomment:commentid"
	PrefixCommentLikeKey2 = "userid"
	//子评论点赞前缀
	PrefixSubCommentLikeKey1 = "userlikesubcomment:subcommentid"
	PrefixSubCommentLikeKey2 = "userid"
)