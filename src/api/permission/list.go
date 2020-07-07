package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/permission"
	"github.com/zngue/go_tool/src/common/response"
)

func List(ctx *gin.Context)  {
	var request permission.ListRequest
	if err :=ctx.ShouldBind(&request);err!=nil{
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	num,list,errs:=Perminssion().List(request)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,gin.H{
		"count":num,
		"item":list,
	})
	return
}
