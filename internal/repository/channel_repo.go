package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// ChannelRepository 渠道数据访问接口
type ChannelRepository interface {
	Create(channel *model.Channel) error
	GetByID(id uint64) (*model.Channel, error)
	GetByCode(code string) (*model.Channel, error)
	GetAll() ([]*model.Channel, error)
	GetByStatus(status int) ([]*model.Channel, error)
	Update(channel *model.Channel) error
	Delete(id uint64) error
}

type channelRepository struct {
	db *gorm.DB
}

// NewChannelRepository 创建渠道数据访问实例
func NewChannelRepository() ChannelRepository {
	return &channelRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建渠道
func (r *channelRepository) Create(channel *model.Channel) error {
	return r.db.Create(channel).Error
}

// GetByID 根据 ID 获取渠道
func (r *channelRepository) GetByID(id uint64) (*model.Channel, error) {
	var channel model.Channel
	err := r.db.First(&channel, id).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

// GetByCode 根据代码获取渠道
func (r *channelRepository) GetByCode(code string) (*model.Channel, error) {
	var channel model.Channel
	err := r.db.Where("code = ?", code).First(&channel).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

// GetAll 获取所有渠道
func (r *channelRepository) GetAll() ([]*model.Channel, error) {
	var channels []*model.Channel
	err := r.db.Order("sort ASC, id ASC").Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

// GetByStatus 根据状态获取渠道
func (r *channelRepository) GetByStatus(status int) ([]*model.Channel, error) {
	var channels []*model.Channel
	err := r.db.Where("status = ?", status).Order("sort ASC, id ASC").Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

// Update 更新渠道
func (r *channelRepository) Update(channel *model.Channel) error {
	return r.db.Save(channel).Error
}

// Delete 删除渠道
func (r *channelRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Channel{}, id).Error
}
