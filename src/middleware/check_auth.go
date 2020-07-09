package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_gin_user/src/model"
	role2 "github.com/zngue/go_gin_user/src/request/role"
	"github.com/zngue/go_gin_user/src/service/role"
	"github.com/zngue/go_tool/src/common/response"
	"reflect"
	"strconv"
	"strings"
)

func CheckAuth() gin.HandlerFunc {
	return Chek
}
func Chek(ctx *gin.Context)  {
	userinfo ,_:=ctx.Get(USER_INFO)
	var userModel model.User
	err:=convert(userinfo,&userModel )
	if err!=nil {
		response.HttpFailWithMessage(ctx,"登录信息错误",err.Error())
		ctx.Abort()
		return
	}
	roleId:=userModel.RoleId
	if roleId==0 {
		response.HttpFailWithMessage(ctx,"角色信息不存在")
		ctx.Abort()
		return
	}
	role,roleErr :=new(role.Role).Detail(role2.IDRequest{
		ID: roleId,
		IsPermission: 1,
	})
	if roleErr!=nil {
		response.HttpFailWithMessage(ctx,"角色信息不存在",roleErr.Error())
		ctx.Abort()
		return
	}
	var urlArr []string
	for _,per :=range role.Permission{
		if per.Uri != ""{
			uri :=strings.Split(per.Uri,"|")
			urlArr=append(urlArr, uri...)
		}
	}
	var NewUriArr []string
	for _,val :=range urlArr {
		if  val!=""  && !StringIsInArray(val,NewUriArr){
			NewUriArr=append(NewUriArr,val)
		}
	}
	if !StringIsInArray(ctx.FullPath(),NewUriArr) {
		response.HttpFailWithMessage(ctx,"暂无权限，请联系管理员")
		ctx.Abort()
		return
	}
	ctx.Next()
}

func  StringIsInArray( str string ,strArr  []string ) bool {
	if len(strArr)==0 {
		return false
	}
	for _,val :=range strArr{
		if val!=" " && str!=" " && val== str{
			return true
		}
	}
	return false
}

func convert(Map interface{}, pointer interface{}) error {
	// reflect.Ptr类型 *main.Person
	pointertype := reflect.TypeOf(pointer)
	// reflect.Value类型
	pointervalue := reflect.ValueOf(pointer)
	// struct类型  main.Person
	structType := pointertype.Elem()
	// 将interface{}类型的map转换为  map[string]interface{}
	m ,ok:= Map.(map[string]interface{})
	if !ok {
		return errors.New("map类型断言失败")
	}
	// 遍历结构体字段
	for i := 0; i < structType.NumField(); i++ {
		// 获取指定字段的反射值
		f := pointervalue.Elem().Field(i)
		//获取struct的指定字段
		stf := structType.Field(i)
		// 获取tag
		name := strings.Split(stf.Tag.Get("json"), ",")[0]
		// 判断是否为忽略字段
		if name == "-" {
			continue
		}
		// 判断是否为空，若为空则使用字段本身的名称获取value值
		if name == "" {
			name = stf.Name
		}
		//获取value值
		v, ok := m[name]
		if !ok {
			continue
		}

		//获取指定字段的类型
		kind := pointervalue.Elem().Field(i).Kind()
		// 若字段为指针类型
		if kind == reflect.Ptr {
			// 获取对应字段的kind
			kind = f.Type().Elem().Kind()
		}
		// 设置对应字段的值
		switch kind {
		case reflect.Int:
			res, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)
			pointervalue.Elem().Field(i).SetInt(res)

		case reflect.String:
			pointervalue.Elem().Field(i).SetString(fmt.Sprint(v))
		}
	}
	return nil
}
