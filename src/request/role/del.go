package role

import "github.com/gin-gonic/gin"

type IDRequest struct {
	ID           int `form:"id" binding:"required"`
	IsPermission int `form:"is_permission"`
}
func (IDRequest) Check (ctx *gin.Context,data interface{}) error {
	return ctx.ShouldBind(&data)
}
