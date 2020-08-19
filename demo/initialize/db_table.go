package initialize

import (
	"github.com/zhenghuajing/demo/global"
	"github.com/zhenghuajing/demo/model"
)

// 注册数据库表专用
func DBTables() {
	// 因为同步最新的api接口到数据库，所以要删除旧表
	if global.DB.HasTable("apis") {
		global.DB.DropTable("apis")
	}

	global.DB.AutoMigrate(
		model.User{},
		model.Api{},
	)
}
