package role

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_tool/src/common/response"
)

func Edit(ctx *gin.Context)  {
	var role model.Role
	err :=ctx.ShouldBind(&role)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	id,errs:=Role().Edit(role)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,id)
	return

}
