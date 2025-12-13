package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sentinel-opinion-monitor/internal/service"
)

// PermissionHandler 权限管理处理器
type PermissionHandler struct {
	permissionService service.PermissionService
}

// NewPermissionHandler 创建权限管理处理器实例
func NewPermissionHandler(permissionService service.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

// CreatePermissionRequest 创建权限请求
type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required,max=50"`
	Code        string `json:"code" binding:"required,max=100"`
	Method      string `json:"method" binding:"required,oneof=GET POST PUT DELETE PATCH"`
	Path        string `json:"path" binding:"required"`
	Description string `json:"description" binding:"omitempty,max=255"`
}

// UpdatePermissionRequest 更新权限请求
type UpdatePermissionRequest struct {
	Name        string `json:"name" binding:"omitempty,max=50"`
	Method      string `json:"method" binding:"omitempty,oneof=GET POST PUT DELETE PATCH"`
	Path        string `json:"path" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Status      int    `json:"status" binding:"omitempty,oneof=1 2"`
}

// CreatePermission 创建权限
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	permission, err := h.permissionService.CreatePermission(req.Name, req.Code, req.Method, req.Path, req.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建成功",
		"data":    permission,
	})
}

// GetPermission 获取权限详情
func (h *PermissionHandler) GetPermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	permission, err := h.permissionService.GetPermissionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "权限不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": permission,
	})
}

// GetPermissions 获取权限列表
func (h *PermissionHandler) GetPermissions(c *gin.Context) {
	permissions, err := h.permissionService.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取权限列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": permissions,
	})
}

// UpdatePermission 更新权限
func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	var req UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if err := h.permissionService.UpdatePermission(id, req.Name, req.Method, req.Path, req.Description, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
	})
}

// DeletePermission 删除权限
func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的ID",
		})
		return
	}

	if err := h.permissionService.DeletePermission(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

