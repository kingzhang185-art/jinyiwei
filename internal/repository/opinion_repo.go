package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// OpinionRepository 舆情数据访问接口
type OpinionRepository interface {
	Create(opinion *model.Opinion) error
	GetByID(id uint64) (*model.Opinion, error)
	GetAll() ([]*model.Opinion, error)
	Update(opinion *model.Opinion) error
	Delete(id uint64) error
}

type opinionRepository struct {
	db *gorm.DB
}

// NewOpinionRepository 创建舆情数据访问实例
func NewOpinionRepository() OpinionRepository {
	return &opinionRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建舆情
func (r *opinionRepository) Create(opinion *model.Opinion) error {
	return r.db.Create(opinion).Error
}

// GetByID 根据 ID 获取舆情
func (r *opinionRepository) GetByID(id uint64) (*model.Opinion, error) {
	var opinion model.Opinion
	err := r.db.First(&opinion, id).Error
	if err != nil {
		return nil, err
	}
	return &opinion, nil
}

// GetAll 获取所有舆情
func (r *opinionRepository) GetAll() ([]*model.Opinion, error) {
	var opinions []*model.Opinion
	err := r.db.Find(&opinions).Error
	if err != nil {
		return nil, err
	}
	return opinions, nil
}

// Update 更新舆情
func (r *opinionRepository) Update(opinion *model.Opinion) error {
	return r.db.Save(opinion).Error
}

// Delete 删除舆情
func (r *opinionRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Opinion{}, id).Error
}

