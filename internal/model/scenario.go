package model

import (
	"time"
)

// Scenario 监测场景模型
type Scenario struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;comment:场景名称" json:"name"`
	TagID     uint64    `gorm:"type:bigint;not null;comment:场景标签ID" json:"tag_id"`
	Tag       Tag       `gorm:"foreignKey:TagID" json:"tag,omitempty"`
	Status    int       `gorm:"type:tinyint;default:1;comment:1-正常,2-禁用" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// 关联关系
	Groups []MonitoringGroup `gorm:"foreignKey:ScenarioID" json:"groups,omitempty"`
}

// TableName 指定表名
func (Scenario) TableName() string {
	return "scenarios"
}

// MonitoringGroup 监测组模型
type MonitoringGroup struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ScenarioID uint64    `gorm:"type:bigint;not null;comment:所属场景ID" json:"scenario_id"`
	Scenario   Scenario  `gorm:"foreignKey:ScenarioID" json:"scenario,omitempty"`
	Name       string    `gorm:"type:varchar(100);not null;comment:监测组名称" json:"name"`
	Sort       int       `gorm:"type:int;default:0;comment:排序" json:"sort"`
	Status     int       `gorm:"type:tinyint;default:1;comment:1-正常,2-禁用" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// 关联关系
	Channels       []Channel            `gorm:"many2many:group_channels;" json:"channels,omitempty"`
	Keywords       []GroupKeyword       `gorm:"foreignKey:GroupID" json:"keywords,omitempty"`
	ExclusionWords []GroupExclusionWord `gorm:"foreignKey:GroupID" json:"exclusion_words,omitempty"`
}

// TableName 指定表名
func (MonitoringGroup) TableName() string {
	return "monitoring_groups"
}

// GroupChannel 监测组-渠道关联表
type GroupChannel struct {
	GroupID   uint64 `gorm:"primaryKey"`
	ChannelID uint64 `gorm:"primaryKey"`
}

// TableName 指定表名
func (GroupChannel) TableName() string {
	return "group_channels"
}

// GroupKeyword 监测组关键词模型
type GroupKeyword struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID   uint64    `gorm:"type:bigint;not null;comment:监测组ID" json:"group_id"`
	Keyword   string    `gorm:"type:varchar(255);not null;comment:关键词" json:"keyword"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName 指定表名
func (GroupKeyword) TableName() string {
	return "group_keywords"
}

// GroupExclusionWord 监测组排除词模型
type GroupExclusionWord struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupID   uint64    `gorm:"type:bigint;not null;comment:监测组ID" json:"group_id"`
	Word      string    `gorm:"type:varchar(255);not null;comment:排除词" json:"word"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName 指定表名
func (GroupExclusionWord) TableName() string {
	return "group_exclusion_words"
}
