package user

import (
	"Mucave/handler"
	"Mucave/model"
	"Mucave/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	s := service.LoginRequest(username, password)
	if s == "" {
		c.JSON(401, gin.H{
			"message": "Invalid username or password",
		})
	} else {
		UserId, err := model.QueryId(username)
		if err != nil {
			user := model.CreateUser(username)
			UserId = user.ID
		}
		tokenString, _ := service.CreateToken(UserId)
		c.JSON(200, gin.H{
			"msg":   "登录成功,获得token.",
			"token": tokenString,
		})
	}
}
func Follow(c *gin.Context) {
	followed_id, _ := strconv.Atoi(c.PostForm("followed_id"))
	relationship := model.Relationship{
		FollowerId: service.GetId(c),
		FollowedId: uint(followed_id),
	}
	err := model.CreateRelationship(relationship)
	if err != nil {
		handler.SendError(c, 400, "关注失败.")
		return
	}
	handler.SendResponse(c, "关注成功", nil)
}
func UnFollow(c *gin.Context) {
	followed_id, _ := strconv.Atoi(c.Query("followed_id"))
	relationship := model.Relationship{
		FollowerId: service.GetId(c),
		FollowedId: uint(followed_id),
	}
	err := model.DeleteRelationship(relationship)
	if err != nil {
		handler.SendError(c, 400, "取消关注失败.")
		return
	}
	handler.SendResponse(c, "取消关注成功", nil)
}
func Outline(c *gin.Context) {
	id, _ := c.Get("UserId")
	outline := model.QueryOutline(id)
	handler.SendResponse(c, "数据为:user,动态数,关注数,粉丝数.", outline)
}
func MyPost(c *gin.Context) {
	id, _ := c.Get("UserId")
	user, _ := model.QueryIdUser(id)
	users := []model.User{user}
	posts, err := model.QueryUserPosts(users)
	if err != nil {
		handler.SendError(c, 410, "查询自己帖子失败.")
		return
	}
	handler.SendResponse(c, "查询到自己的帖子.", posts)
}
func Following(c *gin.Context) {
	id, _ := c.Get("UserId")
	users, err := model.QueryFollowing(id)
	if err != nil {
		handler.SendError(c, 410, "查询自己关注的用户失败.")
		return
	}
	handler.SendResponse(c, "查询到自己关注的用户", users)
}
func Followers(c *gin.Context) {
	id, _ := c.Get("UserId")
	users, err := model.QueryFollowers(id)
	if err != nil {
		handler.SendError(c, 410, "查询自己的粉丝失败.")
		return
	}
	handler.SendResponse(c, "查询到自己的粉丝", users)
}
func UserOutline(c *gin.Context) {
	id := c.Param("id")
	outline := model.QueryOutline(id)
	handler.SendResponse(c, "数据为:user,动态数,关注数,粉丝数.", outline)
}
func UserPost(c *gin.Context) {
	id := c.Param("id")
	user, _ := model.QueryIdUser(id)
	users := []model.User{user}
	posts, err := model.QueryUserPosts(users)
	if err != nil {
		handler.SendError(c, 410, "查询他人的帖子失败.")
		return
	}
	handler.SendResponse(c, "查询到他人的帖子.", posts)
}
func UserFollowers(c *gin.Context) {
	id := c.Param("id")
	users, err := model.QueryFollowers(id)
	if err != nil {
		handler.SendError(c, 410, "查询他人的粉丝失败.")
		return
	}
	handler.SendResponse(c, "查询到他人的粉丝", users)
}
func UserFollowing(c *gin.Context) {
	id := c.Param("id")
	users, err := model.QueryFollowing(id)
	if err != nil {
		handler.SendError(c, 410, "查询他人关注的用户失败.")
		return
	}
	handler.SendResponse(c, "查询到他人关注的用户", users)
}
func UserMsg(c *gin.Context) {
	id := c.Param("id")
	user, err := model.QueryIdUser(id)
	if err != nil {
		handler.SendError(c, 410, "查询指定用户信息失败.")
		return
	}
	handler.SendResponse(c, "查询到指定的用户.", user)
}
func MyMsg(c *gin.Context) {
	id, _ := c.Get("UserId")
	user, err := model.QueryIdUser(id)
	if err != nil {
		handler.SendError(c, 410, "查询自己信息失败.")
		return
	}
	handler.SendResponse(c, "查询到自己的信息", user)
}
func PrivateMsgSend(c *gin.Context) {
	sender, _ := c.Get("UserId")
	receiver := c.Param("id")
	SenderId, _ := sender.(uint)
	senderId := strconv.Itoa(int(SenderId))
	path := service.UploadFile(c, senderId+"-"+receiver)
	sendID, _ := strconv.Atoi(senderId)
	receiverId, _ := strconv.Atoi(receiver)
	privateMsg := model.PrivateMsg{
		SenderId:   sendID,
		ReceiverId: receiverId,
		Status:     1,
		Content:    c.PostForm("content"),
		FilePath:   path,
		SendTime:   time.Now(),
	}
	err := model.CreatePrivateMsg(privateMsg)
	if err != nil {
		handler.SendError(c, 400, "发送私信失败.")
		return
	}
	handler.SendResponse(c, "发送私信成功.", nil)
}
func PrivateMsg(c *gin.Context) {
	senderId := c.Param("id")
	receiverId, _ := c.Get("UserId")
	msgs, err := model.QueryPrivateMsg(senderId, receiverId)
	if err != nil {
		handler.SendError(c, 410, "刷新私信失败.")
		return
	}
	handler.SendResponse(c, "刷新私信成功.", msgs)
}
func MyMsgUpdate(c *gin.Context) {
	path := service.UploadFile(c, service.GetId(c))
	newUserMsg := model.User{
		Model:      gorm.Model{ID: service.GetId(c)},
		Name:       c.PostForm("name"),
		Gender:     c.PostForm("gender"),
		Signature:  c.PostForm("signature"),
		AvatarPath: path,
		Birthday:   time.Now(),
		Hometown:   c.PostForm("hometown"),
		Grader:     c.PostForm("grader"),
		Faculties:  c.PostForm("faculties"),
	}
	err := model.UpdateUserMsg(newUserMsg)
	if err != nil {
		handler.SendError(c, 400, "用户信息修改失败.")
		return
	}
	handler.SendResponse(c, "用户信息修改成功.", nil)
}
func MyComments(c *gin.Context) {
	id, _ := c.Get("UserId")
	comments, err := model.QueryMyComments(id)
	if err != nil {
		handler.SendError(c, 410, "我的评论查询失败.")
		return
	}
	handler.SendResponse(c, "查询到我的评论.", comments)
}
func MyReplies(c *gin.Context) {
	id, _ := c.Get("UserId")
	replies, err := model.QueryMyReplies(id)
	if err != nil {
		handler.SendError(c, 410, "我的回复查询失败.")
		return
	}
	handler.SendResponse(c, "查询到我的回复", replies)
}
func MyLikesPost(c *gin.Context) {
	id, _ := c.Get("UserId")
	posts, err := model.QueryMyLikesPosts(id)
	if err != nil {
		handler.SendError(c, 410, "我点赞过的帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到我点赞的帖子", posts)
}
