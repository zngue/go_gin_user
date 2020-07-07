package user

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/user"
)

func (Service) Detail(request user.DetailRequest) (user *model.User,err error) {
	return new(model.User).Detail(request)
}
