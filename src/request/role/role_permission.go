package role

type PermissionRequest struct {
	RoleId int `json:"role_id" form:"role_id" binding:"required"`
	PermissionIdString string  `json:"permission_id" form:"permission_id"`
	PermissionIDArr []int `json:"permission_id_arr"`
	Status int `json:"status" form:"status"`
}
