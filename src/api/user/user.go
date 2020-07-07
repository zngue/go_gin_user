package user

import "github.com/zngue/go_gin_user/src/service/user"
func Service() *user.Service  {
	return new(user.Service)
}