package role

type StatusRequest struct {
	ID int `form:"id" binding:"required"`
	Status int `form:"status" binding:"required"`
}
