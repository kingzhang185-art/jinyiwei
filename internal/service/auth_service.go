package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/jwt"
	pwd "sentinel-opinion-monitor/internal/pkg/password"
	"sentinel-opinion-monitor/internal/repository"
)

// AuthService 认证服务接口
type AuthService interface {
	Register(username, password, email, nickname string) (*model.User, error)
	Login(username, password string) (string, *model.User, error)
	GetUserInfo(userID uint64) (*model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService 创建认证服务实例
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Register 用户注册
func (s *authService) Register(username, password, email, nickname string) (*model.User, error) {
	// 检查用户名是否已存在
	_, err := s.userRepo.GetByUsername(username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	if email != "" {
		_, err = s.userRepo.GetByEmail(email)
		if err == nil {
			return nil, errors.New("邮箱已存在")
		}
	}

	// 加密密码
	hashedPassword, err := pwd.HashPassword(password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户
	user := &model.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
		Nickname: nickname,
		Status:   1, // 正常状态
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return user, nil
}

// Login 用户登录
func (s *authService) Login(username, password string) (string, *model.User, error) {
	// 获取用户
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != 1 {
		return "", nil, errors.New("用户已被禁用")
	}

	// 验证密码
	if !pwd.CheckPassword(password, user.Password) {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 获取用户角色
	userWithRoles, err := s.userRepo.GetUserWithRoles(user.ID)
	if err != nil {
		return "", nil, errors.New("获取用户信息失败")
	}

	// 提取角色代码
	roles := make([]string, 0, len(userWithRoles.Roles))
	for _, role := range userWithRoles.Roles {
		if role.Status == 1 {
			roles = append(roles, role.Code)
		}
	}

	// 生成 token
	token, err := jwt.GenerateToken(user.ID, user.Username, roles)
	if err != nil {
		return "", nil, errors.New("生成token失败")
	}

	return token, user, nil
}

// GetUserInfo 获取用户信息
func (s *authService) GetUserInfo(userID uint64) (*model.User, error) {
	return s.userRepo.GetUserWithRoles(userID)
}

