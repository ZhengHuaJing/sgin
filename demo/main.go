package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenghuajing/demo/global"
	"github.com/zhenghuajing/demo/initialize"
	"github.com/zhenghuajing/demo/service/timing_tasks_service"
	"net/http"
	"fmt"
)

func init() {
	initialize.Config()
	initialize.Log()
	initialize.Mysql()
	initialize.Redis()
	initialize.DBTables()
	initialize.Casbin()
}

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	// 定时任务
	timing_tasks_service.StartTasks()

	serverCfg := global.Config.Server
	gin.SetMode(serverCfg.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverCfg.HttpPort),
		Handler:        initialize.Router(),
		ReadTimeout:    serverCfg.ReadTimeout,
		WriteTimeout:   serverCfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

	// 程序结束前关闭数据库链接
	defer global.DB.Close()
}
