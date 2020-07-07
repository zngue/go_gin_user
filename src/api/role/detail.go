package role

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_tool/src/common/response"
)

func  Detail (ctx *gin.Context)  {

	var IDrequest role.IDRequest
	err :=ctx.ShouldBind(&IDrequest)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	role ,errs:=Role().Detail(IDrequest)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,role)
	return
}
