package Models

import "gorm.io/gorm"

// 可以通过为标签 constraint 配置 OnUpdate、OnDelete 实现外键约束，在使用 GORM 进行迁移时它会被创建
// `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// cascade：在父表上 update / delete记录时，同步 update / delete掉子表的匹配记录
// No action：如果子表中有匹配的记录, 则不允许对父表对应候选键进行update / delete操作

type User struct {
	gorm.Model
	Name           string    `gorm:"uniqueIndex;size:30"`
	Password       string    `gorm:"size:60"`
	Content        string    `gorm:"size:60"`
	Videos         []Video   `gorm:"ForeignKey:AuthorID"`
	Comments       []Comment `gorm:"many2many:comments;joinForeignKey:UserID"`
	FavoriteVideos []Video   `gorm:"many2many:user_favorite_videos"`
	Subscribers    []User    `gorm:"joinForeignKey:UserID;many2many:subscribes"`
	Followers      []User    `gorm:"joinForeignKey:SubscriberID;many2many:subscribes"`
	Firends        []User    `gorm:"joinForeignKey:SubscriberID;many2many:firends"`
	//ChatRecords
}
