package service

import (
	"go-reggie/internal/dao"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/model/vo/response"
	"strconv"
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

func (m *CategoryService) CategorySave(dto dto.CategoryDto, employeeId int64) response.ResultCode {
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

func (m *CategoryService) CategoryPage(page int, pageSize int) (response.Page[pojo.Category], response.ResultCode) {

	// 1、调用dao层查询分类列表
	categoryPage, err := m.categoryDao.CategoryPage(page, pageSize)

	if err != nil {
		return response.Page[pojo.Category]{}, response.SERVER_ERROR()
	}

	return categoryPage, response.SUCCESS()

}

func (m *CategoryService) CategoryDelete(ID int64) response.ResultCode {
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

func (m *CategoryService) CategoryUpdate(requestMap map[string]interface{}, employeeId int64) response.ResultCode {
	// 0、校验请求参数
	// 0.1、校验ID是否为空
	idStr, ok := requestMap["id"].(string)
	if !ok {
		return response.PARAM_ERROR()
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return response.PARAM_ERROR()
	}

	// 1、构建updateMap
	updateMap := map[string]interface{}{
		"name":        requestMap["name"],
		"sort":        requestMap["sort"],
		"update_user": employeeId,
	}

	// 2、调用dao层更新分类
	err = m.categoryDao.CategoryUpdateById(id, updateMap)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}

func (m *CategoryService) CategoryList(categoryType int) ([]pojo.Category, response.ResultCode) {

	// 1、调用dao层查询分类列表
	categoryList, err := m.categoryDao.CategoryList(categoryType)

	if err != nil {
		return nil, response.SERVER_ERROR()
	}

	return categoryList, response.SUCCESS()

}
