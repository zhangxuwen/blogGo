package main

import (
	"blogWeb/dao/mysql"
	"blogWeb/dao/redis"
	"blogWeb/logger"
	"blogWeb/routers"
	"blogWeb/settings"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Go Web 开发较通用的脚手架模板

func main() {

	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("settings init error: %v\n", err)
		return
	}

	// 初始化日志
	if err := logger.Init(settings.Conf); err != nil {
		fmt.Printf("logger init error: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success ...")

	// 初始化MySQL连接
	if err := mysql.Init(*settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("mysql init error: %v\n", err)
		return
	}
	defer mysql.Close()

	// 初始化Redis连接
	if err := redis.Init(*settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init error: %v\n", err)
		return
	}
	defer redis.Close()

	// 注册路由
	r := routers.Setup()

	// 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("server shutdown: ", zap.Error(err))
	}

	zap.L().Info("server exiting")
}
