package pojo

import (
	"time"
)

// Common 包含所有模型共有的字段
type BasePojo struct {
	ID         int64     `json:"id" gorm:"primaryKey;column:id"`                 // 主键
	CreateTime time.Time `json:"create_time" gorm:"not null;column:create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time" gorm:"not null;column:update_time"` // 更新时间
}
