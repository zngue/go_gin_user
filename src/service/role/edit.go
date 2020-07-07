package role

import "github.com/zngue/go_gin_user/src/model"

func (Role) Edit(role model.Role) (int,error)  {
	if role.Id>0 {
		return role.Update(role)
	}else {
		return role.Add(role)
	}
}
