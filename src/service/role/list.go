package role

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/role"
)
//获取角色列表
func (Role) List (request role.ListRequest) (num int ,roleList  []*model.Role,err error )  {
	dbConn:=new(model.Role).GetList()

	if request.Status!=0 {
		dbConn = dbConn.Where("status = ?",request.Status)
	}
	if request.Name!="" {
		dbConn = dbConn.Where("name like ?","%"+request.Name+"%")
	}
	if request.IsCount==1 {
		err=dbConn.Count(&num).Error
		if err!=nil {
			return
		}
	}
	if request.IsCount==2 {
		err=dbConn.Count(&num).Error
		if err!=nil {
			return
		}else{
			return
		}
	}
	if request.Page.Page>0 && request.PageSize>0 {
		dbConn=dbConn.Offset(request.PageLimitOffset()).Limit(request.PageSize)
	}
	err=dbConn.Find(&roleList).Error
	if err!=nil {
		return
	}
	return
}