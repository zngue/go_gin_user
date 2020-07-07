package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/common/response"
)

func Login(ctx *gin.Context)  {
	var loginRequest user.LoginRequest
	err:=ctx.ShouldBind(&loginRequest)
	if err!=nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	token,errs:=Service().Login(loginRequest)
	if errs!=nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx,token)
	return

}
