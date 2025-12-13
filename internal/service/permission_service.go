package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/repository"
)

// PermissionService 权限服务接口
type PermissionService interface {
	CreatePermission(name, code, method, path, description string) (*model.Permission, error)
	GetPermissionByID(id uint64) (*model.Permission, error)
	GetAllPermissions() ([]*model.Permission, error)
	UpdatePermission(id uint64, name, method, path, description string, status int) error
	DeletePermission(id uint64) error
}

type permissionService struct {
	permissionRepo repository.PermissionRepository
}

// NewPermissionService 创建权限服务实例
func NewPermissionService(permissionRepo repository.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

// CreatePermission 创建权限
func (s *permissionService) CreatePermission(name, code, method, path, description string) (*model.Permission, error) {
	// 检查权限代码是否已存在
	_, err := s.permissionRepo.GetByCode(code)
	if err == nil {
		return nil, errors.New("权限代码已存在")
	}

	permission := &model.Permission{
		Name:        name,
		Code:        code,
		Method:      method,
		Path:        path,
		Description: description,
		Status:      1,
	}

	if err := s.permissionRepo.Create(permission); err != nil {
		return nil, errors.New("创建权限失败")
	}

	return permission, nil
}

// GetPermissionByID 根据 ID 获取权限
func (s *permissionService) GetPermissionByID(id uint64) (*model.Permission, error) {
	return s.permissionRepo.GetByID(id)
}

// GetAllPermissions 获取所有权限
func (s *permissionService) GetAllPermissions() ([]*model.Permission, error) {
	return s.permissionRepo.GetAll()
}

// UpdatePermission 更新权限
func (s *permissionService) UpdatePermission(id uint64, name, method, path, description string, status int) error {
	permission, err := s.permissionRepo.GetByID(id)
	if err != nil {
		return errors.New("权限不存在")
	}

	if name != "" {
		permission.Name = name
	}
	if method != "" {
		permission.Method = method
	}
	if path != "" {
		permission.Path = path
	}
	if description != "" {
		permission.Description = description
	}
	if status > 0 {
		permission.Status = status
	}

	return s.permissionRepo.Update(permission)
}

// DeletePermission 删除权限
func (s *permissionService) DeletePermission(id uint64) error {
	return s.permissionRepo.Delete(id)
}

