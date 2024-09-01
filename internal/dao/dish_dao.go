package dao

import (
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/model/vo/response"
)

type DishDao struct {
	BaseDao
}

var dishDao *DishDao

func NewDishDao() *DishDao {
	if dishDao == nil {
		dishDao = &DishDao{
			BaseDao: NewBaseDao(),
		}
	}

	return dishDao
}

func (m *DishDao) DishCountByCategoryId(categoryId int64) (int64, error) {
	var count int64

	err := m.Orm.Model(pojo.Dish{}).Where("category_id = ?", categoryId).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *DishDao) DishPage(page int, pageSize int, name string) (response.Page[pojo.Dish], error) {
	// 计算偏移量
	offset := (page - 1) * pageSize
	query := m.Orm

	// 构建查询条件
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// 执行分页查询
	var dishes []pojo.Dish
	query.Offset(offset).Limit(pageSize).Order("update_time DESC").Find(&dishes)

	// 获取总记录数
	var total int64
	query.Model(&pojo.Dish{}).Count(&total)

	dishPage := response.Page[pojo.Dish]{
		Total:    total,
		Records:  dishes,
		Page:     page,
		PageSize: pageSize,
	}

	// 将 dishes 数据赋值给 Records
	for i, category := range dishes {
		dishPage.Records[i] = category
	}

	return dishPage, nil
}

// DishSave 菜品保存
func (m *DishDao) DishDelete(id int64) error {
	return m.Orm.Where("id = ?", id).Delete(pojo.Dish{}).Error
}

// DishUpdateStatus 菜品状态更新
func (m *DishDao) DishUpdateStatus(id int64, status int) interface{} {
	return m.Orm.Model(pojo.Dish{}).Where("id = ?", id).Update("status", status).Error
}

func (m *DishDao) DishSave(dish pojo.Dish) error {
	return m.Orm.Create(&dish).Error
}

func (m *DishDao) DishGetById(id int64) (pojo.Dish, error) {
	var dish pojo.Dish

	err := m.Orm.Where("id = ?", id).First(&dish).Error

	if err != nil {
		return pojo.Dish{}, err
	}

	return dish, nil

}

func (m *DishDao) DishList(categoryId int64) ([]pojo.Dish, error) {
	var dishes []pojo.Dish

	err := m.Orm.Where("category_id = ?", categoryId).Find(&dishes).Error

	if err != nil {
		return nil, err
	}

	return dishes, nil
}
