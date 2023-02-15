package Models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint
	VideoID uint
	Content string `gorm:"index;size:100"`
}
