package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_user/src/model"
)

func AutoMigrate(db *gorm.DB)  {
	db.AutoMigrate(
		new(model.User),
		new(model.Role),
		new(model.Permission),
	)

}
