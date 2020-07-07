package permission

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/permission"
	"github.com/zngue/go_gin_user/src/request/role"
)
func (Permission) Delete(request permission.DelRequest) (err error) {
	var ids []int
	ids=append(ids, request.ID)
	err=new(model.RolePermission).PermissionDelRole(role.PermissionRequest{PermissionIDArr: ids})
	if err!=nil {
		return
	}
	return new(model.Permission).Del(request)
}
