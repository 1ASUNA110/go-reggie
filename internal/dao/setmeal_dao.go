package dao

import (
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/model/vo/response"
)

type SetmealDao struct {
	BaseDao
}

var setmealDao *SetmealDao

func NewSetmealDao() *SetmealDao {
	if setmealDao == nil {
		setmealDao = &SetmealDao{
			BaseDao: NewBaseDao(),
		}
	}

	return setmealDao
}

func (m *SetmealDao) SetmealSave(setmeal pojo.Setmeal) error {
	return m.Orm.Create(&setmeal).Error
}

func (m *SetmealDao) SetmealPage(page int, pageSize int, name string) (response.Page[pojo.Setmeal], error) {

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	query := m.Orm

	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}

	// 执行分页查询
	var categories []pojo.Setmeal
	query.Offset(offset).Limit(pageSize).Find(&categories)

	// 获取总记录数
	var total int64
	query.Model(&pojo.Setmeal{}).Count(&total)

	setmealPage := response.Page[pojo.Setmeal]{
		Total:    total,
		Records:  categories,
		Page:     int(page),
		PageSize: int(pageSize),
	}

	return setmealPage, nil

}

func (m *SetmealDao) SetmealDelete(id int64) error {
	return m.Orm.Delete(pojo.Setmeal{}, id).Error
}

// 根据分类ID修改分类信息
func (m *SetmealDao) SetmealUpdateById(id int64, updateMap map[string]interface{}) error {
	return m.Orm.Model(&pojo.Setmeal{}).Where("id = ?", id).Updates(updateMap).Error
}

func (m *SetmealDao) SetmealGetById(id int64) (pojo.Setmeal, error) {
	var setmeal pojo.Setmeal

	err := m.Orm.Where("id = ?", id).First(&setmeal).Error

	return setmeal, err
}

func (m *SetmealDao) SetmealList(setmealType int) ([]pojo.Setmeal, error) {
	var categories []pojo.Setmeal

	err := m.Orm.Where("type = ?", setmealType).Find(&categories).Error

	return categories, err

}

func (m *SetmealDao) SetmealUpdateStatus(id int64, status int) interface{} {

	return m.Orm.Model(&pojo.Setmeal{}).Where("id = ?", id).Update("status", status).Error

}

func (m *SetmealDao) DishDelete(id int64) interface{} {

	return m.Orm.Delete(pojo.Setmeal{}, id).Error

}
