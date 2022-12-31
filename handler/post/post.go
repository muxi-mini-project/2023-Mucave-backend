package post

import (
	"Mucave/handler"
	"Mucave/model"
	"Mucave/service"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func Latest(c *gin.Context) {
	posts, err := model.QueryNewPosts(c.Query("start"), c.Query("length"))
	if err != nil {
		handler.SendError(c, 410, "最新的帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到最新帖子数组.", posts)
}
func Recommendations(c *gin.Context) {
	ty := c.Param("type")
	posts, err := model.QueryHotPosts(c.Query("start"), c.Query("length"), ty)
	if err != nil {
		handler.SendError(c, 410, "推荐帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到点赞多的的帖子数组", posts)
}
func Following(c *gin.Context) {
	UserId, ok := c.Get("UserId")
	if !ok {
		c.JSON(500, gin.H{"message": "无法识别用户身份信息"})
		return
	}
	users, err := model.QueryFollowing(UserId)
	if err != nil {
		handler.SendError(c, 410, "查询关注的用户失败.")
		return
	}
	posts, err := model.QueryUserPosts(users)
	if err != nil {
		handler.SendError(c, 410, "查询关注用户的帖子失败.")
		return
	}
	handler.SendResponse(c, "查询到关注的用户的帖子的数组", posts)
}
func QueryOnePosts(c *gin.Context) {
	id := c.Param("id")
	post, err := model.QueryIdPost(id)
	if err != nil {
		handler.SendError(c, 410, "指定的帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到指定的帖子", post)
}
func Comments(c *gin.Context) {
	postId := c.Param("postId")
	comments, err := model.QueryCommentByPostId(postId)
	if err != nil {
		handler.SendError(c, 410, "帖子的评论查询失败.")
		return
	}
	handler.SendResponse(c, "查询到帖子的评论", comments)
}
func Reply(c *gin.Context) {
	commentId := c.Param("commentId")
	replies, err := model.QueryReplyByCommentId(commentId)
	if err != nil {
		handler.SendError(c, 410, "查询评论的回复失败.")
		return
	}
	handler.SendResponse(c, "查询到评论的回复", replies)
}
func CreatePost(c *gin.Context) {
	post := model.Post{
		AuthorId:  service.GetId(c),
		Type:      c.PostForm("type"),
		Content:   c.PostForm("content"),
		Likes:     0,
		CommentNo: 0,
	}
	post, err := model.CreatePost(post)
	if err != nil {
		handler.SendError(c, 400, "新建帖子失败.")
		return
	}
	DirPath := service.UploadFile(c, post.ID)
	files, _ := os.ReadDir(DirPath)
	paths := make([]string, 4, 4)
	for i, file := range files {
		FilePath := DirPath + "/" + file.Name()
		paths[i] = FilePath
	}
	post.FilePath1 = paths[0]
	post.FilePath2 = paths[1]
	post.FilePath3 = paths[2]
	post.FilePath4 = paths[3]
	model.DB.Save(&post)
	handler.SendResponse(c, "发布成功", gin.H{"data": "null"})
}
func AddReply(c *gin.Context) {
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	object, _ := strconv.Atoi(c.PostForm("object"))
	reply := model.Reply{
		FromWho:   service.GetId(c),
		CommentId: uint(commentId),
		Object:    uint(object),
		Content:   c.PostForm("content"),
	}
	err := model.CreateReply(reply)
	if err != nil {
		handler.SendError(c, 400, "回复失败.")
		return
	}
	handler.SendResponse(c, "回复成功.", nil)
}
func AddComments(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("postId"))
	comment := model.Comment{
		PostId:  uint(postId),
		UserId:  service.GetId(c),
		Content: c.PostForm("content"),
	}
	err := model.CreateComment(comment)
	if err != nil {
		handler.SendError(c, 400, "评论失败.")
		return
	}
	handler.SendResponse(c, "评论成功.", nil)
}
func AddLikes(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("postId"))
	likes := model.Likes{
		PostId: uint(postId),
		UserId: service.GetId(c),
	}
	err := model.CreateLikes(likes)
	if err != nil {
		handler.SendError(c, 400, "点赞失败.")
		return
	}
	handler.SendResponse(c, "点赞成功", nil)
}
