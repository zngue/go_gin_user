package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_tool/src/common/response"
	"github.com/zngue/go_tool/src/jwt"
)
const USER_INFO ="user_info"
	func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtAuth:=jwt.JWTAuth{}
		token := c.Request.Header.Get("token")
		cal,err:=jwtAuth.Parse(token)
		if err!=nil {
			response.HttpFailWithCodeAndMessage(421,"登录信息错误",c,err.Error())
			c.Abort()
			return
		}
		c.Set(USER_INFO,cal.UserInfo.DataInfo)
		c.Next()
	}
}


