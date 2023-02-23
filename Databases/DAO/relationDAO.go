package DAO

import (
	"dousheng-backend/Models"
	"fmt"
	"gorm.io/gorm"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
func FollowUser(db *gorm.DB, user *Models.User, followUserID uint) (*Models.User, error) {
	var follow *Models.User
	err1 := db.First(&follow, followUserID).Error
	if err1 != nil {
		return nil, fmt.Errorf("未找到关注人")
	}
	if err := db.Model(&user).Association("Follows").Append(follow); err != nil {
		return nil, fmt.Errorf("操作失败")
	}
	return user, nil
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
func CancelFollowUser(db *gorm.DB, user *Models.User, followUserID uint) (*Models.User, error) {
	var follow *Models.User
	err1 := db.First(&follow, followUserID).Error
	if err1 != nil {
		return nil, fmt.Errorf("未找到关注人")
	}
	if err := db.Model(&user).Association("Follows").Delete(follow); err != nil {
		return nil, fmt.Errorf("关注不存在")
	}
	return user, nil
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
func GetUserFollowsByID(db *gorm.DB, userID uint) []Models.User {
	var user *Models.User
	db.Preload("Follows").Find(&user, userID)
	return user.Follows
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
func GetUserFollowersByID(db *gorm.DB, userID uint) []Models.User {
	var followers []Models.User
	db.Raw("select * from users where id in(select user_id from follows left join `users`on `users`.id = follow_id where follow_id = ?)", userID).Scan(&followers)
	return followers
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
func GetFirendsListByUserID(db *gorm.DB, userID uint) []Models.User {
	var firends []Models.User
	fmt.Println(userID)
	db.Raw("select * from users where id in(  select  a.follow_id from follows as a inner join follows as b on a.user_id = ? and b.follow_id = ? where a.follow_id = b.user_id)", userID, userID).Scan(&firends)
	return firends

}
