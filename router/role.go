package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/api/role"
)

func RoleRouter(group *gin.RouterGroup)  {
	roleApi:=group.Group("role")
	{
		roleApi.GET("list",role.List)
		roleApi.GET("detail",role.Detail)
		roleApi.POST("status",role.Status)
		roleApi.POST("edit",role.Edit)
		roleApi.POST("del",role.Del)
		roleApi.POST("editPermission",role.RoleAndPermission)
	}
}
