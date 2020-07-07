package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_tool/src/db"
)

type Role struct {
	Id int `json:"id" gorm:"column:id;primary_key" form:"id"`
	Name string `json:"name" gorm:"column:name;" form:"name" sql:"index;unique" binding:"required"`
	Desc string `json:"desc" gorm:"column:desc" form:"desc" sql:"index"`
	Status int8 `json:"status" gorm:"column:status" form:"status" sql:"index"`
	GroupID int `json:"group_id" gorm:"column:group_id" from:"group_id"`
	Permission []Permission `json:"permission" gorm:"many2many:role_permission"`
	db.TimeStampModel
}
type RoleWithPerminssion struct {
	Role Role `json:"role"`
	Permission []Permission `json:"permission" gorm:"many2many:role_permission"`
}
func (Role) TableName() string  {
	return "role"
}
//添加角色
func (Role) Add(role Role) (id int,err error) {
	 err=db.MysqlConn.Create(&role).Error
	 id = role.Id
	 return
}
//修改角色
func (Role) Update(role Role ) (id int ,err error)  {
	err=db.MysqlConn.Model(&Role{}).Update(&role).Error
	id = role.Id
	return
}
//状态修改
func ( Role) UpdateStatus(request role.StatusRequest) error {
	return db.MysqlConn.Model(&Role{}).Update(map[string]interface{}{"status":request.Status}).Error
}
//删除角色
func (Role) Del (request role.IDRequest) error {
	return db.MysqlConn.Where("id = ? ",request.ID).Delete(&Role{}).Error
}
//角色列表
func (Role) GetList() *gorm.DB {
	return db.MysqlConn
}
func (r *Role) GetOne (request role.IDRequest)  ( role *Role,err error) {
	r.Id=request.ID

	dbConn:=db.MysqlConn

	if request.IsPermission ==1 {
		dbConn = dbConn.Preload("Permission")
	}
	err =dbConn.First(&r).Error
	role = r
	return
}
func (Role) RoleAndPermission()  {
	var role Role
	db.MysqlConn.Where("id = ?",1).Preload("Permission").First(&role)
}
func (Role) RoleAddPermission ( )  {
	
}
func (Role) RoleDelPermission ( id int )  {

	db.MysqlConn.Where("id = ?",id).Delete(&RolePermission{})

}

