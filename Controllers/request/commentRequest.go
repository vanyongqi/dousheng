package request

type CommentRequest struct {
	VideoID     uint   `form:"video_id" binding:"required"`
	ActionType  uint   `form:"action_type" binding:"required,gt=0,lt=3"`
	CommentText string `form:"comment_text" binding:"omitempty"`
	CommentID   uint   `form:"comment_id" binding:"omitempty"`
}
