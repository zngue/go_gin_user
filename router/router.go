package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/middleware"
)

func Router(api *gin.RouterGroup)  {
	Login(api)
	api.Use(middleware.JWTAuth())
	{
		Permission(api)
		RoleRouter(api)
		UserRouter(api)
	}

}
