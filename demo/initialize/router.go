package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zhenghuajing/demo/docs"
	"github.com/zhenghuajing/demo/global"
	"github.com/zhenghuajing/demo/middleware"
	"github.com/zhenghuajing/demo/model"
	"github.com/zhenghuajing/demo/pkg/upload"
	"github.com/zhenghuajing/demo/pkg/util"
	"github.com/zhenghuajing/demo/router/api"
	v1 "github.com/zhenghuajing/demo/router/api/v1"
	"github.com/zhenghuajing/demo/service/api_service"
	"net/http"
	"strings"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	{
		// API在线文档
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// 上传文件
		r.POST("/upload", api.UploadImage)
		// 图片访问
		r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	}

	// 无需认证权限路由
	apiNotAuthV1 := r.Group("/api/v1")
	{
		// 生成验证码
		apiNotAuthV1.GET("/captcha", v1.GenerateCaptcha)

		// 用户登录
		apiNotAuthV1.POST("/login", v1.Login)
		// 用户注册
		apiNotAuthV1.POST("/users", v1.AddUser)
	}

	// 需要认证权限路由
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.UserAuth())
	apiV1.Use(middleware.Casbin())
	{
		// 修改用户
		apiV1.PUT("/users/:id", v1.UpdateUser)
		// 用户详情
		apiV1.GET("/users/:id", middleware.IsSelfOperate(), v1.GetUser)
		// 删除用户
		apiV1.DELETE("/users/:id", v1.DeleteUser)
		// 用户列表
		apiV1.GET("/users", v1.GetAllUser)

		// API详情
		apiV1.GET("/apis/:id", v1.GetApi)
		// 查询所有API
		apiV1.GET("/apis", v1.GetAllApi)

		// 为用户添加角色
		apiV1.POST("/user/role", v1.AddRoleForUser)
		// 删除角色
		apiV1.DELETE("/roles/:role_name", v1.DeleteRole)
		// 添加/修改角色
		apiV1.POST("/roles", v1.UpdateRole)
		// 查询所有角色
		apiV1.GET("/roles", v1.GetAllRole)
	}

	apiDocsMigrate()

	return r
}

// 将api接口信息自动同步到数据库apis表中
func apiDocsMigrate() {
	apiDocs := util.JsonFileToMap(global.Config.Casbin.ApiJsonFilePath)
	apiModel := model.Api{}

	for k, v := range apiDocs {
		for k2, v2 := range v {
			apiModel.Path = k
			apiModel.Description = v2["description"].(string)
			apiModel.ApiTag = v2["tags"].([]interface{})[0].(string)
			apiModel.Method = strings.ToUpper(k2)

			api_service.AddApi(apiModel)
		}
	}
}
