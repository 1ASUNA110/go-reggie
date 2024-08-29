package dao

import (
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/utils/response"
)

type CategoryDao struct {
	BaseDao
}

var categoryDao *CategoryDao

func NewCategoryDao() *CategoryDao {
	if categoryDao == nil {
		categoryDao = &CategoryDao{
			BaseDao: NewBaseDao(),
		}
	}

	return categoryDao
}

func (m CategoryDao) CategorySave(category pojo.Category) error {
	return m.Orm.Create(&category).Error
}

func (m CategoryDao) CategoryPage(page int, pageSize int) (response.Page, error) {

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	query := m.Orm

	// 执行分页查询
	var categories []pojo.Category
	query.Offset(offset).Limit(pageSize).Order("sort ASC").Find(&categories)

	// 获取总记录数
	var total int64
	query.Model(&pojo.Category{}).Count(&total)

	categoryPage := response.Page{
		Total:    total,
		Records:  make([]interface{}, len(categories)),
		Page:     int(page),
		PageSize: int(pageSize),
	}

	// 将 employees 数据赋值给 Records
	for i, category := range categories {
		categoryPage.Records[i] = category
	}

	return categoryPage, nil

}

func (m CategoryDao) CategoryDelete(id int64) error {
	return m.Orm.Delete(pojo.Category{}, id).Error
}

// 根据分类ID修改分类信息
func (m CategoryDao) CategoryUpdateById(id int64, updateMap map[string]interface{}) error {
	return m.Orm.Model(&pojo.Category{}).Where("id = ?", id).Updates(updateMap).Error
}
