package main

import (
	"sentinel-opinion-monitor/internal/config"
	"sentinel-opinion-monitor/internal/pkg/logger"
	"sentinel-opinion-monitor/internal/pkg/mysql"
	"sentinel-opinion-monitor/internal/pkg/redis"
	"sentinel-opinion-monitor/internal/router"
	"sentinel-opinion-monitor/internal/server"

	"go.uber.org/zap"
)

func main() {
	// 1. 初始化配置
	cfg, err := config.Load("")
	if err != nil {
		panic("加载配置失败: " + err.Error())
	}

	// 2. 初始化日志
	if err := logger.Init(cfg.Log.Level); err != nil {
		panic("初始化日志失败: " + err.Error())
	}
	defer logger.Sync()

	logger.Get().Info("应用启动中...")

	// 3. 初始化 MySQL
	if err := mysql.Init(&cfg.MySQL); err != nil {
		logger.Get().Fatal("初始化 MySQL 失败", zap.Error(err))
	}
	defer mysql.Close()

	// 4. 初始化 Redis
	if err := redis.Init(&cfg.Redis); err != nil {
		logger.Get().Fatal("初始化 Redis 失败", zap.Error(err))
	}
	defer redis.Close()

	// 5. 注册路由
	r := router.SetupRouter()

	// 6. 启动 Gin Server
	srv := server.NewServer(cfg, r)

	// 7. 优雅退出（graceful shutdown）
	go func() {
		if err := srv.Start(); err != nil {
			logger.Get().Fatal("服务器启动失败", zap.Error(err))
		}
	}()

	srv.GracefulShutdown()
}

