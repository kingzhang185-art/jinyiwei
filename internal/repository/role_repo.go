package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// RoleRepository 角色数据访问接口
type RoleRepository interface {
	Create(role *model.Role) error
	GetByID(id uint64) (*model.Role, error)
	GetByCode(code string) (*model.Role, error)
	GetAll() ([]*model.Role, error)
	Update(role *model.Role) error
	Delete(id uint64) error
	AssignPermissions(roleID uint64, permissionIDs []uint64) error
	GetRoleWithPermissions(roleID uint64) (*model.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository 创建角色数据访问实例
func NewRoleRepository() RoleRepository {
	return &roleRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建角色
func (r *roleRepository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

// GetByID 根据 ID 获取角色
func (r *roleRepository) GetByID(id uint64) (*model.Role, error) {
	var role model.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// GetByCode 根据代码获取角色
func (r *roleRepository) GetByCode(code string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("code = ?", code).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// GetAll 获取所有角色
func (r *roleRepository) GetAll() ([]*model.Role, error) {
	var roles []*model.Role
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// Update 更新角色
func (r *roleRepository) Update(role *model.Role) error {
	return r.db.Save(role).Error
}

// Delete 删除角色
func (r *roleRepository) Delete(id uint64) error {
	// 先删除角色权限关联
	r.db.Where("role_id = ?", id).Delete(&model.RolePermission{})
	// 删除用户角色关联
	r.db.Where("role_id = ?", id).Delete(&model.UserRole{})
	return r.db.Delete(&model.Role{}, id).Error
}

// AssignPermissions 分配权限给角色
func (r *roleRepository) AssignPermissions(roleID uint64, permissionIDs []uint64) error {
	// 先删除现有权限
	r.db.Where("role_id = ?", roleID).Delete(&model.RolePermission{})

	// 添加新权限
	if len(permissionIDs) > 0 {
		rolePermissions := make([]model.RolePermission, 0, len(permissionIDs))
		for _, permissionID := range permissionIDs {
			rolePermissions = append(rolePermissions, model.RolePermission{
				RoleID:       roleID,
				PermissionID: permissionID,
			})
		}
		return r.db.Create(&rolePermissions).Error
	}

	return nil
}

// GetRoleWithPermissions 获取角色及其权限
func (r *roleRepository) GetRoleWithPermissions(roleID uint64) (*model.Role, error) {
	var role model.Role
	err := r.db.Preload("Permissions").First(&role, roleID).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

