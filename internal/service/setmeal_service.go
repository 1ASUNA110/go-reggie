package service

import (
	"github.com/jinzhu/copier"
	"go-reggie/internal/dao"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/model/vo"
	"go-reggie/internal/model/vo/response"
)

type SetmealService struct {
	setmealDao     *dao.SetmealDao
	dishDao        *dao.DishDao
	categoryDao    *dao.CategoryDao
	setmealDishDao *dao.SetmealDishDao
}

var setmealService *SetmealService

func NewSetmealService() *SetmealService {
	if setmealService == nil {
		setmealService = &SetmealService{
			setmealDao:     dao.NewSetmealDao(),
			dishDao:        dao.NewDishDao(),
			categoryDao:    dao.NewCategoryDao(),
			setmealDishDao: dao.NewSetmealDishDao(),
		}
	}

	return setmealService

}

func (m *SetmealService) SetmealPage(page int, pageSize int, name string) (response.Page[vo.SetmealVo], response.ResultCode) {

	// 1、调用dao层查询分类列表
	setmealPage, err := m.setmealDao.SetmealPage(page, pageSize, name)

	// 2、给套餐添加分类信息
	// 这里用最简单的做法直接遍历每个套餐，然后查询每个套餐的分类信息

	records := setmealPage.Records

	// 定义一个vo的切片
	setmealVos := []vo.SetmealVo{}

	for i := 0; i < len(records); i++ {
		setmealVo := vo.SetmealVo{}

		setmeal := records[i]
		category, err := m.categoryDao.CategoryGetById(setmeal.CategoryID)

		if err != nil {
			return response.Page[vo.SetmealVo]{}, response.SERVER_ERROR()
		}

		// 对象拷贝 把setmeal的值拷贝到setmealVo
		copier.Copy(&setmealVo, &setmeal)

		// 设置分类名称
		setmealVo.CategoryName = category.Name
		setmealVos = append(setmealVos, setmealVo)
	}

	if err != nil {
		return response.Page[vo.SetmealVo]{}, response.SERVER_ERROR()
	}

	setmealVoPage := response.NewPage[vo.SetmealVo](setmealPage.Total, setmealVos, setmealPage.Page, setmealPage.PageSize)

	return setmealVoPage, response.SUCCESS()

}

// SetmealUpdate 更新套餐
func (m *SetmealService) SetmealUpdate(setmealDto dto.SetmealDto) response.ResultCode {
	// 1、创建套餐对象
	setmeal := pojo.Setmeal{}
	copier.Copy(&setmeal, &setmealDto)

	// 2、调用dao层更新套餐
	err := m.setmealDao.SetmealUpdate(setmeal)

	if err != nil {
		return response.SERVER_ERROR()
	}

	// 3、删除套餐菜品关系
	err = m.setmealDishDao.SetmealDishDeleteBySetmealId(setmeal.ID)

	// 4、重新添加套餐菜品关系
	for _, dish := range setmealDto.SetmealDishes {
		setmealDish := pojo.SetmealDish{}
		copier.Copy(&setmealDish, &dish)

		setmealDish.SetmealID = setmeal.ID

		err = m.setmealDishDao.SetmealDishSave(setmealDish)
	}

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}

// SetmealUpdateStatus 更新套餐状态
func (m *SetmealService) SetmealUpdateStatus(ids []int64, status int) response.ResultCode {

	// 1、调用dao层

	for _, id := range ids {
		err := m.setmealDao.SetmealUpdateStatus(id, status)

		if err != nil {
			return response.SERVER_ERROR()
		}
	}

	return response.SUCCESS()

	// 2、调用dao层更新分类

}

func (m *SetmealService) SetmealDelete(ids []int64) response.ResultCode {
	// 1、调用dao层
	for _, id := range ids {

		err := m.setmealDao.SetmealDelete(id)

		if err != nil {
			return response.SERVER_ERROR()
		}
	}
	return response.SUCCESS()
}

func (m *SetmealService) SetmealSave(setmealDto dto.SetmealDto) response.ResultCode {
	// 1、获取套餐信息
	var setmeal pojo.Setmeal

	copier.Copy(&setmeal, &setmealDto)

	// 2、调用dao层保存套餐
	err := m.setmealDao.SetmealSave(&setmeal)

	if err != nil {
		return response.SERVER_ERROR()
	}

	// 2、获取套餐菜品信息
	setmealDishes := setmealDto.SetmealDishes

	// 3、调用dao层保存套餐菜品
	for i := 0; i < len(setmealDishes); i++ {
		var setmealDish pojo.SetmealDish

		copier.Copy(&setmealDish, &setmealDishes[i])

		setmealDish.SetmealID = setmeal.ID

		err := m.setmealDishDao.SetmealDishSave(setmealDish)

		if err != nil {
			return response.SERVER_ERROR()
		}

	}

	return response.SUCCESS()

}

func (m *SetmealService) SetmealGetById(id int64) (vo.SetmealVo, response.ResultCode) {

	// 1、调用dao层获取套餐信息
	setmeal, err := m.setmealDao.SetmealGetById(id)

	if err != nil {
		return vo.SetmealVo{}, response.SERVER_ERROR()
	}

	// 2、调用dao层获取套餐菜品信息
	setmealDishes, err := m.setmealDishDao.SetmealDishGetBySetmealId(id)

	if err != nil {
		return vo.SetmealVo{}, response.SERVER_ERROR()
	}

	// 3、调用dao层获取分类信息
	category, err := m.categoryDao.CategoryGetById(setmeal.CategoryID)

	// 3、构建返回值
	var setmealVo vo.SetmealVo

	copier.Copy(&setmealVo, &setmeal)

	// 4、构建setmealDishVo
	var setmealDishVos []vo.SetmealDishVo

	for i := 0; i < len(setmealDishes); i++ {
		var setmealDishVo vo.SetmealDishVo
		copier.Copy(&setmealDishVo, &setmealDishes[i])
		setmealDishVos = append(setmealDishVos, setmealDishVo)
	}
	setmealVo.SetmealDishes = setmealDishVos

	if err != nil {
		return setmealVo, response.SERVER_ERROR()
	}

	setmealVo.CategoryName = category.Name

	return setmealVo, response.SUCCESS()

}
