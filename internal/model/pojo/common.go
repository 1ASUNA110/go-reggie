package pojo

import (
	"go-reggie/internal/utils"
	"gorm.io/gorm"
	"time"
)

// BasePojo 包含所有模型共有的字段
type BasePojo struct {
	ID         int64     `json:"id,string" gorm:"primaryKey;column:id"`               // 主键
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime;column:create_time"` // 创建时间
	UpdateTime time.Time `json:"updateTime" gorm:"autoUpdateTime;column:update_time"` // 更新时间
}

func (basePojo *BasePojo) BeforeCreate(tx *gorm.DB) (err error) {
	if basePojo.ID == 0 {
		basePojo.ID = utils.GenerateID()
	}
	return
}
