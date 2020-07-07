package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_tool/src/common/response"
)

func Edit(ctx *gin.Context)  {
	var permiss model.Permission
	err :=ctx.ShouldBind(&permiss)
	if err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	id ,err:=Perminssion().Edit(permiss)
	if err!=nil {
		response.HttpFail(ctx,gin.H{
			"id":id,
			"err":err.Error(),
		})
		return
	}
	response.HttpOk(ctx,id)
	return

}
