package model2response

import (
	"dousheng-backend/Controllers/response"
	"dousheng-backend/Databases/DAO"
	"dousheng-backend/Models"
	"gorm.io/gorm"
)

func VideoModeltoRes(db *gorm.DB, videoList []Models.Video) []response.Video {
	var videos []response.Video
	for _, video := range videoList {
		videos = append(videos, VideoModeltoRe(db, &video))
	}
	return videos
}

// VideoModeltoRe
func VideoModeltoRe(db *gorm.DB, video *Models.Video) response.Video {

	title := video.Title[11 : len(video.Title)-4]
	return response.Video{
		ID:            int64(video.ID),
		FavoriteCount: DAO.GetFavoriteCount(db, video.ID),
		Author:        UserModeltoRe(db, video.Author),
		CommentCount:  DAO.GetVideoCommentsCountByID(db, video.ID),
		IsFavorite:    true,
		Title:         title,
		//服务器地址 static是资源静态目录
		PlayUrl:  "http://192.168.3.39:8080/static/video/" + video.Title,
		CoverUrl: "http://192.168.3.39:8080/static/cover4video/" + video.Title[:len(video.Title)-4] + ".jpg",
	}
}
