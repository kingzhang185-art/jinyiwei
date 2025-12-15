package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// MonitoringGroupRepository 监测组数据访问接口
type MonitoringGroupRepository interface {
	Create(group *model.MonitoringGroup) error
	CreateWithKeywordsAndExclusionWords(group *model.MonitoringGroup, keywords []string, exclusionWords []string) error
	GetByID(id uint64) (*model.MonitoringGroup, error)
	GetByScenarioID(scenarioID uint64) ([]*model.MonitoringGroup, error)
	Update(group *model.MonitoringGroup) error
	Delete(id uint64) error
	GetWithDetails(id uint64) (*model.MonitoringGroup, error)
	AssignChannels(groupID uint64, channelIDs []uint64) error
	AddKeyword(groupID uint64, keyword string) error
	RemoveKeyword(groupID uint64, keywordID uint64) error
	GetKeywords(groupID uint64) ([]*model.GroupKeyword, error)
	AddExclusionWord(groupID uint64, word string) error
	RemoveExclusionWord(groupID uint64, wordID uint64) error
	GetExclusionWords(groupID uint64) ([]*model.GroupExclusionWord, error)
}

type monitoringGroupRepository struct {
	db *gorm.DB
}

// NewMonitoringGroupRepository 创建监测组数据访问实例
func NewMonitoringGroupRepository() MonitoringGroupRepository {
	return &monitoringGroupRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建监测组
func (r *monitoringGroupRepository) Create(group *model.MonitoringGroup) error {
	return r.db.Create(group).Error
}

// CreateWithKeywordsAndExclusionWords 在事务中创建监测组及其关键词和排除词
func (r *monitoringGroupRepository) CreateWithKeywordsAndExclusionWords(group *model.MonitoringGroup, keywords []string, exclusionWords []string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 创建监测组
		if err := tx.Create(group).Error; err != nil {
			return err
		}

		// 创建关键词
		if len(keywords) > 0 {
			keywordModels := make([]model.GroupKeyword, 0, len(keywords))
			for _, keyword := range keywords {
				if keyword != "" { // 过滤空字符串
					keywordModels = append(keywordModels, model.GroupKeyword{
						GroupID: group.ID,
						Keyword: keyword,
					})
				}
			}
			if len(keywordModels) > 0 {
				if err := tx.Create(&keywordModels).Error; err != nil {
					return err
				}
			}
		}

		// 创建排除词
		if len(exclusionWords) > 0 {
			exclusionWordModels := make([]model.GroupExclusionWord, 0, len(exclusionWords))
			for _, word := range exclusionWords {
				if word != "" { // 过滤空字符串
					exclusionWordModels = append(exclusionWordModels, model.GroupExclusionWord{
						GroupID: group.ID,
						Word:    word,
					})
				}
			}
			if len(exclusionWordModels) > 0 {
				if err := tx.Create(&exclusionWordModels).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// GetByID 根据 ID 获取监测组
func (r *monitoringGroupRepository) GetByID(id uint64) (*model.MonitoringGroup, error) {
	var group model.MonitoringGroup
	err := r.db.First(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// GetByScenarioID 根据场景ID获取监测组列表
func (r *monitoringGroupRepository) GetByScenarioID(scenarioID uint64) ([]*model.MonitoringGroup, error) {
	var groups []*model.MonitoringGroup
	err := r.db.Where("scenario_id = ?", scenarioID).Order("sort ASC, id ASC").Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// Update 更新监测组
func (r *monitoringGroupRepository) Update(group *model.MonitoringGroup) error {
	return r.db.Save(group).Error
}

// Delete 删除监测组
func (r *monitoringGroupRepository) Delete(id uint64) error {
	// 删除关联的关键词
	r.db.Where("group_id = ?", id).Delete(&model.GroupKeyword{})
	// 删除关联的排除词
	r.db.Where("group_id = ?", id).Delete(&model.GroupExclusionWord{})
	// 删除关联的渠道
	r.db.Where("group_id = ?", id).Delete(&model.GroupChannel{})
	// 删除监测组
	return r.db.Delete(&model.MonitoringGroup{}, id).Error
}

// GetWithDetails 获取监测组详细信息（包含渠道、关键词、排除词）
func (r *monitoringGroupRepository) GetWithDetails(id uint64) (*model.MonitoringGroup, error) {
	var group model.MonitoringGroup
	err := r.db.Preload("Scenario").Preload("Channels").Preload("Keywords").Preload("ExclusionWords").First(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// AssignChannels 分配渠道给监测组
func (r *monitoringGroupRepository) AssignChannels(groupID uint64, channelIDs []uint64) error {
	// 先删除现有渠道关联
	r.db.Where("group_id = ?", groupID).Delete(&model.GroupChannel{})

	// 添加新渠道关联
	if len(channelIDs) > 0 {
		groupChannels := make([]model.GroupChannel, 0, len(channelIDs))
		for _, channelID := range channelIDs {
			groupChannels = append(groupChannels, model.GroupChannel{
				GroupID:   groupID,
				ChannelID: channelID,
			})
		}
		return r.db.Create(&groupChannels).Error
	}

	return nil
}

// AddKeyword 添加关键词
func (r *monitoringGroupRepository) AddKeyword(groupID uint64, keyword string) error {
	keywordModel := &model.GroupKeyword{
		GroupID: groupID,
		Keyword: keyword,
	}
	return r.db.Create(keywordModel).Error
}

// RemoveKeyword 删除关键词
func (r *monitoringGroupRepository) RemoveKeyword(groupID uint64, keywordID uint64) error {
	return r.db.Where("group_id = ? AND id = ?", groupID, keywordID).Delete(&model.GroupKeyword{}).Error
}

// GetKeywords 获取关键词列表
func (r *monitoringGroupRepository) GetKeywords(groupID uint64) ([]*model.GroupKeyword, error) {
	var keywords []*model.GroupKeyword
	err := r.db.Where("group_id = ?", groupID).Find(&keywords).Error
	if err != nil {
		return nil, err
	}
	return keywords, nil
}

// AddExclusionWord 添加排除词
func (r *monitoringGroupRepository) AddExclusionWord(groupID uint64, word string) error {
	exclusionWord := &model.GroupExclusionWord{
		GroupID: groupID,
		Word:    word,
	}
	return r.db.Create(exclusionWord).Error
}

// RemoveExclusionWord 删除排除词
func (r *monitoringGroupRepository) RemoveExclusionWord(groupID uint64, wordID uint64) error {
	return r.db.Where("group_id = ? AND id = ?", groupID, wordID).Delete(&model.GroupExclusionWord{}).Error
}

// GetExclusionWords 获取排除词列表
func (r *monitoringGroupRepository) GetExclusionWords(groupID uint64) ([]*model.GroupExclusionWord, error) {
	var words []*model.GroupExclusionWord
	err := r.db.Where("group_id = ?", groupID).Find(&words).Error
	if err != nil {
		return nil, err
	}
	return words, nil
}
