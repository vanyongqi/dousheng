package Controllers

import (
	"dousheng-backend/Controllers/request"
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Models/model2response"
	"dousheng-backend/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostFavoriteAction(c *gin.Context) {
	db := Databases.DatabaseSession()

	var req request.FavoriteRequest

	err := c.ShouldBindQuery(&req)
	if err != nil {
		return
	}

	user := Utils.GetUserFromCtx(c)

	if req.ActionType == 1 {
		if DAO.IsUserFavoriteVideo(db, user.ID, req.VideoID) {

			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: response.BADREQUEST,
				StatusMsg:  "已点赞过视频",
			})
			return
		}
		if err := DAO.ActionFavoriteVideo(db, user, req.VideoID); err != nil {
			c.JSON(http.StatusNotFound, response.Response{
				StatusCode: response.NOTFOUND,
				StatusMsg:  err.Error(),
			})
			return
		}
	} else {

		if !DAO.IsUserFavoriteVideo(db, user.ID, req.VideoID) {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: response.BADREQUEST,
				StatusMsg:  "未点赞过视频",
			})
			return
		}
		if err := DAO.ActionNotFavoriteVideo(db, user, req.VideoID); err != nil {
			c.JSON(http.StatusNotFound, response.Response{
				StatusCode: response.NOTFOUND,
				StatusMsg:  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: response.SUCCESS,
		StatusMsg:  "操作成功",
	})
}

func GetFavoriteList(c *gin.Context) {
	db := Databases.DatabaseSession()

	userID := Utils.QueryIDFromCtx(c, "user_id")

	videos := DAO.GetFavoriteListByUserID(db, userID)

	c.JSON(http.StatusOK, response.VideoListResponse{
		Response: response.Response{
			StatusCode: response.SUCCESS,
			StatusMsg:  "获取成功",
		},
		VideoList: model2response.VideoModeltoRes(db, videos),
	})
}
