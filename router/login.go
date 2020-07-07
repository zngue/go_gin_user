package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/api/user"
)

func Login(group *gin.RouterGroup)  {
	userApi:=group.Group("user")
	userApi.GET("login",user.Login)
}
