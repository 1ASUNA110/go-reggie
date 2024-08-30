package api

import (
	"github.com/gin-gonic/gin"
	"go-reggie/internal/service"
	"go-reggie/internal/utils/response"
	"strconv"
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
