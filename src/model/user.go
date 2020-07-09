package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_user/src/request/user"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/fun/common"
	"github.com/zngue/go_tool/src/log"
	"reflect"
)
type User struct {
	ID 		int 		`json:"user_id" gorm:"column:id;primary_key;" sql:"index" form:"id"`
	Name 		string 		`json:"name" gorm:"column:name" form:"name" sql:"index;unique" binding:"required" sql:"index"`
	Password 	string 		`json:"password" gorm:"column:password" form:"password" `
	HeaderImg 	string 		`json:"header_img" gorm:"column:header_img" form:"header_img"`
	Age 		int8 		`json:"age" gorm:"column:age" form:"age"`
	Status 		int8 		`json:"status" gorm:"column:status;default:1" form:"status"`
	Sex 		int8 		`json:"sex" gorm:"column:sex" form:"sex"`
	RoleId  	int 		`json:"role_id" gorm:"column:role_id" form:"role_id"`
	Role 		Role		`gorm:"ForeignKey:id;AssociationForeignKey:role_id" json:"role"`

	db.TimeStampModel
}
func (User) TableName() string  {
	return "user"
}
func (User)StuckToMap(user User ) (userInfo map[string]interface{}) {
	userInfo = make(map[string]interface{})
	elem := reflect.ValueOf(&user).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		userInfo[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return
}
func (u  *User)Add (user User) error {
	err:=db.MysqlConn.Create(&user).Error
	if err != nil {
		log.LogInfo(err.Error())
		return err
	}
	return nil
}
func (u *User) Update(user User) error  {
	userInfo:=common.StuckToMap(&user)
	err :=db.MysqlConn.Model(&u).Update(&userInfo).Error
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
func (u *User) Detail(request user.DetailRequest) ( user  *User, err error)  {

	if request.Name!="" {
		u.Name = request.Name
	}
	dbConn:=db.MysqlConn
	if request.ID!=0 {
		dbConn= dbConn.Where("user_id = ?",request.ID)
	}
	if request.Name!="" {
		dbConn = dbConn.Where("name = ?",request.Name)
	}
	if request.IsRole {
		dbConn = dbConn.Preload("Role", func( db2 *gorm.DB) *gorm.DB {
			if  request.IsPermission{
				return db2.Preload("Permission")
			}else {
				return nil
			}
		})
	}
	err =dbConn.First(&u).Error
	user = u
	return
}
func ( u *User) UpdateStatus (  request user.StatusRequest ) (err error)  {
	userinfo := make(map[string]interface{})
	userinfo["status"]=request.Status
	err = db.MysqlConn.Model(&u).Where("user_id = ?",request.ID).Update(userinfo).Error
	if err!=nil {
		log.LogInfo(err)
		return
	}
	return
}
func ( u *User ) DelDeleteTime(request user.DeleteRequest) (err error)  {
	u.ID = request.ID
	err = db.MysqlConn.Delete(&u).Error
	return
}
func (u *User) GetUserRoleAndPerminss(request user.RolePermissionRequest) ( user *User,err error) {

	dbConn :=db.MysqlConn
	if request.IsRole {
		dbConn = dbConn.Preload("Role", func(db2 *gorm.DB) *gorm.DB {
			if request.IsPermissoin {
				return  db2.Preload("Permission")
			}else{
				return nil
			}
		})
	}
	err=dbConn.Where("id = ? ",request.ID).First(&user).Error
	return
}
type UserRole struct {
	User
	Role []Role  `json:"role" gorm:"many2many:user_role"`
	//RolePermission RoleWithPerminssion `json:"role_permission" gorm:"many2many:user_role"`
}
func (u *User)GetUserRolePermission()  {


}


