package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/service"
)

// OpinionHandler 舆情处理器
type OpinionHandler struct {
	service service.OpinionService
}

// NewOpinionHandler 创建舆情处理器实例
func NewOpinionHandler(service service.OpinionService) *OpinionHandler {
	return &OpinionHandler{
		service: service,
	}
}

// GetOpinion 根据 ID 获取舆情
func (h *OpinionHandler) GetOpinion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的 ID",
		})
		return
	}

	opinion, err := h.service.GetOpinionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "舆情不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": opinion,
	})
}

// GetAllOpinions 获取所有舆情
func (h *OpinionHandler) GetAllOpinions(c *gin.Context) {
	opinions, err := h.service.GetAllOpinions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取舆情列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": opinions,
	})
}

// CreateOpinion 创建舆情
func (h *OpinionHandler) CreateOpinion(c *gin.Context) {
	var opinion model.Opinion
	if err := c.ShouldBindJSON(&opinion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.service.CreateOpinion(&opinion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建舆情失败",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": opinion,
	})
}

