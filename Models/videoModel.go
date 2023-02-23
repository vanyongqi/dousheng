package Models

import "gorm.io/gorm"

/*
foreignKey	指定外键
references	指定引用
polymorphic	指定多态类型
polymorphicValue	指定多态值、默认表名
many2many	指定连接表表名
joinForeignKey	指定连接表的外键
joinReferences	指定连接表的引用外键

*/

type Video struct {
	gorm.Model
	AuthorID uint
	Title    string `gorm:"size:30"`

	Author        User      `gorm:"reference:ID"`
	UserFavorites []User    `gorm:"many2many:user_favorite_videos"`
	Comments      []Comment `gorm:"many2many:comments;joinForeignKey:VideoID"`
}
