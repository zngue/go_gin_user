package permission

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/permission"
)
func (Permission) Status(request permission.StatusRequest) (err error)  {
	return new(model.Permission).UpdateStatus(request)
}
