package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uint64) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll(page, pageSize int) ([]*model.User, int64, error)
	Update(user *model.User) error
	Delete(id uint64) error
	AssignRoles(userID uint64, roleIDs []uint64) error
	GetUserWithRoles(userID uint64) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户数据访问实例
func NewUserRepository() UserRepository {
	return &userRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建用户
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByID 根据 ID 获取用户
func (r *userRepository) GetByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAll 获取所有用户（分页）
func (r *userRepository) GetAll(page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	offset := (page - 1) * pageSize
	err := r.db.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Update 更新用户
func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
func (r *userRepository) Delete(id uint64) error {
	// 先删除用户角色关联
	r.db.Where("user_id = ?", id).Delete(&model.UserRole{})
	return r.db.Delete(&model.User{}, id).Error
}

// AssignRoles 分配角色给用户
func (r *userRepository) AssignRoles(userID uint64, roleIDs []uint64) error {
	// 先删除现有角色
	r.db.Where("user_id = ?", userID).Delete(&model.UserRole{})

	// 添加新角色
	if len(roleIDs) > 0 {
		userRoles := make([]model.UserRole, 0, len(roleIDs))
		for _, roleID := range roleIDs {
			userRoles = append(userRoles, model.UserRole{
				UserID: userID,
				RoleID: roleID,
			})
		}
		return r.db.Create(&userRoles).Error
	}

	return nil
}

// GetUserWithRoles 获取用户及其角色
func (r *userRepository) GetUserWithRoles(userID uint64) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Roles").Preload("Roles.Permissions").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

