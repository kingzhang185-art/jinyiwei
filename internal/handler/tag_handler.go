package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/service"
)

// TagHandler 标签管理处理器
type TagHandler struct {
	tagService service.TagService
}

// NewTagHandler 创建标签管理处理器实例
func NewTagHandler(tagService service.TagService) *TagHandler {
	return &TagHandler{
		tagService: tagService,
	}
}

// CreateTagRequest 创建标签请求
type CreateTagRequest struct {
	Name        string `json:"name" binding:"required,max=50"`
	Code        string `json:"code" binding:"required,max=50"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Type        string `json:"type" binding:"omitempty,max=20"`
	Sort        int    `json:"sort" binding:"omitempty"`
}

// UpdateTagRequest 更新标签请求
type UpdateTagRequest struct {
	Name        string `json:"name" binding:"omitempty,max=50"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Sort        int    `json:"sort" binding:"omitempty"`
	Status      int    `json:"status" binding:"omitempty,oneof=1 2"`
}

// CreateTag 创建标签
func (h *TagHandler) CreateTag(c *gin.Context) {
	var req CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	tag, err := h.tagService.CreateTag(req.Name, req.Code, req.Description, req.Type, req.Sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建成功",
		"data":    tag,
	})
}

// GetTag 获取标签详情
func (h *TagHandler) GetTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	tag, err := h.tagService.GetTagByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "标签不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tag,
	})
}

// GetTags 获取标签列表
func (h *TagHandler) GetTags(c *gin.Context) {
	tagType := c.Query("type")
	status := c.Query("status")

	var tags []*model.Tag
	var err error

	if status == "active" || status == "1" {
		// 只获取启用的标签
		tags, err = h.tagService.GetActiveTags(tagType)
	} else {
		// 获取所有标签
		tags, err = h.tagService.GetAllTags(tagType)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取标签列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tags,
	})
}

// UpdateTag 更新标签
func (h *TagHandler) UpdateTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.tagService.UpdateTag(id, req.Name, req.Description, req.Sort, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// DeleteTag 删除标签
func (h *TagHandler) DeleteTag(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	if err := h.tagService.DeleteTag(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
