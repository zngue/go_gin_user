package user

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/log"
)

func (Service) Delete(request user.DeleteRequest) (err error,) {

	userModel :=new(model.User)
	err =userModel.DelDeleteTime(request)
	if err != nil {
		log.LogInfo(err)
	}
	return
}
