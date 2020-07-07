package role

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_tool/src/common/response"
)

func List(ctx *gin.Context)  {
	var listRequest  role.ListRequest
	err:=ctx.ShouldBind(&listRequest)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	num,roleList,errs:=Role().List(listRequest)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,gin.H{
		"count":num,
		"item":roleList,
	})
	return




}