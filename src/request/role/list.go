package role

import (
	"github.com/zngue/go_tool/src/common/request"
)

type ListRequest struct {
	ID int `form:"id"`
	Name string `form:"name"`
	Status int `form:"status"`
	request.Page
}
