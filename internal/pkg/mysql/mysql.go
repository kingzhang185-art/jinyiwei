package mysql

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go.uber.org/zap"
	"sentinel-opinion-monitor/internal/config"
	appLogger "sentinel-opinion-monitor/internal/pkg/logger"
)

var db *gorm.DB

// Init 初始化 MySQL 连接
func Init(cfg *config.MySQLConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	// 配置GORM日志级别
	logLevel := gormLogger.Info
	if cfg.LogLevel != "" {
		switch cfg.LogLevel {
		case "silent":
			logLevel = gormLogger.Silent
		case "error":
			logLevel = gormLogger.Error
		case "warn":
			logLevel = gormLogger.Warn
		case "info":
			logLevel = gormLogger.Info
		default:
			logLevel = gormLogger.Info
		}
	}

	// 创建自定义logger，输出SQL到控制台
	// 使用Default logger并设置日志级别，这样可以在终端看到SQL语句
	customLogger := gormLogger.Default.LogMode(logLevel)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		return fmt.Errorf("连接 MySQL 失败: %w", err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	appLogger.Get().Info("MySQL 连接成功",
		zap.String("host", cfg.Host),
		zap.Int("port", cfg.Port),
		zap.String("database", cfg.Database),
	)

	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return db
}

// Close 关闭数据库连接
func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
