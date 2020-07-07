package role

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/role"
)

func (Role) Status( request role.StatusRequest ) error  {
	return new(model.Role).UpdateStatus(request)
}
