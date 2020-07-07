package permission

import "github.com/zngue/go_tool/src/common/request"

type ListRequest struct {
	IDArrString string `form:"id_arr_string"`//id字符串
	Name string `form:"name"`//名称搜索
	Uri string  `form:"uri"`//权限模糊搜索
	IsMenu int `form:"is_menu"`//是否菜单
	Status int `form:"status"` //权限状态
	request.Page//加载分页模型数据
}
