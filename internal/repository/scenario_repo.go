package repository

import (
	"sentinel-opinion-monitor/internal/model"
	"sentinel-opinion-monitor/internal/pkg/mysql"

	"gorm.io/gorm"
)

// ScenarioRepository 场景数据访问接口
type ScenarioRepository interface {
	Create(scenario *model.Scenario) error
	GetByID(id uint64) (*model.Scenario, error)
	GetAll() ([]*model.Scenario, error)
	GetByStatus(status int) ([]*model.Scenario, error)
	Update(scenario *model.Scenario) error
	Delete(id uint64) error
	GetWithGroups(id uint64) (*model.Scenario, error)
}

type scenarioRepository struct {
	db *gorm.DB
}

// NewScenarioRepository 创建场景数据访问实例
func NewScenarioRepository() ScenarioRepository {
	return &scenarioRepository{
		db: mysql.GetDB(),
	}
}

// Create 创建场景
func (r *scenarioRepository) Create(scenario *model.Scenario) error {
	return r.db.Create(scenario).Error
}

// GetByID 根据 ID 获取场景
func (r *scenarioRepository) GetByID(id uint64) (*model.Scenario, error) {
	var scenario model.Scenario
	err := r.db.First(&scenario, id).Error
	if err != nil {
		return nil, err
	}
	return &scenario, nil
}

// GetAll 获取所有场景
func (r *scenarioRepository) GetAll() ([]*model.Scenario, error) {
	var scenarios []*model.Scenario
	err := r.db.Preload("Tag").Order("id DESC").Find(&scenarios).Error
	if err != nil {
		return nil, err
	}
	return scenarios, nil
}

// GetByStatus 根据状态获取场景
func (r *scenarioRepository) GetByStatus(status int) ([]*model.Scenario, error) {
	var scenarios []*model.Scenario
	err := r.db.Where("status = ?", status).Preload("Tag").Order("id DESC").Find(&scenarios).Error
	if err != nil {
		return nil, err
	}
	return scenarios, nil
}

// Update 更新场景
func (r *scenarioRepository) Update(scenario *model.Scenario) error {
	return r.db.Save(scenario).Error
}

// Delete 删除场景
func (r *scenarioRepository) Delete(id uint64) error {
	// 先删除关联的监测组
	r.db.Where("scenario_id = ?", id).Delete(&model.MonitoringGroup{})
	return r.db.Delete(&model.Scenario{}, id).Error
}

// GetWithGroups 获取场景及其监测组
func (r *scenarioRepository) GetWithGroups(id uint64) (*model.Scenario, error) {
	var scenario model.Scenario
	err := r.db.Preload("Tag").Preload("Groups").Preload("Groups.Channels").Preload("Groups.Keywords").Preload("Groups.ExclusionWords").First(&scenario, id).Error
	if err != nil {
		return nil, err
	}
	return &scenario, nil
}
