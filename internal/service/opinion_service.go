package service

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/repository"
)

// OpinionService 舆情业务逻辑接口
type OpinionService interface {
	GetOpinionByID(id uint64) (*model.Opinion, error)
	GetAllOpinions() ([]*model.Opinion, error)
	CreateOpinion(opinion *model.Opinion) error
	UpdateOpinion(opinion *model.Opinion) error
	DeleteOpinion(id uint64) error
}

type opinionService struct {
	repo repository.OpinionRepository
}

// NewOpinionService 创建舆情业务逻辑实例
func NewOpinionService(repo repository.OpinionRepository) OpinionService {
	return &opinionService{
		repo: repo,
	}
}

// GetOpinionByID 根据 ID 获取舆情
func (s *opinionService) GetOpinionByID(id uint64) (*model.Opinion, error) {
	return s.repo.GetByID(id)
}

// GetAllOpinions 获取所有舆情
func (s *opinionService) GetAllOpinions() ([]*model.Opinion, error) {
	return s.repo.GetAll()
}

// CreateOpinion 创建舆情
func (s *opinionService) CreateOpinion(opinion *model.Opinion) error {
	return s.repo.Create(opinion)
}

// UpdateOpinion 更新舆情
func (s *opinionService) UpdateOpinion(opinion *model.Opinion) error {
	return s.repo.Update(opinion)
}

// DeleteOpinion 删除舆情
func (s *opinionService) DeleteOpinion(id uint64) error {
	return s.repo.Delete(id)
}

