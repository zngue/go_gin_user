package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/common/response"
)

func Status(ctx *gin.Context)  {
	var statusRequest user.StatusRequest
	err :=ctx.ShouldBind(&statusRequest)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	serr:=Service().Status(statusRequest)
	if serr!=nil {
		response.HttpFail(ctx,serr.Error())
		return
	}
	response.HttpOk(ctx)
	return
}
