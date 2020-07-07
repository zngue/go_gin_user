package user
//获取用户角色权限
type RolePermissionRequest struct {
	IsRole  bool  `form:"is_role"`
	IsPermissoin bool `form:"is_permission"`
	ID int `form:"id"`
}
