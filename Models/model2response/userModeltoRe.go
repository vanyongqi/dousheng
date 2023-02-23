package model2response

import (
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Models"
	"gorm.io/gorm"
)

// UserModeltoRes
func UserModeltoRes(db *gorm.DB, userList []Models.User) []response.User {
	var users []response.User
	for _, user := range userList {
		users = append(users, UserModeltoRe(db, user))
	}
	return users
}

// UserModeltoRe
func UserModeltoRe(db *gorm.DB, user Models.User) response.User {
	return response.User{
		ID:              int64(user.ID),
		Name:            user.Name,
		FollowCount:     DAO.GetUserFollowsCountByID(db, user.ID),
		FollowerCount:   DAO.GetUserFollowersNumByID(db, user.ID),
		IsFollow:        true,
		Signature:       user.Signature,
		Avatar:          "http://192.168.3.39:8080/static/avatar4user/" + user.Avatar,
		BackgroundImage: "http://192.168.3.39:8080/static/background4user/" + user.BackgroundImage,
		FavoriteCount:   DAO.FavoriteCountNum(db, user.ID),
		TotalFavorited:  DAO.TotalFavoriteNum(db, user.ID),
		WorkCount:       DAO.WorkCountNum(db, user.ID),
	}
}
