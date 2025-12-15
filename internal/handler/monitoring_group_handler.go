package handler

import (
	"net/http"
	"strconv"

	"sentinel-opinion-monitor/internal/service"

	"github.com/gin-gonic/gin"
)

// MonitoringGroupHandler 监测组管理处理器
type MonitoringGroupHandler struct {
	groupService service.MonitoringGroupService
}

// NewMonitoringGroupHandler 创建监测组管理处理器实例
func NewMonitoringGroupHandler(groupService service.MonitoringGroupService) *MonitoringGroupHandler {
	return &MonitoringGroupHandler{
		groupService: groupService,
	}
}

// CreateGroupRequest 创建监测组请求
type CreateGroupRequest struct {
	ScenarioID     uint64   `json:"scenario_id" binding:"required"`
	Name           string   `json:"name" binding:"required,max=100"`
	Sort           int      `json:"sort" binding:"omitempty"`
	Keywords       []string `json:"keywords" binding:"omitempty"`
	ExclusionWords []string `json:"exclusion_words" binding:"omitempty"`
}

// UpdateGroupRequest 更新监测组请求
type UpdateGroupRequest struct {
	Name   string `json:"name" binding:"omitempty,max=100"`
	Sort   int    `json:"sort" binding:"omitempty"`
	Status int    `json:"status" binding:"omitempty,oneof=1 2"`
}

// AssignChannelsRequest 分配渠道请求
type AssignChannelsRequest struct {
	ChannelIDs []uint64 `json:"channel_ids" binding:"required"`
}

// AddKeywordRequest 添加关键词请求
type AddKeywordRequest struct {
	Keyword string `json:"keyword" binding:"required"`
}

// AddExclusionWordRequest 添加排除词请求
type AddExclusionWordRequest struct {
	Word string `json:"word" binding:"required"`
}

// CreateGroup 创建监测组
func (h *MonitoringGroupHandler) CreateGroup(c *gin.Context) {
	var req CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	// 如果提供了关键词或排除词，使用事务方法创建
	if len(req.Keywords) > 0 || len(req.ExclusionWords) > 0 {
		group, err := h.groupService.CreateGroupWithKeywordsAndExclusionWords(
			req.ScenarioID,
			req.Name,
			req.Sort,
			req.Keywords,
			req.ExclusionWords,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "创建成功",
			"data":    group,
		})
		return
	}

	// 如果没有提供关键词和排除词，使用原来的方法
	group, err := h.groupService.CreateGroup(req.ScenarioID, req.Name, req.Sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建成功",
		"data":    group,
	})
}

// GetGroup 获取监测组详情
func (h *MonitoringGroupHandler) GetGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	group, err := h.groupService.GetGroupWithDetails(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "监测组不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": group,
	})
}

// GetGroupsByScenario 根据场景ID获取监测组列表
func (h *MonitoringGroupHandler) GetGroupsByScenario(c *gin.Context) {
	scenarioIDStr := c.Param("scenario_id")
	scenarioID, err := strconv.ParseUint(scenarioIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的场景ID",
		})
		return
	}

	groups, err := h.groupService.GetGroupsByScenarioID(scenarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取监测组列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

// UpdateGroup 更新监测组
func (h *MonitoringGroupHandler) UpdateGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req UpdateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.groupService.UpdateGroup(id, req.Name, req.Sort, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// DeleteGroup 删除监测组
func (h *MonitoringGroupHandler) DeleteGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	if err := h.groupService.DeleteGroup(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// AssignChannels 分配渠道
func (h *MonitoringGroupHandler) AssignChannels(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req AssignChannelsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.groupService.AssignChannels(id, req.ChannelIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "分配渠道成功",
	})
}

// AddKeyword 添加关键词
func (h *MonitoringGroupHandler) AddKeyword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req AddKeywordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.groupService.AddKeyword(id, req.Keyword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "添加关键词成功",
	})
}

// RemoveKeyword 删除关键词
func (h *MonitoringGroupHandler) RemoveKeyword(c *gin.Context) {
	idStr := c.Param("id")
	groupID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的监测组ID",
		})
		return
	}

	keywordIDStr := c.Param("keyword_id")
	keywordID, err := strconv.ParseUint(keywordIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的关键词ID",
		})
		return
	}

	if err := h.groupService.RemoveKeyword(groupID, keywordID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除关键词成功",
	})
}

// GetKeywords 获取关键词列表
func (h *MonitoringGroupHandler) GetKeywords(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	keywords, err := h.groupService.GetKeywords(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取关键词列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": keywords,
	})
}

// AddExclusionWord 添加排除词
func (h *MonitoringGroupHandler) AddExclusionWord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req AddExclusionWordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.groupService.AddExclusionWord(id, req.Word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "添加排除词成功",
	})
}

// RemoveExclusionWord 删除排除词
func (h *MonitoringGroupHandler) RemoveExclusionWord(c *gin.Context) {
	idStr := c.Param("id")
	groupID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的监测组ID",
		})
		return
	}

	wordIDStr := c.Param("word_id")
	wordID, err := strconv.ParseUint(wordIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的排除词ID",
		})
		return
	}

	if err := h.groupService.RemoveExclusionWord(groupID, wordID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除排除词成功",
	})
}

// GetExclusionWords 获取排除词列表
func (h *MonitoringGroupHandler) GetExclusionWords(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	words, err := h.groupService.GetExclusionWords(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取排除词列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": words,
	})
}
