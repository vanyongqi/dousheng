package response

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type Video struct {
	Author        User   `json:"author"`
	CommentCount  int64  `json:"comment_count,"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,"`
	ID            int64  `json:"id,omitempty"`
	IsFavorite    bool   `json:"is_favorite"`
	PlayUrl       string `json:"play_url,omitempty"`
	Title         string `json:"title,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
