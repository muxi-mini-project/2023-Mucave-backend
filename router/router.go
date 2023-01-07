package router

import (
	"Mucave/handler/Midware"
	"Mucave/handler/post"
	"Mucave/handler/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	r.POST("/api/v1/login", user.Login) //登录 ok
	u := r.Group("/api/v1/user")
	u.Use(Midware.TokenMiddleWare)
	{
		//u.GET("/getFile", handler.GetFiles)                     //获取文件 ok
		u.POST("/following", user.Follow)               //关注  ok
		u.DELETE("/following", user.UnFollow)           //取关  ok
		u.GET("/private_msg/:id", user.PrivateMsg)      //刷新某人的来信  ok
		u.GET("/all_private_msg", user.AllPrivateMsg)   //刷新所有给我的私信
		u.POST("/private_msg/:id", user.PrivateMsgSend) //给某人发信息   ok
		u.PUT("/my_msg", user.MyMsgUpdate)              //更新我的资料  ok

		u.GET("/my_outline", user.Outline)              //我的大致信息 ok
		u.GET("/my_post", user.MyPost)                  //我的帖子 ok
		u.GET("/my_following", user.Following)          //我的关注  ok
		u.GET("/my_followers", user.Followers)          //我的粉丝  ok
		u.GET("/my_msg", user.MyMsg)                    //我的详细信息，在编辑页面  ok
		u.GET("/my_replies", user.MyReplies)            //我的回复 ok
		u.GET("/my_comments", user.MyComments)          //我的评论 ok
		u.GET("/my_likes", user.MyLikesPost)            //我的点赞  ok
		u.GET("/likes_of_my_post", user.LikesOfMyPosts) //刷新点赞我的帖子的消息通知
		u.GET("/replies", user.RepliesToMe)             //刷新回复评论我的消息通知

		u.GET("/whether_follow", user.WhetherFollow)     //他人是否被我关注
		u.GET("/:id/user_outline", user.UserOutline)     //他人的大致信息  ok
		u.GET("/:id/user_post", user.UserPost)           //他人的帖子 ok
		u.GET("/:id/user_followers", user.UserFollowers) //他人的粉丝  ok
		u.GET("/:id/user_following", user.UserFollowing) //他人的关注  ok
		u.GET("/:id/user_msg", user.UserMsg)             //他人的详细信息，只读 ok
	}
	p := r.Group("/api/v1/post")
	p.Use(Midware.TokenMiddleWare)
	{
		p.GET("/search", post.SearchPosts)                //搜索帖子 ok
		p.GET("/latest", post.Latest)                     //查看最新帖子 ok
		p.GET("/following", post.Following)               //查看关注的人的帖子 ok
		p.GET("/recommendations", post.Recommendations)   //查看推荐的帖子 ok
		p.GET("/comments/:post_id", post.Comments)        //查看某条帖子的评论 ok
		p.GET("/comment_replies/:comment_id", post.Reply) //查某个评论的回复 ok
		p.GET("/:id", post.QueryOnePosts)                 //查找某一条帖子信息 ok
		p.GET("/whether_like", post.WhetherLike)          //查询是否已点赞 ok

		p.POST("", post.CreatePost)                  //发布帖子  ok
		p.DELETE("", post.DeletePost)                //删除帖子 ok
		p.PUT("", post.UpdatePost)                   //修改帖子 ok
		p.POST("/comment", post.AddComment)          //在帖子下面发评论  ok
		p.DELETE("/comment", post.DeleteComment)     //删除评论
		p.POST("/comment_reply", post.AddReply)      //在评论下回复  ok
		p.DELETE("/comment_reply", post.DeleteReply) //删除评论的回复
		p.POST("/likes", post.AddLikes)              //点赞某个帖子  ok
		p.DELETE("/likes", post.DeleteLikes)         //取消点赞 ok
	}
}
