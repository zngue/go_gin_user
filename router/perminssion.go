package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/api/permission"
)

func Permission(group *gin.RouterGroup)  {
	permssions:=group.Group("permission")
	{
		permssions.GET("list",permission.List)
		permssions.POST("edit",permission.Edit)
		permssions.POST("del",permission.Del)
		permssions.POST("status",permission.Status)
		permssions.GET("detail",permission.Detail)

	}

}
