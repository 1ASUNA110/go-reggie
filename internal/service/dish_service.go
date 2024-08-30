package service

import (
	"github.com/jinzhu/copier"
	"go-reggie/internal/dao"
	vo "go-reggie/internal/model/vo"
	"go-reggie/internal/model/vo/response"
)

type DishService struct {
	dishDao     *dao.DishDao
	categoryDao *dao.CategoryDao
}

var dishService *DishService

func NewDishService() *DishService {
	if dishService == nil {
		dishService = &DishService{
			dishDao:     dao.NewDishDao(),
			categoryDao: dao.NewCategoryDao(),
		}
	}

	return dishService
}

// DishPage 菜品分页查询
func (m *DishService) DishPage(page int, pageSize int, name string) (response.Page[vo.DishVo], response.ResultCode) {
	// 1、调用dao层查询分类列表
	dishPage, err := m.dishDao.DishPage(page, pageSize, name)

	if err != nil {
		return response.Page[vo.DishVo]{}, response.SERVER_ERROR()
	}

	// 2、给菜品添加分类信息
	// 这里用最简单的做法直接遍历每个菜品，然后查询每个菜品的分类信息
	// 但是这样会导致查询数据库的次数过多，可以通过一次性查询所有分类，然后通过map的方式来获取分类信息

	records := dishPage.Records

	// 定义一个vo的切片
	dishVos := []vo.DishVo{}

	for i := 0; i < len(records); i++ {
		dishVo := vo.DishVo{}

		dish := records[i]
		category, err := m.categoryDao.CategoryGetById(dish.CategoryID)

		if err != nil {
			return response.Page[vo.DishVo]{}, response.SERVER_ERROR()
		}

		// 对象拷贝 把dish的值拷贝到dishVo
		copier.Copy(&dishVo, &dish)

		// 设置分类名称
		dishVo.CategoryName = category.Name

		// 添加到切片中
		dishVos = append(dishVos, dishVo)

	}

	dishVoPage := response.NewPage[vo.DishVo](dishPage.Total, dishVos, dishPage.Page, dishPage.PageSize)

	return dishVoPage, response.SUCCESS()
}

// DishDelete 菜品删除
func (m *DishService) DishDelete(id int64) response.ResultCode {

	// 1、调用dao层
	err := m.dishDao.DishDelete(id)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}

func (m *DishService) DishUpdateStatus(id int64, status int) response.ResultCode {

	// 1、调用dao层
	err := m.dishDao.DishUpdateStatus(id, status)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}
