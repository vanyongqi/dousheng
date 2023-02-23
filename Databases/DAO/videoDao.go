package DAO

import (
	"dousheng-backend/Models"
	"gorm.io/gorm"
	"time"
)

func CreateVideo(db *gorm.DB, video *Models.Video) *Models.Video {
	statement := db.Create(&video)
	if err := statement.Error; err != nil {
		return nil
	}
	statement.Commit()
	return video
}

func GetVideoList(db *gorm.DB, latestTime int64, userID uint) []Models.Video {
	var videos []Models.Video
	statement := db.Preload("Author").Limit(30)
	if latestTime != 0 {
		statement = statement.Where("created_at < ?",
			time.Unix(latestTime/1000+43200, 0).Format("2006-01-02 15:04:05"))
	}
	if userID != 0 {
		statement = statement.Where("author_id != ?", userID)
	}
	statement.Order("created_at desc").Find(&videos)
	return videos
}

func GetUserPublishVideosByID(db *gorm.DB, userID uint) []Models.Video {
	var videos []Models.Video
	db.Preload("Author").Where("author_id = ?", userID).Find(&videos)
	return videos
}
