package service

import (
	"go-reggie/internal/dao"
	"go-reggie/internal/utils/response"
)

type DishService struct {
	dishDao *dao.DishDao
}

var dishService *DishService

func NewDishService() *DishService {
	if dishService == nil {
		dishService = &DishService{
			dishDao: dao.NewDishDao(),
		}
	}

	return dishService
}

// DishPage 菜品分页查询
func (m DishService) DishPage(page int, pageSize int, name string) (response.Page, response.ResultCode) {
	// 1、调用dao层查询分类列表
	dishPage, err := m.dishDao.DishPage(page, pageSize, name)

	if err != nil {
		return response.Page{}, response.SERVER_ERROR()
	}

	return dishPage, response.SUCCESS()
}
