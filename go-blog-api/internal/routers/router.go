package routers

import (
	_ "github.com/convee/go-blog-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	"github.com/convee/go-blog-api/internal/handler/api"
	"github.com/convee/go-blog-api/internal/handler/backend"
	"github.com/convee/go-blog-api/internal/routers/middleware"
	"github.com/convee/go-blog-api/pkg/utils"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy. At least healthy enough for demoing purposes.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: utils.GetHostname()})
}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	// pprof router 性能分析路由
	// 默认关闭，开发环境下可以打开
	// 访问方式: HOST/debug/pprof
	// 通过 HOST/debug/pprof/profile 生成profile
	// 查看分析图 go tool pprof -http=:5000 profile (安装graphviz: brew install graphviz)
	// see: https://github.com/gin-contrib/pprof
	pprof.Register(r)
	r.Use(middleware.RequestID())
	r.Use(middleware.Logging())
	r.Use(middleware.Metrics(nil))
	r.Use(gin.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.Cors())
	// HealthCheck 健康检查路由
	r.GET("/health", HealthCheck)
	// metrics router 可以在 prometheus 中进行监控
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	articleHandler := api.NewArticleHandler()
	tagHandler := api.NewTagHandler()
	pageHandler := api.NewPageHandler()
	categoryHandler := api.NewCategoryHandler()

	authHandler := backend.NewAuthHandler()
	backendUserHandler := backend.NewUserHandler()
	backendArticleHandler := backend.NewArticleHandler()
	backendCategoryHandler := backend.NewCategoryHandler()
	backendPageHandler := backend.NewPageHandler()
	backendTagHandler := backend.NewTagHandler()

	// 前台接口
	apiGroup := r.Group("/api")
	// 文章列表
	apiGroup.GET("/article/list", articleHandler.List)
	// 文章详情页
	apiGroup.GET("/article/:id", articleHandler.Detail)
	// 标签页
	apiGroup.GET("/tag/list", tagHandler.List)
	// 分类
	apiGroup.GET("/category/list", categoryHandler.List)
	// 自定义页面
	apiGroup.GET("/page/:ident", pageHandler.Detail)

	backendGroup := r.Group("/backend")
	backendGroup.POST("/auth/login", authHandler.Login)
	backendGroup.POST("/auth/refreshToken", authHandler.RefreshToken)

	backendGroup.Use(middleware.JWT())
	{
		backendGroup.GET("/user", backendUserHandler.Index)
		backendGroup.GET("/article/stat", backendArticleHandler.Stat)
		backendGroup.GET("/article/list", backendArticleHandler.List)
		backendGroup.GET("/article/detail", backendArticleHandler.Detail)
		backendGroup.POST("/article/add", backendArticleHandler.Add)
		backendGroup.POST("/article/update", backendArticleHandler.Update)
		backendGroup.POST("/article/delete", backendArticleHandler.Delete)
		backendGroup.GET("/category/all", backendCategoryHandler.All)
		backendGroup.GET("/category/list", backendCategoryHandler.List)
		backendGroup.GET("/category/detail", backendCategoryHandler.Detail)
		backendGroup.POST("/category/add", backendCategoryHandler.Add)
		backendGroup.POST("/category/update", backendCategoryHandler.Update)
		backendGroup.POST("/category/delete", backendCategoryHandler.Delete)
		backendGroup.GET("/page/list", backendPageHandler.List)
		backendGroup.GET("/page/detail", backendPageHandler.Detail)
		backendGroup.POST("/page/add", backendPageHandler.Add)
		backendGroup.POST("/page/update", backendPageHandler.Update)
		backendGroup.POST("/page/delete", backendPageHandler.Delete)
		backendGroup.GET("/tag/all", backendTagHandler.All)
		backendGroup.GET("/tag/list", backendTagHandler.List)
		backendGroup.POST("/tag/add", backendTagHandler.Add)
		backendGroup.POST("/tag/update", backendTagHandler.Update)
		backendGroup.POST("/tag/delete", backendTagHandler.Delete)
	}
	return r
}
