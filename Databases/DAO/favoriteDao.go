package DAO

import (
	"dousheng-backend/Models"
	"fmt"
	"gorm.io/gorm"
)

func GetFavoriteCount(db *gorm.DB, videoID uint) int64 {
	var count int64
	db.Raw("select count(user_id) from user_favorite_videos where video_id = ?", videoID).Scan(&count)
	return count
}

func IsUserFavoriteVideo(db *gorm.DB, userID, videoID uint) bool {
	var video_id uint
	db.Raw("select video_id from user_favorite_videos where user_id = ? and video_id = ?", userID, videoID).Scan(&video_id)
	if video_id == 0 {
		return false
	}
	return true
}
func ActionFavoriteVideo(db *gorm.DB, user *Models.User, videoID uint) error {
	var video *Models.Video
	err := db.First(&video, videoID).Error
	if err != nil {
		return fmt.Errorf("not found video")
	}
	db.Model(&user).Association("FavoriteVideos").Append(video)
	db.Commit()
	return nil
}

func ActionNotFavoriteVideo(db *gorm.DB, user *Models.User, videoID uint) error {
	var video *Models.Video
	err := db.First(&video, videoID).Error
	if err != nil {
		return fmt.Errorf("not found")
	}
	if db.Model(&user).Association("FavoriteVideos").Delete(video) != nil {
		return fmt.Errorf("not found")
	}
	db.Commit()
	return nil
}

func GetFavoriteListByUserID(db *gorm.DB, userID uint) []Models.Video {
	var videolist []Models.Video
	db.Raw("select * from videos where id in (select video_id from user_favorite_videos where user_id = ?)", userID).Scan(&videolist)
	return videolist
}
