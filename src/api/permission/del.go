package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/permission"
	"github.com/zngue/go_tool/src/common/response"
)

func Del(ctx *gin.Context)  {

	var del permission.DelRequest
	errs :=ctx.ShouldBind(&del)
	if errs!=nil {
		response.HttpFailWithParameter(ctx,errs.Error())
		return
	}
	err :=Perminssion().Delete(del)
	if err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}
	response.HttpOk(ctx)
	return



}
