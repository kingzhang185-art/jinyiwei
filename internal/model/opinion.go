package model

import (
	"time"
)

// Opinion 舆情模型
type Opinion struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Source    string    `gorm:"type:varchar(255);not null" json:"source"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (Opinion) TableName() string {
	return "opinions"
}

