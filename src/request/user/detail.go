package user

type DetailRequest struct {
	ID int `json:"id" form:"id"`
	Name string `form:"name"`
}
