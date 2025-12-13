package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler Ping 处理器
type PingHandler struct{}

// NewPingHandler 创建 Ping 处理器实例
func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Ping 健康检查接口
func (h *PingHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  "ok",
	})
}

