package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// PermissionRepository 权限数据访问接口
type PermissionRepository interface {
	Create(permission *model.Permission) error
	GetByID(id uint64) (*model.Permission, error)
	GetByCode(code string) (*model.Permission, error)
	GetAll() ([]*model.Permission, error)
	Update(permission *model.Permission) error
	Delete(id uint64) error
}

type permissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository 创建权限数据访问实例
func NewPermissionRepository() PermissionRepository {
	return &permissionRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建权限
func (r *permissionRepository) Create(permission *model.Permission) error {
	return r.db.Create(permission).Error
}

// GetByID 根据 ID 获取权限
func (r *permissionRepository) GetByID(id uint64) (*model.Permission, error) {
	var permission model.Permission
	err := r.db.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetByCode 根据代码获取权限
func (r *permissionRepository) GetByCode(code string) (*model.Permission, error) {
	var permission model.Permission
	err := r.db.Where("code = ?", code).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetAll 获取所有权限
func (r *permissionRepository) GetAll() ([]*model.Permission, error) {
	var permissions []*model.Permission
	err := r.db.Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// Update 更新权限
func (r *permissionRepository) Update(permission *model.Permission) error {
	return r.db.Save(permission).Error
}

// Delete 删除权限
func (r *permissionRepository) Delete(id uint64) error {
	// 先删除角色权限关联
	r.db.Where("permission_id = ?", id).Delete(&model.RolePermission{})
	return r.db.Delete(&model.Permission{}, id).Error
}
