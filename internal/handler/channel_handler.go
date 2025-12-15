package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/service"
)

// ChannelHandler 渠道管理处理器
type ChannelHandler struct {
	channelService service.ChannelService
}

// NewChannelHandler 创建渠道管理处理器实例
func NewChannelHandler(channelService service.ChannelService) *ChannelHandler {
	return &ChannelHandler{
		channelService: channelService,
	}
}

// CreateChannelRequest 创建渠道请求
type CreateChannelRequest struct {
	Name        string `json:"name" binding:"required,max=50"`
	Code        string `json:"code" binding:"required,max=50"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Icon        string `json:"icon" binding:"omitempty,max=255"`
	Sort        int    `json:"sort" binding:"omitempty"`
}

// UpdateChannelRequest 更新渠道请求
type UpdateChannelRequest struct {
	Name        string `json:"name" binding:"omitempty,max=50"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Icon        string `json:"icon" binding:"omitempty,max=255"`
	Sort        int    `json:"sort" binding:"omitempty"`
	Status      int    `json:"status" binding:"omitempty,oneof=1 2"`
}

// CreateChannel 创建渠道
func (h *ChannelHandler) CreateChannel(c *gin.Context) {
	var req CreateChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	channel, err := h.channelService.CreateChannel(req.Name, req.Code, req.Description, req.Icon, req.Sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建成功",
		"data":    channel,
	})
}

// GetChannel 获取渠道详情
func (h *ChannelHandler) GetChannel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	channel, err := h.channelService.GetChannelByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "渠道不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": channel,
	})
}

// GetChannels 获取渠道列表
func (h *ChannelHandler) GetChannels(c *gin.Context) {
	status := c.Query("status")

	var channels []*model.Channel
	var err error

	if status == "active" || status == "1" {
		// 只获取启用的渠道
		channels, err = h.channelService.GetActiveChannels()
	} else {
		// 获取所有渠道
		channels, err = h.channelService.GetAllChannels()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取渠道列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": channels,
	})
}

// UpdateChannel 更新渠道
func (h *ChannelHandler) UpdateChannel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req UpdateChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.channelService.UpdateChannel(id, req.Name, req.Description, req.Icon, req.Sort, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// DeleteChannel 删除渠道
func (h *ChannelHandler) DeleteChannel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	if err := h.channelService.DeleteChannel(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
