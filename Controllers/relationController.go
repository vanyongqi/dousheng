package Controllers

import (
	"dousheng-backend/Controllers/request"
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Middlewares"
	"dousheng-backend/Models"
	"dousheng-backend/Models/model2response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostRelationAction(c *gin.Context) {
	db := Databases.DatabaseSession()

	var realUser *Models.User
	var req request.RelationRequest
	token := c.Query("token")
	userID, _ := Middlewares.ParseToken(c, token)
	realUser, _ = DAO.GetUserByID(db, userID)

	visitUserID := c.Query("to_user_id")
	visitID, _ := strconv.ParseInt(visitUserID, 10, 64)
	//responseUser := model2response.UserModelChange(db, *visitUser)
	if req.ActionType == 1 {
		DAO.FollowUser(db, realUser, uint(visitID))
	}
	if req.ActionType == 2 {
		DAO.CancelFollowUser(db, realUser, uint(visitID))
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status_msg":  "action is ok",
	})
}

func GetFollowList(c *gin.Context) {
	db := Databases.DatabaseSession()
	userID, _ := Middlewares.ParseToken(c, c.Query("token"))
	Users := DAO.GetUserFollowsByID(db, userID)
	res := model2response.UserModeltoRes(db, Users)
	c.JSON(http.StatusOK, response.UserListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "已找到用户",
		},
		UserList: res,
	})

}
func GetFollowerList(c *gin.Context) {
	db := Databases.DatabaseSession()
	userID, _ := Middlewares.ParseToken(c, c.Query("token"))
	Users := DAO.GetUserFollowersByID(db, userID)
	res := model2response.UserModeltoRes(db, Users)
	c.JSON(http.StatusOK, response.UserListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "已找到用户",
		},
		UserList: res,
	})

}
func GetFriendList(c *gin.Context) {
	db := Databases.DatabaseSession()
	userID, _ := Middlewares.ParseToken(c, c.Query("token"))
	Users := DAO.GetFirendsListByUserID(db, userID)
	res := model2response.UserModeltoRes(db, Users)
	c.JSON(http.StatusOK, response.UserListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "已找到用户",
		},
		UserList: res,
	})

}
