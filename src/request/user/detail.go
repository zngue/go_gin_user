package user

type DetailRequest struct {
	ID int `json:"id" form:"id"`
	Name string `form:"name"`
	IsRole bool `form:"is_role"`
	IsPermission bool `form:"is_permission"`
}
