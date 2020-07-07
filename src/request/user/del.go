package user

type DeleteRequest struct {
	ID int `form:"id" binding:"required"`
	Status int8 `form:"status"`
}
