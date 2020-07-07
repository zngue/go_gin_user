package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/common/response"
)

func Detail(ctx *gin.Context)  {
	var detailRequst user.DetailRequest
	err :=ctx.ShouldBind(&detailRequst)
	if err!=nil || (detailRequst.Name =="" && detailRequst.ID==0 )   {
		response.HttpFailWithParameter(ctx,err.Error())
	}
	one ,err :=Service().Detail(detailRequst)
	if err!=nil {
		response.HttpFail(ctx,err.Error())
		return
	}
	response.HttpOk(ctx,one)
	return
}
