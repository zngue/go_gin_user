package role

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_tool/src/common/response"
)

func Status(ctx *gin.Context)  {
	var statueRequest role.StatusRequest
	err:=ctx.ShouldBind(&statueRequest)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	errs:=Role().Status(statueRequest)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx)
	return
}
