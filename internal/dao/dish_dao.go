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

	err := m.Orm.Table("dish").Where("category_id = ?", categoryId).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *DishDao) DishPage(page int, pageSize int, name string) (response.Page, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	query := m.Orm

	// 构建查询条件
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// 执行分页查询
	var dishes []pojo.Dish
	query.Offset(offset).Limit(pageSize).Order("update_time ASC").Find(&dishes)

	// 获取总记录数
	var total int64
	query.Model(&pojo.Dish{}).Count(&total)

	dishPage := response.Page{
		Total:    total,
		Records:  make([]interface{}, len(dishes)),
		Page:     int(page),
		PageSize: int(pageSize),
	}

	// 将 dishes 数据赋值给 Records
	for i, category := range dishes {
		dishPage.Records[i] = category
	}

	return dishPage, nil
}
