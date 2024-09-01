package dao

import "go-reggie/internal/model/pojo"

type SetmealDishDao struct {
	BaseDao
}

var setmealDishDao *SetmealDishDao

func NewSetmealDishDao() *SetmealDishDao {
	if setmealDishDao == nil {
		setmealDishDao = &SetmealDishDao{
			BaseDao: NewBaseDao(),
		}
	}

	return setmealDishDao
}

// SetmealDishSave 保存套餐菜品
func (d SetmealDishDao) SetmealDishSave(dish pojo.SetmealDish) interface{} {
	return d.Orm.Create(&dish).Error
}
