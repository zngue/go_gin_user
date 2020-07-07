package model

import (
	"errors"
	"github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_tool/src/db"
)

type RolePermission struct {
	PermissionID int `gorm:"column:permission_id"`
	RoleID int `gorm:"column:role_id"`
}
func (RolePermission)TableName() string {
	return  "role_permission"
}
//给角色添加权限
func (RolePermission) RoleAddPermission (request role.PermissionRequest) error  {
	dbConn :=db.MysqlConn.Begin()
	defer func() {
		dbConn.Rollback()
	}()
	var err []string
	for _,permissionId :=range request.PermissionIDArr{
		role :=RolePermission{
			PermissionID: permissionId,
			RoleID: request.RoleId,
		}
		errs :=dbConn.Create(&role).Error
		if errs!=nil {
			err=append(err, errs.Error())
		}
	}
	if len(err)==0 {
		errss:=dbConn.Commit().Error
		if errss!=nil {
			return errss
		}else{
			return nil
		}
	}else{
		errs:=dbConn.Rollback().Error
		if errs!=nil {
			return errs
		}else{
			return errors.New("role add permission fail")
		}
	}
}
//删除角色的对应的权限
func (RolePermission) RoleDelPermission (request role.PermissionRequest) error {
	return db.MysqlConn.Where(" role_id = ? ",request.RoleId).Delete(&RolePermission{}).Error
}
//从角色中删除某个权限
func (RolePermission) PermissionDelRole (request role.PermissionRequest) error  {
	return db.MysqlConn.Where("permission_id in (?) ",request.PermissionIDArr).Error
}



