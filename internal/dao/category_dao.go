package dao

import (
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/model/vo/response"
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

func (m *CategoryDao) CategorySave(category pojo.Category) error {
	return m.Orm.Create(&category).Error
}

func (m *CategoryDao) CategoryPage(page int, pageSize int) (response.Page[pojo.Category], error) {

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

	categoryPage := response.Page[pojo.Category]{
		Total:    total,
		Records:  categories,
		Page:     int(page),
		PageSize: int(pageSize),
	}

	return categoryPage, nil

}

func (m *CategoryDao) CategoryDelete(id int64) error {
	return m.Orm.Delete(pojo.Category{}, id).Error
}

// 根据分类ID修改分类信息
func (m *CategoryDao) CategoryUpdateById(id int64, updateMap map[string]interface{}) error {
	return m.Orm.Model(&pojo.Category{}).Where("id = ?", id).Updates(updateMap).Error
}

func (m *CategoryDao) CategoryGetById(id int64) (pojo.Category, error) {
	var category pojo.Category

	err := m.Orm.Where("id = ?", id).First(&category).Error

	return category, err
}

func (m *CategoryDao) CategoryList(categoryType int) ([]pojo.Category, error) {
	var categories []pojo.Category

	err := m.Orm.Where("type = ?", categoryType).Find(&categories).Error

	return categories, err

}
