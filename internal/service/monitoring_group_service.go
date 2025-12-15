package service

import (
	"errors"
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/repository"
)

// MonitoringGroupService 监测组服务接口
type MonitoringGroupService interface {
	CreateGroup(scenarioID uint64, name string, sort int) (*model.MonitoringGroup, error)
	CreateGroupWithKeywordsAndExclusionWords(scenarioID uint64, name string, sort int, keywords []string, exclusionWords []string) (*model.MonitoringGroup, error)
	GetGroupByID(id uint64) (*model.MonitoringGroup, error)
	GetGroupsByScenarioID(scenarioID uint64) ([]*model.MonitoringGroup, error)
	UpdateGroup(id uint64, name string, sort, status int) error
	DeleteGroup(id uint64) error
	GetGroupWithDetails(id uint64) (*model.MonitoringGroup, error)
	AssignChannels(groupID uint64, channelIDs []uint64) error
	AddKeyword(groupID uint64, keyword string) error
	RemoveKeyword(groupID uint64, keywordID uint64) error
	GetKeywords(groupID uint64) ([]*model.GroupKeyword, error)
	AddExclusionWord(groupID uint64, word string) error
	RemoveExclusionWord(groupID uint64, wordID uint64) error
	GetExclusionWords(groupID uint64) ([]*model.GroupExclusionWord, error)
}

type monitoringGroupService struct {
	groupRepo    repository.MonitoringGroupRepository
	scenarioRepo repository.ScenarioRepository
}

// NewMonitoringGroupService 创建监测组服务实例
func NewMonitoringGroupService(groupRepo repository.MonitoringGroupRepository, scenarioRepo repository.ScenarioRepository) MonitoringGroupService {
	return &monitoringGroupService{
		groupRepo:    groupRepo,
		scenarioRepo: scenarioRepo,
	}
}

// CreateGroup 创建监测组
func (s *monitoringGroupService) CreateGroup(scenarioID uint64, name string, sort int) (*model.MonitoringGroup, error) {
	// 验证场景是否存在
	_, err := s.scenarioRepo.GetByID(scenarioID)
	if err != nil {
		return nil, errors.New("场景不存在")
	}

	group := &model.MonitoringGroup{
		ScenarioID: scenarioID,
		Name:       name,
		Sort:       sort,
		Status:     1, // 正常状态
	}

	if err := s.groupRepo.Create(group); err != nil {
		return nil, errors.New("创建监测组失败")
	}

	return group, nil
}

// CreateGroupWithKeywordsAndExclusionWords 在事务中创建监测组及其关键词和排除词
func (s *monitoringGroupService) CreateGroupWithKeywordsAndExclusionWords(scenarioID uint64, name string, sort int, keywords []string, exclusionWords []string) (*model.MonitoringGroup, error) {
	// 验证场景是否存在
	_, err := s.scenarioRepo.GetByID(scenarioID)
	if err != nil {
		return nil, errors.New("场景不存在")
	}

	group := &model.MonitoringGroup{
		ScenarioID: scenarioID,
		Name:       name,
		Sort:       sort,
		Status:     1, // 正常状态
	}

	// 在事务中创建监测组、关键词和排除词
	if err := s.groupRepo.CreateWithKeywordsAndExclusionWords(group, keywords, exclusionWords); err != nil {
		return nil, errors.New("创建监测组失败: " + err.Error())
	}

	return group, nil
}

// GetGroupByID 根据 ID 获取监测组
func (s *monitoringGroupService) GetGroupByID(id uint64) (*model.MonitoringGroup, error) {
	return s.groupRepo.GetByID(id)
}

// GetGroupsByScenarioID 根据场景ID获取监测组列表
func (s *monitoringGroupService) GetGroupsByScenarioID(scenarioID uint64) ([]*model.MonitoringGroup, error) {
	return s.groupRepo.GetByScenarioID(scenarioID)
}

// UpdateGroup 更新监测组
func (s *monitoringGroupService) UpdateGroup(id uint64, name string, sort, status int) error {
	group, err := s.groupRepo.GetByID(id)
	if err != nil {
		return errors.New("监测组不存在")
	}

	if name != "" {
		group.Name = name
	}
	if sort >= 0 {
		group.Sort = sort
	}
	if status > 0 {
		group.Status = status
	}

	return s.groupRepo.Update(group)
}

// DeleteGroup 删除监测组
func (s *monitoringGroupService) DeleteGroup(id uint64) error {
	return s.groupRepo.Delete(id)
}

// GetGroupWithDetails 获取监测组详细信息
func (s *monitoringGroupService) GetGroupWithDetails(id uint64) (*model.MonitoringGroup, error) {
	return s.groupRepo.GetWithDetails(id)
}

// AssignChannels 分配渠道
func (s *monitoringGroupService) AssignChannels(groupID uint64, channelIDs []uint64) error {
	return s.groupRepo.AssignChannels(groupID, channelIDs)
}

// AddKeyword 添加关键词
func (s *monitoringGroupService) AddKeyword(groupID uint64, keyword string) error {
	if keyword == "" {
		return errors.New("关键词不能为空")
	}
	return s.groupRepo.AddKeyword(groupID, keyword)
}

// RemoveKeyword 删除关键词
func (s *monitoringGroupService) RemoveKeyword(groupID uint64, keywordID uint64) error {
	return s.groupRepo.RemoveKeyword(groupID, keywordID)
}

// GetKeywords 获取关键词列表
func (s *monitoringGroupService) GetKeywords(groupID uint64) ([]*model.GroupKeyword, error) {
	return s.groupRepo.GetKeywords(groupID)
}

// AddExclusionWord 添加排除词
func (s *monitoringGroupService) AddExclusionWord(groupID uint64, word string) error {
	if word == "" {
		return errors.New("排除词不能为空")
	}
	return s.groupRepo.AddExclusionWord(groupID, word)
}

// RemoveExclusionWord 删除排除词
func (s *monitoringGroupService) RemoveExclusionWord(groupID uint64, wordID uint64) error {
	return s.groupRepo.RemoveExclusionWord(groupID, wordID)
}

// GetExclusionWords 获取排除词列表
func (s *monitoringGroupService) GetExclusionWords(groupID uint64) ([]*model.GroupExclusionWord, error) {
	return s.groupRepo.GetExclusionWords(groupID)
}
