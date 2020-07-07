package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/permission"
	"github.com/zngue/go_tool/src/common/response"
)

func Detail(ctx *gin.Context)  {

	var detail permission.DetailRequest
	err :=ctx.ShouldBind(&detail)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	p,errs:=Perminssion().Detail(detail)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,p)
	return




}
