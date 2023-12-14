package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/convee/go-blog-api/internal/crons"
	"github.com/convee/go-blog-api/internal/daos"
	"github.com/convee/go-blog-api/pkg/jwt"

	"github.com/convee/go-blog-api/configs"
	"github.com/convee/go-blog-api/pkg/logger"
	"github.com/convee/go-blog-api/pkg/redis"
	"github.com/convee/go-blog-api/pkg/shutdown"

	"github.com/convee/go-blog-api/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

var (
	cfgFile = pflag.StringP("config", "c", "./configs/config.yml", "config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *version {
		log.Println("version:", "v1.0")
	}
	// init config
	cfg := configs.Init(*cfgFile)
	// init logger
	logger.Init(&cfg.Logger)
	// init redis
	redis.Init(&cfg.Redis)
	// init mysql
	daos.Init(&cfg.ORM)
	// init jwt
	jwt.Init(&cfg.JWT)

	gin.SetMode(cfg.App.Mode)

	log.Println("http server startup", cfg.App.Addr)
	logger.Info("http server startup")

	srv := &http.Server{
		Addr:    cfg.App.Addr,
		Handler: routers.InitRouter(),
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//定时执行任务
	crons.Init()

	// 优雅关闭
	shutdown.NewHook().Close(

		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Println("http server closed err", err)
			} else {
				log.Println("http server closed")
			}
		},
	)

}
