package permission

import "github.com/zngue/go_gin_user/src/model"

func (Permission) Edit( permission model.Permission ) (id int, err error) {
	if permission.ID>0 {
		return permission.Update(permission)
	}else{
		return permission.Add(permission)
	}
}
