# go-vue-blog
go vue blog

## 项目介绍
## 技术介绍
### 前端技术栈
* 基于 TypeScript
* pnpm 包管理工具
* Vue3
* Pinia
* Vue Router
* Axios
* Naive UI
### 后端技术栈
- Go 语言
- 框架路由使用 [Gin](https://github.com/gin-gonic/gin) 路由
- 中间件使用 [Gin](https://github.com/gin-gonic/gin) 框架的中间件
- 数据库组件 [GORM](https://github.com/jinzhu/gorm)
- 配置文件解析库 [Viper](https://github.com/spf13/viper)
- 校验器使用 [validator](https://github.com/go-playground/validator.v10)  也是 Gin 框架默认的校验器
- 任务调度 [cron](https://github.com/robfig/cron)
- 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- 高性能日志库[Zap](https://github.com/uber-go/zap)
- JWT 登录认证 [JWT](https://github.com/golang-jwt/jwt)
- 使用 make 来管理 Go 工程
- 使用 shell(startup.sh) 脚本来管理进程
- 使用 supervisor 管理进程
- 使用 YAML 文件进行多环境配置
- 使用 Ide 自带 [REST Client](https://www.jetbrains.com/help/idea/http-client-in-product-code-editor.html#converting-curl-requests) 工具测试 API
- 使用 Prometheus 监控 QPS、分位耗时等
- 使用钉钉告警panic异常
## 目录结构