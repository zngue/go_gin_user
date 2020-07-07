package user

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/log"
)
func (Service) List(request user.ListRequest) ( userList []*model.User,num int, err error,)  {
	dbConn:=db.MysqlConn
	if request.Status!=0 {
		dbConn = dbConn.Where("status = ?" ,request.Status)
	}
	if request.Name !=""{
		dbConn = dbConn.Where("name like ?","%"+request.Name+"%")
	}
	if request.Sex!=0 {
		dbConn = dbConn.Where("sex = ?" ,request.Sex)
	}
	if request.IsCount==1 {
		err =dbConn.Model(&model.User{}).Count(&num).Error
		if err != nil {
			return
		}
	}
	if request.Page.Page>0 && request.PageSize>0 {
		dbConn = dbConn.Offset(request.PageLimitOffset()).Limit(request.PageSize)
	}
	err =dbConn.Find(&userList).Error
	if err!=nil {
		log.LogInfo(err)
		return
	}
	return
}
