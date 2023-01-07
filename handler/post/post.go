package post

import (
	"Mucave/handler"
	"Mucave/model"
	qiniu "Mucave/pkg"
	"Mucave/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 最新的帖子
// @Description  查询最新发布的帖子，返回的开始点和数量由query参数决定
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  start  query integer true "开始点"
// @Param  length query integer true "需要返回的帖子数量"
// @Success 200 {object} handler.Response "{"msg":"查询到最新的帖子（数组）"}"
// @Failure 410 {object} handler.Error  "{"msg":"最新帖子查询失败"}"
// @Router /post/latest [GET]
func Latest(c *gin.Context) {
	posts, err := model.QueryNewPosts(c.Query("start"), c.Query("length"))
	if err != nil {
		handler.SendError(c, 410, "最新的帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到最新帖子（数组）.", posts)
}

// @Summary 推荐的帖子(json)
// @Description  查询点赞最多的帖子，返回的开始点和数量由query参数决定，类型有type参数决定(type,start_time,end_time,start_index,length)
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  type  path string true "帖子类型"
// @Param  start  query integer true "开始点"
// @Param  length query integer true "需要返回的帖子数量"
// @Success 200 {object} handler.Response "{"msg":"查询到推荐的帖子（数组）"}"
// @Failure 410 {object} handler.Error  "{"msg":"推荐帖子查询失败"}"
// @Router /post/recommendations [GET]
func Recommendations(c *gin.Context) {
	m := service.RawToMap(c)
	ty := m["type"]
	startTime := m["start_time"]
	endTime := m["end_time"]
	startIndex := m["start_index"]
	length := m["length"]
	posts, err := model.QueryHotPosts(startIndex, length, ty, startTime, endTime)
	if err != nil {
		handler.SendError(c, 410, "推荐帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到点赞多的的帖子组", posts)
}

// @Summary 关注用户帖子
// @Description  查询关注用户的帖子（所有）
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Success 200 {object} handler.Response "{"msg":"查询到关注的用户的帖子的数组"}"
// @Failure 410 {object} handler.Error  "{"msg":"查询关注的用户失败."}"
// @Failure 410 {object} handler.Error  "{"msg":"查询关注用户的帖子失败."}"
// @Router /post/following [GET]
func Following(c *gin.Context) {
	UserId, _ := c.Get("UserId")
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

// @Summary 查某个帖子
// @Description  通过帖子的id获得某个帖子的详细信息
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  id path integer true "指定帖子的id"
// @Success 200 {object} handler.Response "{"msg":"查询到指定的帖子"}"
// @Failure 410 {object} handler.Error  "{"msg":"指定的帖子查询失败."}"
// @Router /post/{id} [GET]
func QueryOnePosts(c *gin.Context) {
	id := c.Param("id")
	post, err := model.QueryIdPost(id)
	if err != nil {
		handler.SendError(c, 410, "指定的帖子查询失败.")
		return
	}
	handler.SendResponse(c, "查询到指定的帖子", post)
}

// @Summary 帖子的评论
// @Description 根据帖子的id查询指定帖子的所有评论(不包括评论的回复)
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  post_id  path integer true "指定帖子的id"
// @Success 200 {object} handler.Response "{"msg":"查询到帖子的评论."}"
// @Failure 410 {object} handler.Error  "{"msg":"帖子的评论查询失败."}"
// @Router /post/comments/{post_id} [GET]
func Comments(c *gin.Context) {
	postId := c.Param("post_id")
	my := service.WhetherMyPost(service.GetId(c), postId)
	comments, err := model.QueryCommentByPostId(postId, my)
	if err != nil {
		handler.SendError(c, 410, "帖子的评论查询失败.")
		return
	}
	handler.SendResponse(c, "查询到帖子的评论.", comments)
}

// @Summary 评论的回复
// @Description 根据评论的id查询指定评论的所有回复
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  comment_id  path integer true "指定评论的id"
// @Success 200 {object} handler.Response "{"msg":"查询到评论的回复"}"
// @Failure 410 {object} handler.Error  "{"msg":"查询评论的回复失败."}"
// @Router /post/comment_replies/{comment_id} [GET]
func Reply(c *gin.Context) {
	commentId := c.Param("comment_id")
	replies, err := model.QueryReplyByCommentId(commentId, service.GetId(c))
	if err != nil {
		handler.SendError(c, 410, "查询评论的回复失败.")
		return
	}
	handler.SendResponse(c, "查询到评论的回复", replies)
}

// @Summary 发布帖子
// @Description 上传帖子的各项信息，发布帖子
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param type  formData string true "类型"
// @Param title  formData string true "标题"
// @Param content  formData string true "文本内容"
// @Param file  formData file false "文件组(有数量限制)"
// @Param file_have  query string true "yse/no(说明是否上传了文件)"
// @Success 200 {object} handler.Response "{"msg":"发布成功"}"
// @Failure 400 {object} handler.Error  "{"msg":"新建帖子失败"}"
// @Router /post [POST]
func CreatePost(c *gin.Context) {
	//post := model.Post{
	//	AuthorId:  service.GetId(c),
	//	Title:     c.PostForm("title"),
	//	Type:      c.PostForm("type"),
	//	Content:   c.PostForm("content"),
	//	Likes:     0,
	//	CommentNo: 0,
	//}
	var post model.Post
	c.ShouldBind(&post)
	post.AuthorId = service.GetId(c)
	if c.Query("file_have") == "yes" {
		urls, _ := qiniu.UploadFile(c)
		for len(urls) < 6 {
			urls = append(urls, "")
		}
		post.FilePath1 = urls[0]
		post.FilePath2 = urls[1]
		post.FilePath3 = urls[2]
		post.FilePath4 = urls[3]
		post.FilePath5 = urls[4]
		post.FilePath6 = urls[5]
	}
	err := model.CreatePost(post)
	if err != nil {
		handler.SendError(c, 400, "新建帖子失败.")
		return
	}
	handler.SendResponse(c, "发布成功", gin.H{"data": "null"})
}

// @Summary 回复(json)
// @Description 给指定评论添加回复(comment_id,post_id,object,content,private)
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  comment_id  query integer true "评论的id"
// @Param  object  formData integer true "回复对象的id"
// @Param  content  formData string true "回复的内容"
// @Success 200 {object} handler.Response "{"msg":"回复成功"}"
// @Failure 400 {object} handler.Error  "{"msg":"回复失败"}"
// @Router /post/comment_reply [POST]
func AddReply(c *gin.Context) {
	m := service.RawToMap(c)
	commentId, _ := strconv.Atoi(m["comment_id"])
	postId, _ := strconv.Atoi(m["post_id"])
	object, _ := strconv.Atoi(m["object"])
	var b bool
	a := 1
	d := 1
	if m["private"] == "true" {
		b = true
		a = 0
	} else {
		b = false
		if m["object"] == m["author_id"] {
			d = 0
		}
	}
	reply := model.Reply{
		FromWho:        service.GetId(c),
		CommentId:      uint(commentId),
		Object:         uint(object),
		Content:        m["content"],
		PostId:         uint(postId),
		Private:        b,
		StatusToObject: d,
		StatusToAuthor: a,
	}
	err := model.CreateReply(reply)
	if err != nil {
		handler.SendError(c, 400, "回复失败.")
		return
	}
	handler.SendResponse(c, "回复成功.", nil)
}

// @Summary 评论（json)
// @Description 评论指定的帖子(post_id,private,content)
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  post_id  query integer true "帖子id"
// @Param  content  formData string true "评论的内容"
// @Success 200 {object} handler.Response "{"msg":"评论成功."}"
// @Failure 400 {object} handler.Error  "{"msg":"评论失败."}"
// @Router /post/comment [POST]
func AddComment(c *gin.Context) {
	m := service.RawToMap(c)
	postId, _ := strconv.Atoi(m["post_id"])
	private := m["private"]
	var b bool
	if private == "true" {
		b = true
	} else {
		b = false
	}
	comment := model.Comment{
		PostId:  uint(postId),
		UserId:  service.GetId(c),
		Content: m["content"],
		Status:  1,
		Private: b,
	}
	err := model.CreateComment(comment)
	if err != nil {
		handler.SendError(c, 400, "评论失败.")
		return
	}
	handler.SendResponse(c, "评论成功.", nil)
}

// @Summary 点赞
// @Description 点赞指定的帖子
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  post_id  query integer true "帖子id"
// @Success 200 {object} handler.Response "{"msg":"点赞成功."}"
// @Failure 400 {object} handler.Error  "{"msg":"点赞失败."}"
// @Router /post/likes [POST]
func AddLikes(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Query("post_id"))
	likes := model.Likes{
		PostId: uint(postId),
		UserId: service.GetId(c),
		Status: 1,
	}
	err := model.CreateLikes(likes)
	if err != nil {
		handler.SendError(c, 400, "点赞失败.")
		return
	}
	handler.SendResponse(c, "点赞成功", nil)
}

// @Summary  是否已赞
// @Description 通过帖子id查询是否已经点赞
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param post_id  query integer true "帖子id"
// @Success 200 {object} handler.Response "{"msg":"no"}"
// @Success 200 {object} handler.Response "{"msg":"yes"}"
// @Router /post/whether_like [GET]
func WhetherLike(c *gin.Context) {
	id, _ := c.Get("UserId")
	like := model.WhetherLike(id, c.Query("post_id"))
	if !like {
		handler.SendResponse(c, "no", nil)
		return
	}
	handler.SendResponse(c, "yes", nil)
}

// @Summary 取消点赞
// @Description  通过帖子id取消点赞
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  post_id  query integer true "帖子id"
// @Success 200 {object} handler.Response "{"msg":"取消点赞成功"}"
// @Failure 400 {object} handler.Error  "{"msg":"取消点赞失败"}"
// @Router /post/likes [DELETE]
func DeleteLikes(c *gin.Context) {
	err := model.DeleteLikes(service.GetId(c), c.Query("post_id"))
	if err != nil {
		handler.SendError(c, 400, "取消点赞失败.")
		return
	}
	handler.SendResponse(c, "取消点赞成功", nil)
}

// @Summary  修改帖子
// @Description  上传所有内容然后修改所有内容,
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param post_id  query integer  true "帖子id"
// @Param type  formData string true "类型"
// @Param title  formData string true "标题"
// @Param content  formData string true "文本内容"
// @Param file  formData file false "新文件组(有数量限制)"
// @Param file  query string true "yse/no(说明是否修改了文件)"
// @Success 200 {object} handler.Response "{"msg":"修改帖子成功"}"
// @Failure 400 {object} handler.Error  "{"msg":"修改帖子失败"}"
// @Router /post [PUT]
func UpdatePost(c *gin.Context) {
	if !service.WhetherMyPost(service.GetId(c), c.Query("post_id")) {
		handler.SendError(c, 401, "没有权限.")
		return
	}
	var post model.Post
	c.ShouldBind(&post)
	if c.Query("file") == "yes" {
		urls, _ := qiniu.UploadFile(c)
		for len(urls) < 6 {
			urls = append(urls, "")
		}
		post.FilePath1 = urls[0]
		post.FilePath2 = urls[1]
		post.FilePath3 = urls[2]
		post.FilePath4 = urls[3]
		post.FilePath5 = urls[4]
		post.FilePath6 = urls[5]
	}
	err := model.UpdatePost(c.Query("file"), post, c.Query("post_id"))
	if err != nil {
		handler.SendError(c, 400, "修改帖子失败.")
		return
	}
	handler.SendResponse(c, "修改帖子成功", nil)
}

// @Summary 删帖
// @Description  根据post_id删除指定帖子
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param post_id query integer true "帖子id"
// @Success 200 {object} handler.Response "{"msg":"删帖成功."}"
// @Failure 400 {object} handler.Error  "{"msg":"删帖失败."}"
// @Router /post [DELETE]
func DeletePost(c *gin.Context) {
	if !service.WhetherMyPost(service.GetId(c), c.Query("post_id")) {
		handler.SendError(c, 401, "没有权限.")
		return
	}
	err := model.DeletePost(c.Query("post_id"))
	if err != nil {
		handler.SendError(c, 400, "删帖失败.")
		return
	}
	handler.SendResponse(c, "删帖成功.", nil)
}

// @Summary  搜索帖子
// @Description  通过关键词搜索标题有关键词的帖子
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param  query_string  query string true "关键词"
// @Success 200 {object} handler.Response "{"msg":"搜索成功."}"
// @Failure 410 {object} handler.Error  "{"msg":"搜索失败."}"
// @Router /post/search [GET]
func SearchPosts(c *gin.Context) {
	point := c.Query("query_string")
	posts, err := model.SearchPosts(point)
	if err != nil {
		handler.SendError(c, 410, "搜索失败.")
		return
	}
	handler.SendResponse(c, "搜索成功.", posts)
}

// @Summary 删评论
// @Description  根据comment_id删除指定评论
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param comment_id query integer true "评论id"
// @Param post_id query integer true "帖子id"
// @Success 200 {object} handler.Response "{"msg":"删评论成功."}"
// @Failure 400 {object} handler.Error  "{"msg":"删评论失败."}"
// @Router /post/comment [DELETE]
func DeleteComment(c *gin.Context) {
	if !service.WhetherMyComment(service.GetId(c), c.Query("comment_id")) {
		handler.SendError(c, 401, "没有权限.")
		return
	}
	err := model.DeleteComment(c.Query("comment_id"), c.Query("post_id"))
	if err != nil {
		handler.SendError(c, 400, "删评论失败.")
		return
	}
	handler.SendResponse(c, "删评论成功.", nil)
}

// @Summary 删回复
// @Description  根据reply_id删除指定帖子
// @Tags post
// @Accept  multipart/form-data
// @Produce application/json
// @Param Authorization header string true "token"
// @Param reply_id query integer true "回复id"
// @Success 200 {object} handler.Response "{"msg":"删回复成功."}"
// @Failure 400 {object} handler.Error  "{"msg":"删回复失败."}"
// @Router /post/comment_reply [DELETE]
func DeleteReply(c *gin.Context) {
	if !service.WhetherMyReply(service.GetId(c), c.Query("reply_id")) {
		handler.SendError(c, 401, "没有权限.")
		return
	}
	err := model.DeleteReply(c.Query("reply_id"))
	if err != nil {
		handler.SendError(c, 400, "删回复失败.")
		return
	}
	handler.SendResponse(c, "删回复成功.", nil)
}
