package response

type Comment struct {
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date"`
	ID         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}
