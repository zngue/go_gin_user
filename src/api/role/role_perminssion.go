package role

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_tool/src/common/response"
	"github.com/zngue/go_tool/src/fun/zng_str"
)

//角色添加权限
func RoleAndPermission(ctx *gin.Context)  {
	var role role.PermissionRequest
	err :=ctx.ShouldBind(&role)
	if err != nil {
		response.HttpFailWithParameter(ctx,err.Error())
		return
	}
	role.PermissionIDArr=zng_str.IDStringToSlice(role.PermissionIdString,",")
	errs :=Role().RolePermission(role)
	if errs != nil {
		response.HttpFail(ctx,errs.Error())
		return
	}
	response.HttpOk(ctx)
	return


}
