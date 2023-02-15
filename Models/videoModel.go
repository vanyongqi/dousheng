package Models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorID      uint
	Title         string    `gorm:"size:30"`
	Author        User      `gorm:"reference:ID"`
	UserFavorites []User    `gorm:"many2many:user_favorite_videos"`
	Comments      []Comment `gorm:"many2many:comments;joinForeignKey:VideoID"`
}
