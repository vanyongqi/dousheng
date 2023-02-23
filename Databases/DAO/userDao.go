package DAO

import (
	"dousheng-backend/Models"
	"fmt"
	"gorm.io/gorm"
)

// CreateUser used in ：register
func CreateUser(db *gorm.DB, user *Models.User) *Models.User {
	statement := db.Create(&user)
	if err := statement.Error; err != nil {
		return nil
	}
	statement.Commit()
	return user
}

// GetUserByID used in ： userinfo、 create token、comment
func GetUserByID(db *gorm.DB, userID uint) (*Models.User, error) {
	var user *Models.User
	err := db.First(&user, userID).Error
	if err != nil {
		return nil, fmt.Errorf("not found")
	}
	return user, nil
}

// GetUserByName  used in： login and register to check repeat
func GetUserByName(db *gorm.DB, name string) *Models.User {
	var user *Models.User
	err := db.Where(&Models.User{Name: name}).First(&user).Error
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}
	return user
}

// FavoriteCountNum 喜欢的视频数量
func FavoriteCountNum(db *gorm.DB, userID uint) int64 {
	var num int64
	db.Raw("select count(*) from user_favorite_videos where user_id =?", userID).Scan(&num)
	return num
}

// TotalFavoriteNum 获赞总数量
func TotalFavoriteNum(db *gorm.DB, userID uint) int64 {
	var num int64
	db.Raw("select count(*) FROM user_favorite_videos  where video_id in(SELECT id from videos   where author_id =?)", userID).Scan(&num)
	//SELECT count(*) from videos left join user_favorite_videos on author_id =user_id where author_id =1; 用户发出的赞数量
	return num
}

// WorkCountNum 作品数量
func WorkCountNum(db *gorm.DB, userID uint) int64 {
	var num int64
	db.Raw("SELECT count(*) from videos where author_id=?", userID).Scan(&num)
	return num
}

// IsUserFollow
// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
func IsUserFollow(db *gorm.DB, nowuserID, anotherUserID uint) bool {
	if nowuserID == anotherUserID {
		return false
	}
	var user_id uint
	db.Raw(
		"select user_id from follows where user_id = ? and follow_id = ?", nowuserID, anotherUserID).Scan(&user_id)
	if user_id == 0 {
		return false
	}
	return true
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
func GetUserFollowsCountByID(db *gorm.DB, userID uint) int64 {
	var count int64
	db.Raw("select count(follow_id) from follows where user_id = ?", userID).Scan(&count)
	return count
}

func GetUserFollowersNumByID(db *gorm.DB, userID uint) int64 {
	var count int64
	db.Raw("select count(user_id) from follows where follow_id = ?", &userID).Scan(&count)
	return count
}
