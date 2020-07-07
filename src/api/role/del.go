package role

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_tool/src/common/response"
)

func Del(ctx *gin.Context)  {

	var IdRequest role.IDRequest
	err:=ctx.ShouldBind(&IdRequest)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	errs:=Role().Del(IdRequest)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx)
	return
}
