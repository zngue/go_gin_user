#运行
```go
package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zngue/go_gin_user/migration"
	"github.com/zngue/go_gin_user/router"
	"github.com/zngue/go_tool/src/db"
	"github.com/zngue/go_tool/src/gin_run"
)
gin_run.GinRun(func(group *gin.RouterGroup) {
    router.Router(group)//路由加载自己加载的路由
}, func(db *gorm.DB) {
    migration.AutoMigrate(db)//数据库迁移
})
defer db.ConnClose()  

```