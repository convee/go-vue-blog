<p align="center">
	<strong>Go-Vue-Blog 是一款简洁、实用的基于 Golang Vue Markdown 前后端分离博客系统</strong>
</p>
<p align="center">
   <a target="_blank" href="#">
      <img style="display: inline-block;" src="https://img.shields.io/badge/Go-1.17.13-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/Gin-v1.9.0-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/Mysql-5.7-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/GORM-v1.24.3-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/Redis-6.0-red"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/vue-v3.X-green"/>
    </a>
</p>

[在线预览](#在线预览) | [项目介绍](#项目介绍) | [技术介绍](#技术介绍) | [目录结构](#目录结构) | [环境说明](#环境说明) | [快速开始](#快速开始) | [总结&鸣谢](#总结鸣谢)  | [后续计划](#后续计划) | [更新日志](#更新日志)
## 在线预览

* 博客前台链接：[convee.cn](https://www.convee.cn)

* 博客后台链接：[convee.cn/admin](https://www.convee.cn/admin)

## 项目介绍

### 前台：

- 前台界面使用NaiveUI
- 响应式布局，适配了移动端
- 关于我自定义专题页
- 标签页
- 文章页/专题页支持代码高亮

### 后台：

- 鉴权使用 JWT
- 基于 RBAC 的权限管理
- 文章管理
- Markdown 编辑器
- 标签管理
- 自定义专题页管理

## 技术介绍
### 前端技术栈
* 语言：[TypeScript](https://www.typescriptlang.org/zh/)
* 框架：[Vue3](https://cn.vuejs.org/guide/introduction.html)
* 项目构建工具：[Vite](https://cn.vitejs.dev/)
* 软件包管理工具：[Pnpm](https://www.pnpm.cn/)
* Vue 状态管理库：[Pinia](https://pinia.vuejs.org/zh/introduction.html)
* Vue 官方路由：[Vue Router](https://router.vuejs.org/zh/)
* 网络请求库：[Axios](https://www.axios-http.cn/docs/intro)
* Vue 3 组件库：[Naive UI](https://www.naiveui.com/zh-CN/os-theme)
* Markdown 编辑器：[md-editor-v3](https://github.com/imzbf/md-editor-v3)
### 后端技术栈
- 语言：[Golang](https://go.dev/)
- 框架路由使用 [Gin](https://github.com/gin-gonic/gin) 路由
- 中间件使用 [Gin](https://github.com/gin-gonic/gin) 框架的中间件
- 数据库组件 [GORM](https://github.com/jinzhu/gorm)
- 配置文件解析库 [Viper](https://github.com/spf13/viper)
- 校验器使用 [validator](https://github.com/go-playground/validator.v10)  也是 Gin 框架默认的校验器
- 任务调度 [Cron](https://github.com/robfig/cron)
- 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- 高性能日志库 [Zap](https://github.com/uber-go/zap)
- JWT 登录认证 [JWT](https://github.com/golang-jwt/jwt)
- 使用 make 来管理 Go 工程
- 使用 Supervisor 管理进程
- 使用 YAML 文件进行多环境配置
- 使用 Ide 自带 [REST Client](https://www.jetbrains.com/help/idea/http-client-in-product-code-editor.html#converting-curl-requests) 工具测试 API
- [Swagger](https://github.com/swaggo/swag) 自动生成 API 文档
- 使用 Prometheus 监控 QPS、分位耗时等
- 使用钉钉告警panic异常
## 目录结构
```bash
.
├── LICENSE
├── README.md
├── go-blog-api
│   ├── Makefile
│   ├── README.md
│   ├── cmd
│   ├── configs
│   ├── docs
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   ├── logs
│   ├── main.go
│   ├── pkg
│   ├── resource
│   ├── scripts
│   ├── startup.sh
│   └── tests
├── vue-blog-admin
│   ├── LICENSE
│   ├── README.md
│   ├── README.zh-CN.md
│   ├── build
│   ├── commitlint.config.js
│   ├── index.html
│   ├── mock
│   ├── node_modules
│   ├── package.json
│   ├── pnpm-lock.yaml
│   ├── public
│   ├── settings
│   ├── src
│   ├── tsconfig.json
│   ├── types
│   ├── uno.config.ts
│   └── vite.config.ts
└── vue-blog-front
    ├── README.md
    ├── build
    ├── index.html
    ├── node_modules
    ├── package-lock.json
    ├── package.json
    ├── public
    ├── src
    └── vite.config.js
```

## 环境说明

golang 1.17.13
mysql 5.7
redis 6.0

## 快速开始
克隆项目
```bash
git clone https://github.com/szluyu99/gin-vue-blog.git
```
## 接口文档
示例：convee.cn/
### 后端部署
```bash

# 1、进入后端项目根目录 
cd go-blog-api

# 2、修改项目运行的配置文件 
vim configs/config.yml 

# 3、MySQL 导入 vblog.sql
mysql > source vblog.sql

# 4、启动 Redis 
redis-server 6379.conf

# 5、运行项目
go mod tidy
go run main.go

```

### 前台前端部署
```bash

# 1、进入前台前端项目
cd vue-blog-front

# 2、安装依赖
pnpm install

# 3、运行项目
pnpm dev
```

### 后台前端部署
```bash

# 1、进入前台前端项目
cd vue-blog-admin

# 2、安装依赖
pnpm install

# 3、运行项目
pnpm dev
```

## 接口文档
```bash
## 部署
go install github.com/swaggo/swag/cmd/swag@latest
## 在包含main.go文件的项目根目录运行swag init。这将会解析注释并生成需要的文件（docs文件夹和docs/docs.go）。
swag init
## 运行项目
go run main.go
## 访问地址
http://localhost:8000/swagger/index.html
```

## 总结鸣谢
鸣谢项目：

https://github.com/zclzone/vue-naive-admin

博客后台的前端基于 vue-naive-admin 二次开发，感谢作者的开源。

## 后续计划

* 图片上传

## 更新日志