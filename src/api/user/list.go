package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/user"
	user2 "github.com/zngue/go_gin_user/src/service/user"
	"github.com/zngue/go_tool/src/common/response"
)

func List(c *gin.Context)  {
	var userListRequest user.ListRequest
	err :=c.ShouldBind(&userListRequest)
	if err != nil {
		response.HttpFailWithMessage(c,"参数错误",err.Error())
		return
	}
	service :=user2.Service{}
	list,num,errs:=service.List(userListRequest)
	if errs!=nil {
		response.HttpFail(c,errs.Error())
		return
	}
	response.HttpOk(c,gin.H{
		"count":num,
		"item":list,
	})
	return
}




