package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_user/router"
	"github.com/zngue/go_gin_user/src/model"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/gin_run"
	"github.com/zngue/go_tool/src/sign_chan"
	"testing"
)
func TestUser(t *testing.T)  {
	http,_:=gin_run.HttpRouterServe(db.Config.System.Port, func(group *gin.RouterGroup) {
		router.Router(group)
	})
	go gin_run.HttpRun(func() error {
		return http.ListenAndServe()
	})
	go db.InitDB(func(db *gorm.DB) {
		db.AutoMigrate(
			new(model.User),
			new(model.Role),
			new(model.Permission),
		)
	})
	defer db.ConnClose()
	sign_chan.ListClose(func(ctx context.Context) error {
		return http.Shutdown(ctx)
	})
}
