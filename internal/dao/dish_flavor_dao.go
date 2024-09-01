package dao

import "go-reggie/internal/model/pojo"

type DishFlavorDao struct {
	BaseDao
}

var dishFlavorDao *DishFlavorDao

func NewDishFlavorDao() *DishFlavorDao {
	if dishFlavorDao == nil {
		dishFlavorDao = &DishFlavorDao{
			BaseDao: NewBaseDao(),
		}
	}

	return dishFlavorDao
}

func (d DishFlavorDao) DishFlavorSave(flavor pojo.DishFlavor) error {

	return d.Orm.Create(&flavor).Error

}

func (d DishFlavorDao) DishFlavorGetByDishId(id int64) ([]pojo.DishFlavor, error) {
	var flavors []pojo.DishFlavor

	err := d.Orm.Where("dish_id = ?", id).Find(&flavors).Error

	if err != nil {
		return nil, err
	}

	return flavors, nil

}

func (d DishFlavorDao) DishFlavorDeleteByDishId(dishId int64) error {
	return d.Orm.Where("dish_id = ?", dishId).Delete(&pojo.DishFlavor{}).Error
}
