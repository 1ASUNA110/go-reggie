package service

import (
	"go-reggie/internal/dao"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/utils/response"
)

type CategoryService struct {
	categoryDao *dao.CategoryDao
}

var categoryService *CategoryService

func NewCategoryService() *CategoryService {
	if categoryService == nil {
		categoryService = &CategoryService{
			categoryDao: dao.NewCategoryDao(),
		}
	}
	return categoryService

}

func (s CategoryService) CategorySave(dto dto.CategoryDto, employeeId int64) response.ResultCode {
	// 1、创建分类对象
	category := pojo.Category{
		Type:       dto.Type,
		Name:       dto.Name,
		Sort:       dto.Sort,
		CreateUser: employeeId,
		UpdateUser: employeeId,
	}

	// 2、调用dao层保存分类
	err := s.categoryDao.CategorySave(category)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}

func (s CategoryService) CategoryPage(page int, pageSize int) (response.Page, response.ResultCode) {

	// 1、调用dao层查询分类列表
	categoryPage, err := s.categoryDao.CategoryPage(page, pageSize)

	if err != nil {
		return response.Page{}, response.SERVER_ERROR()
	}

	return categoryPage, response.SUCCESS()

}
