package user

import "github.com/zngue/go_tool/src/common/request"
type ListRequest struct {
	request.Page
	Status int 	`form:"status"`
	Sex int `form:"sex"`
	Name string `form:"name"`

}
