package Controllers

import (
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Middlewares"
	"dousheng-backend/Models/model2response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//  get /douyin/feed :不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
// params:
//			latest_time query 			 否  可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
//			token       query 			 否  用户登录状态下设置

func GetFeed(c *gin.Context) {
	db := Databases.DatabaseSession()
	var latestTime, nextTime int64
	token := c.Query("token")
	latestTime_ := c.Query("latest_time")
	var userID uint
	if latestTime_ != "" {
		latestTime, _ = strconv.ParseInt(latestTime_, 10, 64)
	} else {
		latestTime = 0
	}

	if token != "" {
		var err error
		userID, err = Middlewares.ParseToken(c, token)
		if err != nil {
			return
		}

	} else {
		userID = 0
	}

	videos := DAO.GetVideoList(db, latestTime, userID)
	responseVideos := model2response.VideoModeltoRes(db, videos)
	for i := 0; i < len(responseVideos); i++ {
		responseVideos[i].IsFavorite = DAO.IsUserFavoriteVideo(db, userID, uint(responseVideos[i].ID))
	}
	if len(videos)-1 < 0 {
		nextTime = 0
	} else {
		nextTime = videos[len(videos)-1].CreatedAt.Unix()
	}
	c.JSON(http.StatusOK, response.FeedResponse{
		Response: response.Response{
			StatusCode: response.SUCCESS,
			StatusMsg:  "获取成功",
		},
		VideoList: responseVideos,
		NextTime:  nextTime,
	})
}
