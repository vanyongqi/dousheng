package response

// 添加 omitempty 关键字，来表示这条信息如果没有提供，
// 在序列化成 json 的时候就不要包含其默认值。

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type User struct {
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Password        string `json:"password,omitempty"`
	FavoriteCount   int64  `json:"favorite_count"` //喜欢数
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"` //总点赞数
	WorkCount       int64  `json:"work_count"`      //作品数

}

type UserLoginRegisterResponse struct {
	Response
	UserID   int64  `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Token    string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}
