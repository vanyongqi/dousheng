package DAO

import (
	"dousheng-backend/Models"
	"gorm.io/gorm"
)

func CreateComment(db *gorm.DB, comment *Models.Comment) *Models.Comment {
	db.Create(&comment).Commit()
	return comment
}

// GetComments desc实现降序
func GetComments(db *gorm.DB, videoID uint) []Models.Comment {
	var comments []Models.Comment
	db.Where("video_id = ?", videoID).Order("created_at desc").Find(&comments)
	return comments
}

func DeleteComment(db *gorm.DB, commentID uint) error {
	var comment *Models.Comment
	return db.Model(&comment).Delete("id = ?", commentID).Error
}

func GetVideoCommentsCountByID(db *gorm.DB, videoID uint) int64 {
	var count int64
	db.Raw("select count(*) from comments where video_id = ? and deleted_at is null", videoID).Scan(&count)
	return count
}
