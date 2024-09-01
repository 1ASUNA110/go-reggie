package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/vo/response"
	"go-reggie/internal/service"
	"strconv"
	"strings"
)

type DishApi struct {
	dishService *service.DishService
}

var dishApi *DishApi

func NewDishApi() *DishApi {
	if dishApi == nil {
		dishApi = &DishApi{
			dishService: service.NewDishService(),
		}
	}

	return dishApi
}

// DishPage 菜品分页查询
func (m *DishApi) DishPage(c *gin.Context) {
	// 1、获取page pageSize
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	// 将字符串转换为整数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// 2、获取菜品名
	name := c.DefaultQuery("name", "")

	// 3、调用service层 分页查询菜品
	dishPage, resultCode := m.dishService.DishPage(page, pageSize, name)

	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(dishPage, c)
		return
	}

	response.Fail(resultCode, c)

}

// DishDelete 删除菜品
func (m *DishApi) DishDelete(c *gin.Context) {
	// 1、获取菜品id
	idsStr := c.Query("ids")

	idList := strings.Split(idsStr, ",")

	ids := make([]int64, 0)

	for _, idStr := range idList {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.Fail(response.PARAM_ERROR(), c)
			return
		}
		ids = append(ids, id)
	}

	// 2、调用service层 删除菜品
	resultCode := m.dishService.DishDelete(ids)

	// 3、返回响应
	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(resultCode, c)
}

// DishUpdateStatus 更新菜品状态
func (m *DishApi) DishUpdateStatus(c *gin.Context) {
	// 1、获取状态
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、获取菜品ids
	idsStr := c.Query("ids")

	idList := strings.Split(idsStr, ",")

	ids := make([]int64, 0)

	for _, idStr := range idList {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.Fail(response.PARAM_ERROR(), c)
			return
		}
		ids = append(ids, id)
	}

	// 3、调用service层 更新菜品状态
	resultCode := m.dishService.DishUpdateStatus(ids, status)

	// 4、返回响应
	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(resultCode, c)
}

// DishSave 保存菜品
func (m *DishApi) DishSave(c *gin.Context) {
	// 1、获取菜品信息
	var dishDto dto.DishDto

	err := c.ShouldBindJSON(&dishDto)

	if err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		print(err.Error())
		return
	}

	fmt.Println(dishDto)

	// 2、调用service层 保存菜品
	resultCode := m.dishService.DishSave(dishDto)

	// 3、返回响应
	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(resultCode, c)
}

// DishGetById 根据id查询菜品
func (m *DishApi) DishGetById(c *gin.Context) {
	// 1、获取菜品id
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(response.PARAM_ERROR(), c)
	}

	// 2、调用service层 根据id查询菜品
	dishVo, resultCode := m.dishService.DishGetById(id)

	// 3、返回响应
	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(dishVo, c)
		return
	}

	response.Fail(resultCode, c)

}

// DishUpdate 更新菜品

// DishList 菜品列表
func (m *DishApi) DishList(c *gin.Context) {

	// 1、获取参数
	categoryIdStr := c.Query("categoryId")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)

	if err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、调用service层 查询菜品列表
	dishList, resultCode := m.dishService.DishList(categoryId)

	// 3、返回响应
	if resultCode.Code != response.SUCCESS().Code {
		response.Fail(resultCode, c)
		return
	}

	response.Ok(dishList, c)

}
