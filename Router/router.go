package Router

import (
	"dousheng-backend/Controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.Static("/static", "./public") //资源目录

	apiRouter := r.Group("/api") //路由分组

	//basic apis
	apiRouter.GET("/feed/", Controllers.GetFeed)                    //主页面的视频流
	apiRouter.GET("/user/", Controllers.GetUserInfo)                //获取用户的身份信息
	apiRouter.POST("/user/register/", Controllers.PostRegisterUser) //用户注册
	apiRouter.POST("/user/login/", Controllers.PostLoginUser)       //用户登录
	apiRouter.GET("/publish/action/", Controllers.GetPublishAction) //发布视频
	apiRouter.POST("/publish/list/", Controllers.PostPublishList)   //获取发布视频列表

	// user_behaviour_video apis
	apiRouter.POST("/favorite/action/", Controllers.PostFavorateAction) //赞
	apiRouter.GET("/favorite/list/", Controllers.GetFavoritList)        //赞列表
	apiRouter.POST("/comment/action/", Controllers.PostCommentAction)   //评论
	apiRouter.GET("/comment/list/", Controllers.GetCommentList)         //评论列表

	// user_behaviour_people apis
	apiRouter.POST("relation/action/", Controllers.PostRelationAction)    //关注
	apiRouter.GET("relation/follow/list/", Controllers.GetFollowList)     //关注列表
	apiRouter.GET("relation/follower/list/", Controllers.GetFollowerList) //好友
	apiRouter.GET("relation/friend/list/", Controllers.GetFriendList)     //好友列表
	apiRouter.POST("message/action/", Controllers.PostMessageAction)      //发送消息
	apiRouter.GET("message/chat/", Controllers.GetMessageList)            //聊天记录

}
