package router

import (
	"github.com/gin-gonic/gin"
)

func Router(api *gin.RouterGroup)  {
	Login(api)
	//api.Use(middleware.JWTAuth())
	{
		Permission(api)
		RoleRouter(api)
		UserRouter(api)
	}

}
