package router

import (
	"Mucave/handler"
	"Mucave/handler/Midware"
	"Mucave/handler/post"
	"Mucave/handler/user"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/login", user.Login) //登录 ok
	u := r.Group("/api/v1/user")
	u.Use(Midware.TokenMiddleWare)
	{
		u.GET("/getFile", handler.GetFiles)            //获取文件 ok
		u.POST("/follower", user.Follow)               //关注  ok
		u.DELETE("/follower", user.UnFollow)           //取关  ok
		u.GET("/privateMsg/:id", user.PrivateMsg)      //刷新某人的来信  ok
		u.POST("/privateMsg/:id", user.PrivateMsgSend) //给某人发信息   ok
		u.PUT("/myMsg", user.MyMsgUpdate)              //更新我的资料  ok

		u.GET("/myOutline", user.Outline)     //我的大致信息 ok
		u.GET("/myPost", user.MyPost)         //我的帖子 ok
		u.GET("/following", user.Following)   //我的关注  ok
		u.GET("/followers", user.Followers)   //我的粉丝  ok
		u.GET("/myMsg", user.MyMsg)           //我的详细信息，在编辑页面  ok
		u.GET("/myReplies", user.MyReplies)   //我的回复 ok
		u.GET("/myComments", user.MyComments) //我的评论 ok
		u.GET("/myLikes", user.MyLikesPost)   //我的点赞  ok

		u.GET("/userOutline/:id", user.UserOutline)     //他人的大致信息  ok
		u.GET("/userPost/:id", user.UserPost)           //他人的帖子 ok
		u.GET("/userFollowers/:id", user.UserFollowers) //他人的粉丝  ok
		u.GET("/userFollowing/:id", user.UserFollowing) //他人的关注  ok
		u.GET("/userMsg/:id", user.UserMsg)             //他人的详细信息，只读 ok
	}
	p := r.Group("/api/v1/post")
	p.Use(Midware.TokenMiddleWare)
	{
		p.GET("/latest", post.Latest)                         //查看最新帖子 ok
		p.GET("/following", post.Following)                   //查看关注的人的帖子 ok
		p.GET("/recommendations/:type", post.Recommendations) //查看推荐的帖子 ok
		p.GET("/comments/:postId", post.Comments)             //查看某条帖子的评论 ok
		p.GET("/reply/:commentId", post.Reply)                //查某个评论的回复 ok
		p.GET("/:id", post.QueryOnePosts)                     //查找某一条帖子信息 ok

		p.POST("/postCreate", post.CreatePost)             //发布帖子  ok
		p.POST("/comments/:postId", post.AddComments)      //在帖子下面发评论  ok
		p.POST("/commentsReply/:commentId", post.AddReply) //在评论下回复  ok
		p.POST("/likes/:postId", post.AddLikes)            //点赞某个帖子  ok
	}
}
