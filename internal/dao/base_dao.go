package dao

import (
	"go-reggie/internal/global"
	"gorm.io/gorm"
)

type BaseDao struct {
	Orm *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		Orm: global.DB,
	}
}
