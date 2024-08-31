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

func (d DishFlavorDao) DishFlavorSave(flavors []pojo.DishFlavor) error {

	return d.Orm.Create(&flavors).Error

}

func (d DishFlavorDao) DishFlavorGetByDishId(id int64) ([]pojo.DishFlavor, error) {
	var flavors []pojo.DishFlavor

	err := d.Orm.Where("dish_id = ?", id).Find(&flavors).Error

	if err != nil {
		return nil, err
	}

	return flavors, nil

}
