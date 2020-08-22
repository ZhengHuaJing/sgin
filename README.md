## sgin

快速启动Gin项目的脚本

## 技术栈

gin + gorm + casbin + swagger + validation

## 目录结构

```bash
├─demo  	        （项目名）
    │  ├─config         （配置包）
    │  ├─docs  	        （swagger文档目录）
    │  ├─global         （全局对象）
    │  ├─initialiaze    （初始化）
    │  ├─middleware     （中间件）
    │  ├─model          （数据库模型层）
    │  ├─pkg            （公共功能）
    │  ├─router         （路由）
    |  ├─runtime         (运行时产生的文件，如日志、上传等)
    │  ├─service         (业务逻辑层)
```

## 使用说明

1. 加入环境变量

   将 sgin 可执行程序加入到环境变量，确保可以在命令行中使用，不同系统加入方式不同所以不做演示。

2. 启动项目

   进入到 GOPATH 中的项目根路径下。如：

   ``` bash
   cd $GOPATH/src/github.com/zhenghuajing
   ```

   创建项目，并初始化文档。

   ``` bash
   sgin demo      // 命令格式：sgin 项目名  
   cd demo        // 进入新创建项目
   swag init      // 生成 api 文档，也就是那个 docs 目录
   go mod tidy    // 安装所需包
   ```

   配置数据库，打开 config/config.yaml。

   ``` yaml
   mysql:
       Type: 'mysql'
       User: 'root' 		// 你的用户名
       Password: 'root123456' 	// 你的密码
       Host: '127.0.0.1:3306'
       Name: ''  			// 你的数据库名
       TablePrefix:
   ```

   在新创建的项目根目录下执行。

   ``` bash
   go run main.go
   ```

3. 运行结果

   - 命令行

   ![1-1命令行结果图](https://github.com/ZhengHuaJing/sgin/blob/master/images/1-1命令行结果图.jpg)

   - API在线文档

   ![1-2 API在线文档结果图](https://github.com/ZhengHuaJing/sgin/blob/master/images/1-2%20API在线文档结果图.jpg)



