package service

import (
	"go-reggie/internal/dao"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/utils/response"
)

type CategoryService struct {
	categoryDao *dao.CategoryDao
	dishDao     *dao.DishDao
}

var categoryService *CategoryService

func NewCategoryService() *CategoryService {
	if categoryService == nil {
		categoryService = &CategoryService{
			categoryDao: dao.NewCategoryDao(),
			dishDao:     dao.NewDishDao(),
		}
	}

	return categoryService

}

func (m CategoryService) CategorySave(dto dto.CategoryDto, employeeId int64) response.ResultCode {
	// 1、创建分类对象
	category := pojo.Category{
		Type:       dto.Type,
		Name:       dto.Name,
		Sort:       dto.Sort,
		CreateUser: employeeId,
		UpdateUser: employeeId,
	}

	// 2、调用dao层保存分类
	err := m.categoryDao.CategorySave(category)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}

func (m CategoryService) CategoryPage(page int, pageSize int) (response.Page, response.ResultCode) {

	// 1、调用dao层查询分类列表
	categoryPage, err := m.categoryDao.CategoryPage(page, pageSize)

	if err != nil {
		return response.Page{}, response.SERVER_ERROR()
	}

	return categoryPage, response.SUCCESS()

}

func (m CategoryService) CategoryDelete(ID int64) response.ResultCode {
	// 1、查询当前分类是否关联了菜品，如果已关联，抛出一个业务异常
	count, err := m.dishDao.DishCountByCategoryId(ID)

	if err != nil {
		return response.SERVER_ERROR()
	}

	if count > 0 {
		// 已经关联菜品 抛出一个业务异常状态
		return response.ERROR_CATEGORY_BE_RELATED()
	}

	// 2、调用dao层删除分类
	err = m.categoryDao.CategoryDelete(ID)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}
