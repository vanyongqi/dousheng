package model2response

import (
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Models"
	"gorm.io/gorm"
)

func CommentModeltoRes(db *gorm.DB, commentList []Models.Comment) []response.Comment {
	var comments []response.Comment
	for _, comment := range commentList {
		comments = append(comments, CommentModeltoRe(db, comment))
	}
	return comments
}

// CommentModeltoRe
func CommentModeltoRe(db *gorm.DB, comment Models.Comment) response.Comment {
	user, _ := DAO.GetUserByID(db, comment.UserID)
	return response.Comment{
		ID:         int64(comment.ID),
		Content:    comment.Content,
		CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		User:       UserModeltoRe(db, *user),
	}
}
