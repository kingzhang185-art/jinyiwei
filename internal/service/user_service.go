package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	pwd "sentinel-opinion-monitor/internal/pkg/password"
	"sentinel-opinion-monitor/internal/repository"
)

// UserService 用户服务接口
type UserService interface {
	CreateUser(username, password, email, nickname string) (*model.User, error)
	GetUserByID(id uint64) (*model.User, error)
	GetUsers(page, pageSize int) ([]*model.User, int64, error)
	UpdateUser(id uint64, email, nickname string, status int) error
	DeleteUser(id uint64) error
	AssignRoles(userID uint64, roleIDs []uint64) error
	ChangePassword(userID uint64, oldPassword, newPassword string) error
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser 创建用户
func (s *userService) CreateUser(username, password, email, nickname string) (*model.User, error) {
	// 检查用户名是否已存在
	_, err := s.userRepo.GetByUsername(username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := pwd.HashPassword(password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	user := &model.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
		Nickname: nickname,
		Status:   1,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return user, nil
}

// GetUserByID 根据 ID 获取用户
func (s *userService) GetUserByID(id uint64) (*model.User, error) {
	return s.userRepo.GetUserWithRoles(id)
}

// GetUsers 获取用户列表（分页）
func (s *userService) GetUsers(page, pageSize int) ([]*model.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return s.userRepo.GetAll(page, pageSize)
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(id uint64, email, nickname string, status int) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return errors.New("用户不存在")
	}

	if email != "" {
		user.Email = email
	}
	if nickname != "" {
		user.Nickname = nickname
	}
	if status > 0 {
		user.Status = status
	}

	return s.userRepo.Update(user)
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id uint64) error {
	return s.userRepo.Delete(id)
}

// AssignRoles 分配角色
func (s *userService) AssignRoles(userID uint64, roleIDs []uint64) error {
	return s.userRepo.AssignRoles(userID, roleIDs)
}

// ChangePassword 修改密码
func (s *userService) ChangePassword(userID uint64, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 验证旧密码
	if !pwd.CheckPassword(oldPassword, user.Password) {
		return errors.New("原密码错误")
	}

	// 加密新密码
	hashedPassword, err := pwd.HashPassword(newPassword)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user.Password = hashedPassword
	return s.userRepo.Update(user)
}

