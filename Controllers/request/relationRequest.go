package request

type RelationRequest struct {
	UserID     uint `form:"video_id" binding:"required"`
	ActionType uint `form:"action_type" binding:"required,min=1,max=2"`
}
