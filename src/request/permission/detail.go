package permission

type DetailRequest struct {
	ID int `form:"id" binding:"required"`
}
