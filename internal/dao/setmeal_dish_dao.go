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
func (d SetmealDishDao) SetmealDishSave(setmealDish pojo.SetmealDish) error {
	return d.Orm.Create(&setmealDish).Error
}

func (d SetmealDishDao) SetmealDishGetBySetmealId(id int64) ([]pojo.SetmealDish, error) {
	var setmealDishes []pojo.SetmealDish

	err := d.Orm.Where("setmeal_id = ?", id).Find(&setmealDishes).Error

	if err != nil {
		return nil, err
	}

	return setmealDishes, nil
}

func (d SetmealDishDao) SetmealDishDeleteBySetmealId(id int64) error {
	return d.Orm.Where("setmeal_id = ?", id).Delete(&pojo.SetmealDish{}).Error
}
