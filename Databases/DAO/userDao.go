package DAO

import (
	"dousheng-backend/Models"
	_ "dousheng-backend/Models"
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *Models.User) *Models.User {
	statement := db.Create(&user)
	if err := statement.Error; err != nil {
		return nil
	}
	statement.Commit()
	return user
}

func GetUserByID(db *gorm.DB, userID uint) (*Models.User, error) {
	var user *Models.User
	err := db.First(&user, userID).Error
	if err != nil {
		return nil, fmt.Errorf("未找到用户")
	}
	return user, nil
}

func GetUserByName(db *gorm.DB, name string) *Models.User {
	var user *Models.User
	err := db.Where(&Models.User{Name: name}).First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return user
}

func SubscribeUser(db *gorm.DB, user *Models.User, subscriberUserID uint) (*Models.User, error) {
	var subscriber *Models.User
	err1 := db.First(&subscriber, subscriberUserID).Error
	if err1 != nil {
		return nil, fmt.Errorf("未找到关注人")
	}
	if err := db.Model(&user).Association("Subscribers").Append(subscriber); err != nil {
		return nil, fmt.Errorf("操作失败")
	}
	return user, nil
}

func CancelSubscribeUser(db *gorm.DB, user *Models.User, subscriberUserID uint) (*Models.User, error) {
	var subscriber *Models.User
	err1 := db.First(&subscriber, subscriberUserID).Error
	if err1 != nil {
		return nil, fmt.Errorf("未找到关注人")
	}
	if err := db.Model(&user).Association("Subscribers").Delete(subscriber); err != nil {
		return nil, fmt.Errorf("关注不存在")
	}
	return user, nil
}

func GetUserSubscribersByID(db *gorm.DB, userID uint) []Models.User {
	var user *Models.User
	db.Preload("Subscribers").Find(&user, userID)
	return user.Subscribers
}

func GetUserSubscribersCountByID(db *gorm.DB, userID uint) int64 {
	var count int64
	db.Raw("select count(subscriber_id) from subscribes where user_id = ?",
		userID).Scan(&count)
	return count
}

func GetUserFollowersByID(db *gorm.DB, userID uint) []Models.User {
	var followers []Models.User
	db.Raw("select * from users where id in"+
		"(select user_id from subscribes left join `users`"+
		"on `users`.id = subscriber_id "+
		"where subscriber_id = ?)", userID).Scan(&followers)
	return followers
}

func GetUserFollowersCountByID(db *gorm.DB, userID uint) int64 {
	var count int64
	db.Raw("select count(user_id) from subscribes where subscriber_id = ?",
		&userID).Scan(&count)
	return count
}

func IsUserFollow(db *gorm.DB, userID, anotherUserID uint) bool {
	if userID == anotherUserID {
		return false
	}
	var user_id uint
	db.Raw(
		"select user_id from subscribes where user_id = ? and subscriber_id = ?",
		userID, anotherUserID).Scan(&user_id)
	return user_id != 0
}
