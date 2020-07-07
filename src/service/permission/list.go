package permission

import (
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_gin_user/src/request/permission"
	"github.com/zngue/go_tool/src/fun/zng_str"
)
func (Permission) List(request permission.ListRequest) (num int,permission2 []*model.Permission ,err error) {
	dbConn :=new(model.Permission).GetList()
	if request.Name!="" {
		dbConn = dbConn.Where("name like ? ","%"+request.Name+"%")
	}
	if request.Status!=0 {
		dbConn = dbConn.Where("status = ? ",request.Status)
	}
	if request.IDArrString!="" {
		ids :=zng_str.IDStringToSlice(request.IDArrString,",")
		if len(ids)>0 {
			dbConn = dbConn.Where("id in (?)",ids)
		}
	}
	if request.IsMenu==1 {
		dbConn = dbConn.Where("is_menu = ?",request.IsMenu)
	}
	if request.IsCount==1 {//获取分页数据
		err =dbConn.Count(&num).Error
		if err!=nil {
			return
		}
	}else if request.IsCount==2 {//只返回数量
		err =dbConn.Count(&num).Error
		return
	}
	if request.Page.Page >0 && request.PageSize>0 {
		dbConn = dbConn.Offset(request.PageLimitOffset()).Limit(request.PageSize)
	}
	err=dbConn.Find(&permission2).Error
	return

}
