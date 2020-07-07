package permission

type DelRequest struct {
	ID int `json:"id" form:"id" binding:"required"`
}
