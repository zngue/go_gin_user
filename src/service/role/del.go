package role

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/role"
)
func (Role) Del(request role.IDRequest) (err error) {
	//删除角色的时候删除角色对应的权限
	err=new(model.RolePermission).RoleDelPermission(role.PermissionRequest{RoleId: request.ID})
	if err!=nil {
		return
	}
	return new(model.Role).Del(request)
}
