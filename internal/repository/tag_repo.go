package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// TagRepository 标签数据访问接口
type TagRepository interface {
	Create(tag *model.Tag) error
	GetByID(id uint64) (*model.Tag, error)
	GetByCode(code string) (*model.Tag, error)
	GetAll(tagType string) ([]*model.Tag, error)
	GetByTypeAndStatus(tagType string, status int) ([]*model.Tag, error)
	Update(tag *model.Tag) error
	Delete(id uint64) error
}

type tagRepository struct {
	db *gorm.DB
}

// NewTagRepository 创建标签数据访问实例
func NewTagRepository() TagRepository {
	return &tagRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建标签
func (r *tagRepository) Create(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

// GetByID 根据 ID 获取标签
func (r *tagRepository) GetByID(id uint64) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetByCode 根据代码获取标签
func (r *tagRepository) GetByCode(code string) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.Where("code = ?", code).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetAll 获取所有标签
func (r *tagRepository) GetAll(tagType string) ([]*model.Tag, error) {
	var tags []*model.Tag
	query := r.db.Model(&model.Tag{})

	if tagType != "" {
		query = query.Where("type = ?", tagType)
	}

	err := query.Order("sort ASC, id ASC").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// GetByTypeAndStatus 根据类型和状态获取标签
func (r *tagRepository) GetByTypeAndStatus(tagType string, status int) ([]*model.Tag, error) {
	var tags []*model.Tag
	query := r.db.Model(&model.Tag{})

	if tagType != "" {
		query = query.Where("type = ?", tagType)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	err := query.Order("sort ASC, id ASC").Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// Update 更新标签
func (r *tagRepository) Update(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

// Delete 删除标签
func (r *tagRepository) Delete(id uint64) error {
	return r.db.Delete(&model.Tag{}, id).Error
}
