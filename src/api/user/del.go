package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/common/response"
)

func Delete(ctx *gin.Context)  {
	var delRequest user.DeleteRequest
	err :=ctx.ShouldBind(&delRequest)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	err =Service().Delete(delRequest)
	if err != nil {
		response.HttpFail(ctx,err.Error())
		return
	}
	response.HttpOk(ctx)

}
