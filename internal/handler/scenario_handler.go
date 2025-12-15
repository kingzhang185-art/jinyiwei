package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/service"
)

// ScenarioHandler 场景管理处理器
type ScenarioHandler struct {
	scenarioService service.ScenarioService
}

// NewScenarioHandler 创建场景管理处理器实例
func NewScenarioHandler(scenarioService service.ScenarioService) *ScenarioHandler {
	return &ScenarioHandler{
		scenarioService: scenarioService,
	}
}

// CreateScenarioRequest 创建场景请求
type CreateScenarioRequest struct {
	Name  string `json:"name" binding:"required,max=100"`
	TagID uint64 `json:"tag_id" binding:"required"`
}

// UpdateScenarioRequest 更新场景请求
type UpdateScenarioRequest struct {
	Name   string `json:"name" binding:"omitempty,max=100"`
	TagID  uint64 `json:"tag_id" binding:"omitempty"`
	Status int    `json:"status" binding:"omitempty,oneof=1 2"`
}

// CreateScenario 创建场景
func (h *ScenarioHandler) CreateScenario(c *gin.Context) {
	var req CreateScenarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	scenario, err := h.scenarioService.CreateScenario(req.Name, req.TagID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建成功",
		"data":    scenario,
	})
}

// GetScenario 获取场景详情
func (h *ScenarioHandler) GetScenario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	scenario, err := h.scenarioService.GetScenarioByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "场景不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": scenario,
	})
}

// GetScenarios 获取场景列表
func (h *ScenarioHandler) GetScenarios(c *gin.Context) {
	status := c.Query("status")

	var scenarios []*model.Scenario
	var err error

	if status == "active" || status == "1" {
		scenarios, err = h.scenarioService.GetActiveScenarios()
	} else {
		scenarios, err = h.scenarioService.GetAllScenarios()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取场景列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": scenarios,
	})
}

// UpdateScenario 更新场景
func (h *ScenarioHandler) UpdateScenario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req UpdateScenarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.scenarioService.UpdateScenario(id, req.Name, req.TagID, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// DeleteScenario 删除场景
func (h *ScenarioHandler) DeleteScenario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	if err := h.scenarioService.DeleteScenario(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// GetScenarioWithGroups 获取场景及其监测组
func (h *ScenarioHandler) GetScenarioWithGroups(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	scenario, err := h.scenarioService.GetScenarioWithGroups(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "场景不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": scenario,
	})
}
