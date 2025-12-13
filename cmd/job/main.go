package main

import (
	"flag"
	"os"

	"sentinel-opinion-monitor/internal/config"
	"sentinel-opinion-monitor/internal/job"
	"sentinel-opinion-monitor/internal/pkg/logger"
	"sentinel-opinion-monitor/internal/pkg/mysql"
	"sentinel-opinion-monitor/internal/pkg/redis"

	"go.uber.org/zap"
)

func main() {
	// 解析命令行参数
	var task = flag.String("task", "", "要执行的任务名称 (例如: scan)")
	flag.Parse()

	if *task == "" {
		flag.Usage()
		os.Exit(1)
	}

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

	logger.Get().Info("任务脚本启动中...")

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

	// 5. 执行任务
	switch *task {
	case "scan":
		logger.Get().Info("执行舆情扫描任务")
		job.ScanOpinionJob()
	default:
		logger.Get().Error("未知的任务", zap.String("task", *task))
		os.Exit(1)
	}

	logger.Get().Info("任务执行完成")
}

