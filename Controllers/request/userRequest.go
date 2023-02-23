package request

type RegisterLoginRequest struct {
	UserName string `json:"username" binding:"required,min=4,max=32" form:"username"`
	Password string `json:"password" binding:"required,min=4,max=32" form:"password"`
}
