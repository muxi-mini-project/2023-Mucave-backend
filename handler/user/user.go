package user

import (
	"Mucave/handler"
	"Mucave/model"
	qiniu "Mucave/pkg"
	"Mucave/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 登录验证
// @Description  通过学号密码验证身份(username,password)
// @Tags login
// @Accept  application/json
// @Produce application/json
// @Param  object  body handler.LoginData true "登录需要的信息"
// @Success 200 {object} handler.Response "{"msg":"登录成功，获得token."}"
// @Failure 400 {object} handler.Error  "{"msg":""Invalid username or password}"
// @Router /login [POST]
func Login(c *gin.Context) {
	var loginData handler.LoginData
	err1 := c.ShouldBind(&loginData)
	username := loginData.Username
	password := loginData.Password
	s := service.LoginRequest(username, password)
	if s == "" {
		c.JSON(401, gin.H{
			"message": "Invalid username or password",
		})
	} else {
		UserId, err2 := model.QueryId(username)
		if err1 != nil || err2 != nil {
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

// @Summary  关注
// @Description 通过关注用户id和被关注用户id建立关注关系
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param followed_id query string true "被关注的用户的id"
// @Success 200 {object} handler.Response "{"msg":"关注成功"}"
// @Failure 400 {object} handler.Error  "{"msg":"关注失败"}"
// @Router /user/following [POST]
func Follow(c *gin.Context) {
	followed_id, _ := strconv.Atoi(c.Query("followed_id"))
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

// @Summary  取消关注
// @Description  通过关注用户id和被关注用户id删除关注关系
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  followed_id  query string true "被关注者id"
// @Success 200 {object} handler.Response "{"msg":"取消关注成功"}"
// @Failure 400 {object} handler.Error  "{"msg":"取消关注失败"}"
// @Router /user/following [DELETE]
func UnFollow(c *gin.Context) {
	err := model.DeleteRelationship(c.Query("followed_id"), service.GetId(c))
	if err != nil {
		handler.SendError(c, 400, "取消关注失败.")
		return
	}
	handler.SendResponse(c, "取消关注成功", nil)
}

// @Summary 我的大致信息
// @Description  通过id获取我的大致信息：User,动态数，关注数，粉丝数.
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} handler.Response "{"msg":"数据为:user,动态数,关注数,粉丝数."}"
// @Router /user/my_outline [GET]
func Outline(c *gin.Context) {
	id, _ := c.Get("UserId")
	outline := model.QueryOutline(id)
	handler.SendResponse(c, "数据为:user,动态数,关注数,粉丝数.", outline)
}

// @Summary 我的帖子
// @Description   查询到我所有发布过的帖子信息
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []handler.Post "{"msg":"查询到自己的帖子"}"
// @Failure 410 {object} handler.Error  "{"msg":"查询自己帖子失败"}"
// @Router /user/my_post [GET]
func MyPost(c *gin.Context) {
	id, _ := c.Get("UserId")
	posts, err := model.QueryOneUserPosts(id)
	if err != nil {
		handler.SendError(c, 410, "查询自己帖子失败.")
		return
	}
	handler.SendResponse(c, "查询到自己的帖子.", posts)
}

// @Summary 我的关注
// @Description  查询我关注的用户
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []handler.User "{"msg":"查询到自己关注的用户"}"
// @Failure 410 {object} handler.Error  "{"msg":"查询自己关注的用户失败"}"
// @Router /user/my_following [GET]
func Following(c *gin.Context) {
	id, _ := c.Get("UserId")
	users, err := model.QueryFollowing(id)
	if err != nil {
		handler.SendError(c, 410, "查询自己关注的用户失败.")
		return
	}
	handler.SendResponse(c, "查询到自己关注的用户", users)
}

// @Summary 我的粉丝
// @Description 查询到关注我的用户
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []handler.User "{"msg":"查询到自己的粉丝."}"
// @Failure 410 {object} handler.Error  "{"msg":"查询自己的粉丝失败."}"
// @Router /user/my_followers [GET]
func Followers(c *gin.Context) {
	id, _ := c.Get("UserId")
	users, err := model.QueryFollowers(id)
	if err != nil {
		handler.SendError(c, 410, "查询自己的粉丝失败.")
		return
	}
	handler.SendResponse(c, "查询到自己的粉丝", users)
}

// @Summary 指定用户的大致信息
// @Description  指定id用户的User,动态数,关注数,粉丝数.
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id  path string true "指定用户的id"
// @Success 200 {object} handler.Response "{"msg":"数据为:user,动态数,关注数,粉丝数."}"
// @Router /user/{id}/user_outline [GET]
func UserOutline(c *gin.Context) {
	id := c.Param("id")
	outline := model.QueryOutline(id)
	outline.User.UserName = "" //change
	handler.SendResponse(c, "数据为:user,动态数,关注数,粉丝数.", outline)
}

// @Summary 指定用户的帖子
// @Description  查询到指定id用户的所有帖子
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id  path string true "指定用户的id"
// @Success 200 {object} []handler.Post "{"msg":"查询到他人的帖子."}"
// @Failure 410 {object} handler.Error  "{"msg":"查询他人的帖子失败."}"
// @Router /user/{id}/user_post [GET]
func UserPost(c *gin.Context) {
	id := c.Param("id")
	posts, err := model.QueryOneUserPosts(id)
	if err != nil {
		handler.SendError(c, 410, "查询他人的帖子失败.")
		return
	}
	handler.SendResponse(c, "查询到他人的帖子.", posts)
}

// @Summary 用户的粉丝
// @Description  查询指定id用户的粉丝列表
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id  path string true "指定用户的id"
// @Success 200 {object} []handler.User "{"msg":"查询到他人的粉丝."}"
// @Failure 410 {object} handler.Error  "{"msg":"查询他人的粉丝失败."}"
// @Router /user/{id}/user_followers [GET]
func UserFollowers(c *gin.Context) {
	id := c.Param("id")
	users, err := model.QueryFollowers(id)
	if err != nil {
		handler.SendError(c, 410, "查询他人的粉丝失败.")
		return
	}
	handler.SendResponse(c, "查询到他人的粉丝.", users)
}

// @Summary 用户的关注
// @Description  指定id用户的关注列表
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id  path string true "指定用户的id"
// @Success 200 {object} []handler.User "{"msg":"查询到他人关注的用户."}"
// @Failure 410 {object} handler.Error  "{"msg":"查询他人关注的用户失败."}"
// @Router /user/{id}/user_following [GET]
func UserFollowing(c *gin.Context) {
	id := c.Param("id")
	users, err := model.QueryFollowing(id)
	if err != nil {
		handler.SendError(c, 410, "查询他人关注的用户失败.")
		return
	}
	handler.SendResponse(c, "查询到他人关注的用户.", users)
}

// @Summary 用户信息
// @Description 查询指定id用户的信息
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id  path string true "指定用户的id"
// @Success 200 {object} handler.User "{"msg":"查询到指定的用户."}"
// @Failure 410 {object} handler.Error  "{"msg":"查询指定用户信息失败."}"
// @Router /user/{id}/user_msg [GET]
func UserMsg(c *gin.Context) {
	id := c.Param("id")
	user, err := model.QueryIdUser(id)
	user.UserName = "" //改
	if err != nil {
		handler.SendError(c, 410, "查询指定用户信息失败.")
		return
	}
	handler.SendResponse(c, "查询到指定的用户.", user)
}

// @Summary 我的信息
// @Description 根据token中的id查我的User体
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} handler.User "{"msg":"查询到自己的信息."}"
// @Failure 410 {object} handler.Error  "{"msg":"查询自己信息失败."}"
// @Router /user/my_msg [GET]
func MyMsg(c *gin.Context) {
	id, _ := c.Get("UserId")
	user, err := model.QueryIdUser(id)
	if err != nil {
		handler.SendError(c, 410, "查询自己信息失败.")
		return
	}
	handler.SendResponse(c, "查询到自己的信息.", user)
}

// @Summary 发私信
// @Description 通过id指定发私信的对象进行发私信
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id  path string true "指定用户的id"
// @Param  file  formData file false "一个文件"
// @Param  file_have  query string true "yes/no(说明是否附带文件)"
// @Param  content_have  query string true "yes/no(说明是否有文本)"
// @Param  content  formData string false "私信的文本内容"
// @Success 200 {object} handler.Response "{"msg":"发送私信成功."}"
// @Failure 400 {object} handler.Error  "{"msg":"发送私信失败."}"
// @Router /user/private_msg/{id} [POST]
func PrivateMsgSend(c *gin.Context) {
	privateMsg := service.MsgTmp(c)
	if c.Query("file_have") == "yes" {
		urls := make([]string, 1, 1)
		urls, _ = qiniu.UploadFile(c)
		privateMsg.FilePath = urls[0]
	}
	if c.Query("content_have") != "no" {
		privateMsg.Content = c.PostForm("content")
	}
	err := model.CreatePrivateMsg(privateMsg)
	if err != nil {
		handler.SendError(c, 400, "发送私信失败.")
		return
	}
	handler.SendResponse(c, "发送私信成功.", nil)
}

// @Summary 刷新指定私信
// @Description  通过id刷新指定用户的向自己发的信息
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id  path string true "指定用户的id"
// @Success 200 {object} handler.Response "{"msg":"刷新私信成功"}"
// @Failure 410 {object} handler.Error  "{"msg":"刷新私信失败"}"
// @Router /user/private_msg/{id} [GET]
func PrivateMsg(c *gin.Context) {
	msgs, err := model.QueryPrivateMsg(c.Param("id"), service.GetId(c))
	if err != nil {
		handler.SendError(c, 410, "刷新指定私信失败.")
		return
	}
	handler.SendResponse(c, "刷新指定私信成功.", msgs)
}

// @Summary 刷新所有私信
// @Description  刷新所有发向我的私信
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} handler.Response "{"msg":"刷新所有私信成功"}"
// @Failure 410 {object} handler.Error  "{"msg":"刷新所有私信失败"}"
// @Router /user/all_private_msg [GET]
func AllPrivateMsg(c *gin.Context) {
	msgs, err := model.QueryPrivateMsg("all", service.GetId(c))
	m := make(map[uint][]model.PrivateMsg)
	for _, msg := range msgs {
		m[msg.SenderId] = append(m[msg.SenderId], msg)
	}
	if err != nil {
		handler.SendError(c, 410, "刷新所有私信失败.")
		return
	}
	handler.SendResponse(c, "刷新所有私信成功.(每个对象的信息在一个数组里ID:[])", m)
}

// @Summary 修改我的信息
// @Description  上传我的新User更新我的信息
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param object  body handler.User false "除了头像地址之外的其他信息"
// @Success 200 {object} handler.Response "{"msg":"用户信息修改成功."}"
// @Failure 400 {object} handler.Error  "{"msg":"用户信息修改失败."}"
// @Router /user/my_msg [PUT]
func MyMsgUpdate(c *gin.Context) {
	if c.Query("avatar_only") == "yes" {
		urls, _ := qiniu.UploadFile(c)
		path := urls[0]
		err := model.UpdateAvatar(service.GetId(c), path)
		if err != nil {
			handler.SendError(c, 400, "用户头像修改失败.")
			return
		}
		handler.SendResponse(c, "用户头像修改成功.", nil)
		return
	}
	var newUserMsg model.User
	err1 := c.ShouldBind(&newUserMsg)
	err2 := model.UpdateUserMsg(newUserMsg, service.GetId(c))
	if err1 != nil || err2 != nil {
		handler.SendError(c, 400, "用户信息修改失败.")
		return
	}
	handler.SendResponse(c, "用户信息修改成功.", nil)
}

// @Summary 我的评论
// @Description  查询我的所有评论
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []handler.Comment "{"msg":"查询到我的评论."}"
// @Failure 410 {object} handler.Error  "{"msg":"我的评论查询失败"}"
// @Router /user/my_comments [GET]
func MyComments(c *gin.Context) {
	id, _ := c.Get("UserId")
	comments, err := model.QueryMyComments(id)
	if err != nil {
		handler.SendError(c, 410, "我的评论查询失败.")
		return
	}
	handler.SendResponse(c, "查询到我的评论.", comments)
}

// @Summary 我的回复
// @Description  查询我的所有回复
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []handler.Reply "{"msg":"查询到我的回复."}"
// @Failure 410 {object} handler.Error  "{"msg":"我的回复查询失败."}"
// @Router /user/my_replies [GET]
func MyReplies(c *gin.Context) {
	id, _ := c.Get("UserId")
	replies, err := model.QueryMyReplies(id)
	if err != nil {
		handler.SendError(c, 410, "我的回复查询失败.")
		return
	}
	handler.SendResponse(c, "查询到我的回复", replies)
}

// @Summary 我的点赞
// @Description 我点赞的所有的帖子
// @Tags user
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} []handler.Post "{"msg":"查询到我的点赞的帖子."}"
// @Failure 410 {object} handler.Error  "{"msg":"我点赞过的帖子查询失败."}"
// @Router /user/my_likes [GET]
func MyLikesPost(c *gin.Context) {
	id, _ := c.Get("UserId")
	posts, err := model.QueryMyLikesPosts(id)
	if err != nil {
		handler.SendError(c, 410, "我点赞过的帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到我点赞的帖子", posts)
}

// @Summary  是否关注
// @Description 通过用户id查询是否已经关注
// @Tags post
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Param user_id query integer true "用户id"
// @Success 200 {object} handler.Response "{"msg":"no"}"
// @Success 200 {object} handler.Response "{"msg":"yes"}"
// @Router /user/whether_follow [GET]
func WhetherFollow(c *gin.Context) {
	follow := model.WhetherFollow(c.Query("user_id"), service.GetId(c))
	if !follow {
		handler.SendResponse(c, "no", nil)
		return
	}
	handler.SendResponse(c, "yes", nil)
}

// @Summary  点赞通知
// @Description  刷新后获得我的贴子的点赞信心
// @Tags post
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Failure 410 {object} handler.Error "{"msg":"刷新点赞我的失败."}"
// @Success 200 {object} []handler.Likes "{"msg":"刷新点赞我的成功."}"
// @Router /user/likes_of_my_post [GET]
func LikesOfMyPosts(c *gin.Context) {
	posts, _ := model.QueryOneUserPosts(service.GetId(c))
	likes, err := model.QueryLikesSend(posts)
	if err != nil {
		handler.SendError(c, 410, "刷新点赞我的失败.")
		return
	}
	handler.SendResponse(c, "刷新点赞我的成功.", likes)
}

// @Summary  回复通知
// @Description  获得我的贴子下面的评论和回复所有新增信息和他人贴子下面回复对象是我的信息
// @Tags post
// @Accept  application/json
// @Produce application/json
// @Param Authorization header string true "token"
// @Failure 410 {object} handler.Error "{"msg":"查询回复失败."}"
// @Success 200 {object} handler.ReCo "{"msg":"查询回复成功."}"
// @Router /user/replies [GET]
func RepliesToMe(c *gin.Context) {
	posts, _ := model.QueryOneUserPosts(service.GetId(c))
	comments, err1 := model.QueryCommentSend(posts)
	replies, err2 := model.QueryRepliesSend(posts, service.GetId(c))
	if err1 != nil || err2 != nil {
		handler.SendError(c, 410, "查询回复失败.")
		return
	}
	handler.SendResponse(c, "查询回复成功.", gin.H{
		"comments": comments,
		"replies":  replies,
	})
}
