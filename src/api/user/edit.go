package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_tool/src/common/response"
	"github.com/zngue/go_tool/src/fun/md5"
)

func Edit(ctx *gin.Context)  {
	var user2 model.User
	err :=ctx.ShouldBind(&user2)
	if err!=nil {
		response.HttpFailWithCodeAndMessage(422,"参数验证错误",ctx,err.Error())
		return
	}
	if user2.Password!=""{
		user2.Password = md5.MD5(user2.Password)
	}
	ser:=Service()
	r:=ser.Edit(user2)
	if r != nil {
		response.HttpFail(ctx,r.Error())
		return
	}
	response.HttpOk(ctx)
	return


}
