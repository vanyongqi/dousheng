package response

type RelationResponse struct {
	Response
	UserList []User `json:"video_list,omitempty"`
}
