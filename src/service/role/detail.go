package role

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/role"
)

func (Role) Detail(request role.IDRequest) (*model.Role,error)  {
	return new(model.Role).GetOne(request)
}
