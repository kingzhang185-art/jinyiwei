package model

import (
	"time"
)

// Channel 渠道模型
type Channel struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	Code        string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Icon        string    `gorm:"type:varchar(255);comment:图标URL或图标标识" json:"icon"`
	Sort        int       `gorm:"type:int;default:0;comment:排序" json:"sort"`
	Status      int       `gorm:"type:tinyint;default:1;comment:1-正常,2-禁用" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Channel) TableName() string {
	return "channels"
}
