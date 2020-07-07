package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/api/user"
)

func UserRouter(group *gin.RouterGroup)  {
	userApi:=group.Group("user")
	{
		userApi.GET("list",user.List)
		userApi.POST("edit",user.Edit)
		userApi.POST("status",user.Status)
		userApi.POST("del",user.Delete)
		userApi.GET("detail",user.Detail)
	}



}
