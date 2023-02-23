package Router

import (
	"dousheng-backend/Controllers"
	"dousheng-backend/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.Static("/static", "./public") //资源目录

	apiRouter := r.Group("/douyin")                                   //路由分组
	login_apiRouter := apiRouter.Group("/", Middlewares.Token2User()) //登陆后才可操作的路由分组
	//test apis
	//apiRouter.GET("/re", Controllers.Register) //test api from camp code ,test is over at  2023/02/16/12:44

	//basic apis
	apiRouter.GET("/feed/", Controllers.GetFeed)                      //主页面的视频流
	apiRouter.GET("/user/", Controllers.GetUserInfo)                  //用户信息  			TOKEN
	apiRouter.POST("/user/register/", Controllers.PostRegisterUser)   //用户注册
	apiRouter.POST("/user/login/", Controllers.PostLoginUser)         //用户登录
	apiRouter.POST("/publish/action/", Controllers.PostPublishAction) //发布视频			TOKEN
	login_apiRouter.GET("/publish/list/", Controllers.GetPublishList) //获取发布视频列表 	TOKEN	LOGIN

	// user_behaviour_video apis
	login_apiRouter.POST("/favorite/action/", Controllers.PostFavoriteAction) //赞 		TOKEN	LOGIN
	login_apiRouter.GET("/favorite/list/", Controllers.GetFavoriteList)       //赞列表 	TOKEN	LOGIN
	login_apiRouter.POST("/comment/action/", Controllers.PostCommentAction)   //评论 				LOGIN
	apiRouter.GET("/comment/list/", Controllers.GetCommentList)               //评论列表	TOKEN

	// user_behaviour_social apis
	apiRouter.POST("relation/action/", Controllers.PostRelationAction)    //关注 			TOKEN	LOGIN
	apiRouter.GET("relation/follow/list/", Controllers.GetFollowList)     //关注列表 		TOKEN	LOGIN
	apiRouter.GET("relation/follower/list/", Controllers.GetFollowerList) //粉丝列表 		TOKEN	LOGIN
	apiRouter.GET("relation/friend/list/", Controllers.GetFriendList)     //好友列表 		TOKEN	LOGIN（ 粉丝 ∩ 关注 = 好友）
	apiRouter.POST("message/action/", Controllers.PostMessageAction)      //发送消息 		TOKEN	LOGIN
	apiRouter.GET("message/chat/", Controllers.GetMessageList)            //聊天记录 		TOKEN	LOGIN

}
