package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/repository"
)

// RoleService 角色服务接口
type RoleService interface {
	CreateRole(name, code, description string) (*model.Role, error)
	GetRoleByID(id uint64) (*model.Role, error)
	GetAllRoles() ([]*model.Role, error)
	UpdateRole(id uint64, name, description string, status int) error
	DeleteRole(id uint64) error
	AssignPermissions(roleID uint64, permissionIDs []uint64) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

// NewRoleService 创建角色服务实例
func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

// CreateRole 创建角色
func (s *roleService) CreateRole(name, code, description string) (*model.Role, error) {
	// 检查角色代码是否已存在
	_, err := s.roleRepo.GetByCode(code)
	if err == nil {
		return nil, errors.New("角色代码已存在")
	}

	role := &model.Role{
		Name:        name,
		Code:        code,
		Description: description,
		Status:      1,
	}

	if err := s.roleRepo.Create(role); err != nil {
		return nil, errors.New("创建角色失败")
	}

	return role, nil
}

// GetRoleByID 根据 ID 获取角色
func (s *roleService) GetRoleByID(id uint64) (*model.Role, error) {
	return s.roleRepo.GetRoleWithPermissions(id)
}

// GetAllRoles 获取所有角色
func (s *roleService) GetAllRoles() ([]*model.Role, error) {
	return s.roleRepo.GetAll()
}

// UpdateRole 更新角色
func (s *roleService) UpdateRole(id uint64, name, description string, status int) error {
	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		return errors.New("角色不存在")
	}

	if name != "" {
		role.Name = name
	}
	if description != "" {
		role.Description = description
	}
	if status > 0 {
		role.Status = status
	}

	return s.roleRepo.Update(role)
}

// DeleteRole 删除角色
func (s *roleService) DeleteRole(id uint64) error {
	return s.roleRepo.Delete(id)
}

// AssignPermissions 分配权限
func (s *roleService) AssignPermissions(roleID uint64, permissionIDs []uint64) error {
	return s.roleRepo.AssignPermissions(roleID, permissionIDs)
}

