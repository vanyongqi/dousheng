package request

type FavoriteRequest struct {
	VideoID    uint `form:"video_id" binding:"required"`
	ActionType uint `form:"action_type" binding:"required,min=1,max=2"`
}
