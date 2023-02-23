package Models

import "gorm.io/gorm"

// 可以通过为标签 constraint 配置 OnUpdate、OnDelete 实现外键约束，在使用 GORM 进行迁移时它会被创建
// `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// cascade：在父表上 update / delete记录时，同步 update / delete掉子表的匹配记录
// No action：如果子表中有匹配的记录, 则不允许对父表对应候选键进行update / delete操作

type User struct {
	gorm.Model
	Avatar          string `gorm:"size:100"`
	BackgroundImage string `gorm:"size:100"`
	Name            string `gorm:"uniqueIndex;size:20"`
	Password        string `gorm:"size:100"`
	Signature       string `gorm:"size:30"`
	//FavoriteCount   int    `json:"favorite_count"` //喜欢数
	//TotalFavorite   string `json:"total_favorite"` //点赞数
	//WorkCount       int    `json:"work_count"`     //作品数

	Videos         []Video   `gorm:"ForeignKey:AuthorID"`
	Comments       []Comment `gorm:"many2many:comments;joinForeignKey:UserID"`
	FavoriteVideos []Video   `gorm:"many2many:user_favorite_videos"`
	Follows        []User    `gorm:"joinForeignKey:UserID;many2many:follows"`   //我关注的人
	Followers      []User    `gorm:"joinForeignKey:FollowID;many2many:follows"` //关注我的人
	//Firends      []User    `gorm:"joinForeignKey:SubscriberID;many2many:firends"`
	//ChatRecords  []User    `gorm:"joinForeignKey:SubscriberID;many2many:firends"`
}
