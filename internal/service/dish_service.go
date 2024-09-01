package service

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"go-reggie/internal/dao"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/pojo"
	vo "go-reggie/internal/model/vo"
	"go-reggie/internal/model/vo/response"
)

type DishService struct {
	dishDao       *dao.DishDao
	categoryDao   *dao.CategoryDao
	dishFlavorDao *dao.DishFlavorDao
}

var dishService *DishService

func NewDishService() *DishService {
	if dishService == nil {
		dishService = &DishService{
			dishDao:       dao.NewDishDao(),
			categoryDao:   dao.NewCategoryDao(),
			dishFlavorDao: dao.NewDishFlavorDao(),
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
func (m *DishService) DishDelete(ids []int64) response.ResultCode {

	// 1、调用dao层
	for _, id := range ids {

		err := m.dishDao.DishDelete(id)

		if err != nil {
			return response.SERVER_ERROR()
		}
	}
	return response.SUCCESS()

}

func (m *DishService) DishUpdateStatus(ids []int64, status int) response.ResultCode {

	// 1、调用dao层
	for _, id := range ids {

		err := m.dishDao.DishUpdateStatus(id, status)

		if err != nil {
			return response.SERVER_ERROR()
		}
	}

	return response.SUCCESS()

}

func (m *DishService) DishSave(dto dto.DishDto) response.ResultCode {
	// 1、获取菜品信息
	var dish pojo.Dish

	// 2、对象拷贝
	copier.Copy(&dish, &dto)

	// 3、获取菜品口味信息
	var flavors []pojo.DishFlavor

	for i := 0; i < len(dto.Flavors); i++ {
		flavor := pojo.DishFlavor{}
		copier.Copy(&flavor, &dto.Flavors[i])
		flavor.DishID = dish.ID
		flavors = append(flavors, flavor)
	}

	// 4、开启事务
	tx := m.dishDao.Orm.Begin()

	// 5、在事务中保存菜品
	err := tx.Create(&dish).Error
	if err != nil {
		tx.Rollback() // 如果出错，回滚事务

		// 判断是否是 唯一约束冲突
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return response.ERROR_DISH_NAME_UNIQUE()
		}

		return response.SERVER_ERROR()
	}

	// 6、在事务中保存菜品口味
	if len(flavors) > 0 {
		err = tx.Create(&flavors).Error
		if err != nil {
			tx.Rollback() // 如果出错，回滚事务
			return response.SERVER_ERROR()
		}
	}

	// 7、提交事务
	tx.Commit()
	return response.SUCCESS()
}

func (m *DishService) DishGetById(id int64) (vo.DishVo, response.ResultCode) {

	// 1、调用dao层查询菜品
	dish, err := m.dishDao.DishGetById(id)

	if err != nil {
		return vo.DishVo{}, response.SERVER_ERROR()
	}

	// 2、判断是否查询到菜品
	if dish.ID == 0 {
		return vo.DishVo{}, response.ERROR_DISH_NOT_FOUND()
	}

	// 3、查询菜品口味
	flavors, err := m.dishFlavorDao.DishFlavorGetByDishId(id)

	if err != nil {
		return vo.DishVo{}, response.SERVER_ERROR()
	}

	// 4、查询菜品分类
	category, err := m.categoryDao.CategoryGetById(dish.CategoryID)

	if err != nil {
		return vo.DishVo{}, response.SERVER_ERROR()
	}

	// 5、构建返回值
	dishVo := vo.DishVo{}

	// 对象拷贝
	copier.Copy(&dishVo, &dish)

	// 设置分类名称
	dishVo.CategoryName = category.Name

	// 设置口味
	dishVo.Flavors = []vo.DishFlavorVo{}

	for i := 0; i < len(flavors); i++ {
		flavorVo := vo.DishFlavorVo{}
		copier.Copy(&flavorVo, &flavors[i])
		dishVo.Flavors = append(dishVo.Flavors, flavorVo)
	}

	return dishVo, response.SUCCESS()

}

func (m *DishService) DishList(categoryId int64) ([]vo.DishVo, response.ResultCode) {

	// 1、调用dao层查询菜品
	dishes, err := m.dishDao.DishList(categoryId)

	if err != nil {
		return []vo.DishVo{}, response.SERVER_ERROR()
	}

	// 2、判断是否查询到菜品
	if len(dishes) == 0 {
		return []vo.DishVo{}, response.SUCCESS()
	}

	dishVoList := []vo.DishVo{}
	// 3、构建返回值
	for i := 0; i < len(dishes); i++ {

		// 5、构建返回值
		dishVo := vo.DishVo{}

		// 对象拷贝
		copier.Copy(&dishVo, &dishes[i])

		dishVoList = append(dishVoList, dishVo)
	}

	return dishVoList, response.SUCCESS()
}
