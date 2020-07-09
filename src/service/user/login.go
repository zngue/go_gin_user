package user

import (
	"errors"
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/fun/md5"
	"github.com/zngue/go_tool/src/jwt"
)

func (Service) Login(request user.LoginRequest) (token string,errs error)   {
	getOne :=user.DetailRequest{
		Name: request.Name,
	}
	model :=model.User{}
	info,err:=model.Detail(getOne)
	if err!=nil || info.Password!=md5.MD5(request.Password) {
		errs = errors.New("账号或者密码错误")
		return
	}
	auth:=jwt.JWTAuth{}
	token,errs=auth.CreateToken(info)
	return

}
