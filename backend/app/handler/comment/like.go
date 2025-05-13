package handler

import "github.com/gin-gonic/gin"

// @Summary 评论点赞【TODO】
// @Description 还没做
// @Tags Comment模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param edit body types.CommentLikeReq true "请求体"
// @Success 200 {object} types.CommentLikeResp "成功响应"
// @Router /comment/auth/like [post]
func HandlerCommentLike(c * gin.Context){
	// TODO
	
}
// @Summary 回复点赞【TODO】
// @Description 还没做
// @Tags Comment模块
// @Security     BearerAuth 
// @Accept json
// @Produce json
// @Param edit body types.SubCommentLikeReq true "请求体"
// @Success 200 {object} types.SubCommentLikeResp "成功响应"
// @Router /comment/auth/likereply [post]
func HandlerSubCommentLike(c * gin.Context){
	// TODO
}