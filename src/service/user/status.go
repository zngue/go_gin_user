package user

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/user"
)

func (Service) Status(request user.StatusRequest) (err error) {
	user :=new(model.User)
	return user.UpdateStatus(request)
}
