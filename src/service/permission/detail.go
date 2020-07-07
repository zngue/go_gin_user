package permission

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/permission"
)
func (Permission) Detail( request permission.DetailRequest ) (*model.Permission,error)  {
	return new(model.Permission).Detail(request)
}
