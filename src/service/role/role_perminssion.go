package role

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/role"
)
func (Role) RolePermission(request role.PermissionRequest) (err error) {
	m :=new(model.RolePermission)
	err=m.RoleDelPermission(request)
	if err!=nil {
		return
	}
	if request.Status==1 && len(request.PermissionIDArr)>0 {
		err=m.RoleAddPermission(request)
		if err!=nil {
			return
		}
	}
	return
}
