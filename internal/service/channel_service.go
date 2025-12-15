package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/repository"
)

// ChannelService 渠道服务接口
type ChannelService interface {
	CreateChannel(name, code, description, icon string, sort int) (*model.Channel, error)
	GetChannelByID(id uint64) (*model.Channel, error)
	GetAllChannels() ([]*model.Channel, error)
	GetActiveChannels() ([]*model.Channel, error)
	UpdateChannel(id uint64, name, description, icon string, sort, status int) error
	DeleteChannel(id uint64) error
}

type channelService struct {
	channelRepo repository.ChannelRepository
}

// NewChannelService 创建渠道服务实例
func NewChannelService(channelRepo repository.ChannelRepository) ChannelService {
	return &channelService{
		channelRepo: channelRepo,
	}
}

// CreateChannel 创建渠道
func (s *channelService) CreateChannel(name, code, description, icon string, sort int) (*model.Channel, error) {
	// 检查渠道代码是否已存在
	_, err := s.channelRepo.GetByCode(code)
	if err == nil {
		return nil, errors.New("渠道代码已存在")
	}

	channel := &model.Channel{
		Name:        name,
		Code:        code,
		Description: description,
		Icon:        icon,
		Sort:        sort,
		Status:      1, // 正常状态
	}

	if err := s.channelRepo.Create(channel); err != nil {
		return nil, errors.New("创建渠道失败")
	}

	return channel, nil
}

// GetChannelByID 根据 ID 获取渠道
func (s *channelService) GetChannelByID(id uint64) (*model.Channel, error) {
	return s.channelRepo.GetByID(id)
}

// GetAllChannels 获取所有渠道
func (s *channelService) GetAllChannels() ([]*model.Channel, error) {
	return s.channelRepo.GetAll()
}

// GetActiveChannels 获取启用的渠道
func (s *channelService) GetActiveChannels() ([]*model.Channel, error) {
	return s.channelRepo.GetByStatus(1)
}

// UpdateChannel 更新渠道
func (s *channelService) UpdateChannel(id uint64, name, description, icon string, sort, status int) error {
	channel, err := s.channelRepo.GetByID(id)
	if err != nil {
		return errors.New("渠道不存在")
	}

	if name != "" {
		channel.Name = name
	}
	if description != "" {
		channel.Description = description
	}
	if icon != "" {
		channel.Icon = icon
	}
	if sort >= 0 {
		channel.Sort = sort
	}
	if status > 0 {
		channel.Status = status
	}

	return s.channelRepo.Update(channel)
}

// DeleteChannel 删除渠道
func (s *channelService) DeleteChannel(id uint64) error {
	return s.channelRepo.Delete(id)
}
