package permission

type StatusRequest struct {
	ID int `form:"id" binding:"required"`
	Status int8 `form:"status" binding:"required"`
}
