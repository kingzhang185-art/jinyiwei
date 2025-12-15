package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/repository"
)

// ScenarioService 场景服务接口
type ScenarioService interface {
	CreateScenario(name string, tagID uint64) (*model.Scenario, error)
	GetScenarioByID(id uint64) (*model.Scenario, error)
	GetAllScenarios() ([]*model.Scenario, error)
	GetActiveScenarios() ([]*model.Scenario, error)
	UpdateScenario(id uint64, name string, tagID uint64, status int) error
	DeleteScenario(id uint64) error
	GetScenarioWithGroups(id uint64) (*model.Scenario, error)
}

type scenarioService struct {
	scenarioRepo repository.ScenarioRepository
	tagRepo      repository.TagRepository
}

// NewScenarioService 创建场景服务实例
func NewScenarioService(scenarioRepo repository.ScenarioRepository, tagRepo repository.TagRepository) ScenarioService {
	return &scenarioService{
		scenarioRepo: scenarioRepo,
		tagRepo:      tagRepo,
	}
}

// CreateScenario 创建场景
func (s *scenarioService) CreateScenario(name string, tagID uint64) (*model.Scenario, error) {
	// 验证标签是否存在
	_, err := s.tagRepo.GetByID(tagID)
	if err != nil {
		return nil, errors.New("场景标签不存在")
	}

	scenario := &model.Scenario{
		Name:   name,
		TagID:  tagID,
		Status: 1, // 正常状态
	}

	if err := s.scenarioRepo.Create(scenario); err != nil {
		return nil, errors.New("创建场景失败")
	}

	return scenario, nil
}

// GetScenarioByID 根据 ID 获取场景
func (s *scenarioService) GetScenarioByID(id uint64) (*model.Scenario, error) {
	return s.scenarioRepo.GetByID(id)
}

// GetAllScenarios 获取所有场景
func (s *scenarioService) GetAllScenarios() ([]*model.Scenario, error) {
	return s.scenarioRepo.GetAll()
}

// GetActiveScenarios 获取启用的场景
func (s *scenarioService) GetActiveScenarios() ([]*model.Scenario, error) {
	return s.scenarioRepo.GetByStatus(1)
}

// UpdateScenario 更新场景
func (s *scenarioService) UpdateScenario(id uint64, name string, tagID uint64, status int) error {
	scenario, err := s.scenarioRepo.GetByID(id)
	if err != nil {
		return errors.New("场景不存在")
	}

	if name != "" {
		scenario.Name = name
	}
	if tagID > 0 {
		// 验证标签是否存在
		_, err := s.tagRepo.GetByID(tagID)
		if err != nil {
			return errors.New("场景标签不存在")
		}
		scenario.TagID = tagID
	}
	if status > 0 {
		scenario.Status = status
	}

	return s.scenarioRepo.Update(scenario)
}

// DeleteScenario 删除场景
func (s *scenarioService) DeleteScenario(id uint64) error {
	return s.scenarioRepo.Delete(id)
}

// GetScenarioWithGroups 获取场景及其监测组
func (s *scenarioService) GetScenarioWithGroups(id uint64) (*model.Scenario, error) {
	return s.scenarioRepo.GetWithGroups(id)
}
