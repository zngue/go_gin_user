package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_user/src/request/permission"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
	"reflect"
)

type Permission struct {

	ID int `json:"id" gorm:"column:id;primary_key" form:"id" ` //id
	Status int8 `json:"status" gorm:"column:status;tinyint(1);comment:'1 正常，2禁用'" form:"status"  ` //状态
	Name string `json:"name" gorm:"column:name;varchar(50);comment:'权限名称'" form:"name" binding:"required"`//名称
	Url string `json:"url" gorm:"column:url;varchar(100);comment:'跳转链接'" form:"url"` //跳转地址
	Uri string  `json:"uri" gorm:"column:uri;varchar(1024);comment:'权限地址，多个用|分割'" form:"uri"`
	IsMenu bool `json:"is_menu" gorm:"column:is_menu;tinyint(1);comment:'true 菜单 false 不是菜单'" form:"is_menu"`
	Desc string `json:"desc" gorm:"column:desc;varchar(50);comment:'权限描述'" form:"desc"`
	Pid int `json:"pid" gorm:"column:pid;int(11);comment:'权限pid'"  form:"pid"`
	db.TimeStampModel
}

func (Permission) TableName() string  {
	return "permission"
}
func (Permission)StuckToMap(user Permission ) (userInfo map[string]interface{}) {
	userInfo = make(map[string]interface{})
	elem := reflect.ValueOf(&user).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		userInfo[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return
}
//添加
func (Permission) Add (permission Permission) (id int, err error)  {
	err = db.MysqlConn.Create(&permission).Error
	id = permission.ID
	return
}
//更新
func (p *Permission) Update(permission Permission) (id int, err error ) {
	permissionInfo:=common.StuckToMap(permission)
	err =db.MysqlConn.Model(&p).Update(permissionInfo).Error
	id  = permission.ID
	return
}
//删除
func (Permission) Del ( request permission.DelRequest ) (err error) {
	err =db.MysqlConn.Where("id = ?",request.ID).Delete(&Permission{}).Error
	return
}
//修改状态
func ( p *Permission)UpdateStatus(request permission.StatusRequest) (err error) {
	err = db.MysqlConn.Model(&p).Where("id = ?",request.ID).Updates(map[string]interface{}{"status":request.Status}).Error
	return
}
//获取权限列表
func (Permission) GetList()(dbConn *gorm.DB)  {
	dbConn = db.MysqlConn
	return
}
//获取权限详情
func (p *Permission) Detail(request permission.DetailRequest) (permission *Permission,err error)  {
	err=db.MysqlConn.Where("id = ? ",request.ID).First(&p).Error
	permission = p
	return
}


