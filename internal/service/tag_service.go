package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/repository"
)

// TagService 标签服务接口
type TagService interface {
	CreateTag(name, code, description, tagType string, sort int) (*model.Tag, error)
	GetTagByID(id uint64) (*model.Tag, error)
	GetAllTags(tagType string) ([]*model.Tag, error)
	GetActiveTags(tagType string) ([]*model.Tag, error)
	UpdateTag(id uint64, name, description string, sort, status int) error
	DeleteTag(id uint64) error
}

type tagService struct {
	tagRepo repository.TagRepository
}

// NewTagService 创建标签服务实例
func NewTagService(tagRepo repository.TagRepository) TagService {
	return &tagService{
		tagRepo: tagRepo,
	}
}

// CreateTag 创建标签
func (s *tagService) CreateTag(name, code, description, tagType string, sort int) (*model.Tag, error) {
	// 检查标签代码是否已存在
	_, err := s.tagRepo.GetByCode(code)
	if err == nil {
		return nil, errors.New("标签代码已存在")
	}

	if tagType == "" {
		tagType = "scene"
	}

	tag := &model.Tag{
		Name:        name,
		Code:        code,
		Description: description,
		Type:        tagType,
		Sort:        sort,
		Status:      1, // 正常状态
	}

	if err := s.tagRepo.Create(tag); err != nil {
		return nil, errors.New("创建标签失败")
	}

	return tag, nil
}

// GetTagByID 根据 ID 获取标签
func (s *tagService) GetTagByID(id uint64) (*model.Tag, error) {
	return s.tagRepo.GetByID(id)
}

// GetAllTags 获取所有标签
func (s *tagService) GetAllTags(tagType string) ([]*model.Tag, error) {
	return s.tagRepo.GetAll(tagType)
}

// GetActiveTags 获取启用的标签
func (s *tagService) GetActiveTags(tagType string) ([]*model.Tag, error) {
	return s.tagRepo.GetByTypeAndStatus(tagType, 1)
}

// UpdateTag 更新标签
func (s *tagService) UpdateTag(id uint64, name, description string, sort, status int) error {
	tag, err := s.tagRepo.GetByID(id)
	if err != nil {
		return errors.New("标签不存在")
	}

	if name != "" {
		tag.Name = name
	}
	if description != "" {
		tag.Description = description
	}
	if sort >= 0 {
		tag.Sort = sort
	}
	if status > 0 {
		tag.Status = status
	}

	return s.tagRepo.Update(tag)
}

// DeleteTag 删除标签
func (s *tagService) DeleteTag(id uint64) error {
	return s.tagRepo.Delete(id)
}
