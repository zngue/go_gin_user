package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/permission"
	"github.com/zngue/go_tool/src/common/response"
)

func Status(ctx *gin.Context)  {
	var request permission.StatusRequest
	if err:=ctx.ShouldBind(&request);err!=nil{
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	errs:=Perminssion().Status(request)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx)
	return

}
