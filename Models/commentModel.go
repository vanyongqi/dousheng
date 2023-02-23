package Models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"index;size:100"`
	UserID  uint
	VideoID uint
}
