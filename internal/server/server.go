package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sentinel-opinion-monitor/internal/config"
	appLogger "sentinel-opinion-monitor/internal/pkg/logger"
)

// Server HTTP 服务器
type Server struct {
	httpServer *http.Server
	config     *config.Config
}

// NewServer 创建服务器实例
func NewServer(cfg *config.Config, router *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
			Handler: router,
		},
		config: cfg,
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	appLogger.Get().Info("服务器启动中",
		zap.Int("port", s.config.Server.Port),
	)

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("服务器启动失败: %w", err)
	}

	return nil
}

// Stop 停止服务器
func (s *Server) Stop(ctx context.Context) error {
	appLogger.Get().Info("服务器关闭中...")
	return s.httpServer.Shutdown(ctx)
}

// GracefulShutdown 优雅关闭
func (s *Server) GracefulShutdown() {
	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Get().Info("收到关闭信号，开始优雅关闭...")

	// 设置 5 秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器
	if err := s.Stop(ctx); err != nil {
		appLogger.Get().Error("服务器关闭失败", zap.Error(err))
	}

	appLogger.Get().Info("服务器已关闭")
}

