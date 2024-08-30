package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/vo/response"
	"go-reggie/internal/service"
	"strconv"
)

type CategoryApi struct {
	categoryService *service.CategoryService
}

var categoryApi *CategoryApi

func NewCategoryApi() *CategoryApi {
	if categoryApi == nil {
		categoryApi = &CategoryApi{
			categoryService: service.NewCategoryService(),
		}
	}

	return categoryApi
}

// CategorySave 保存分类
func (m *CategoryApi) CategorySave(c *gin.Context) {
	// 1、校验请求参数
	var categoryDto dto.CategoryDto

	// 绑定失败 抛出参数错误异常
	if err := c.ShouldBind(&categoryDto); err != nil {
		fmt.Println(err)
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、从session中获取当前登录用户
	session := sessions.Default(c)
	employeeId := session.Get("employee").(int64)

	// 3、调用service层 保存分类
	resultCode := m.categoryService.CategorySave(categoryDto, employeeId)

	// 4、返回响应
	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(resultCode, c)
}

// CategoryPage 分类分页查询
func (m *CategoryApi) CategoryPage(c *gin.Context) {

	// 2、获取page pageSize
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

	// 3、调用service层 分页查询分类信息
	categoryPage, resultCode := m.categoryService.CategoryPage(page, pageSize)

	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(categoryPage, c)
		return
	}

	response.Fail(resultCode, c)

}

// CategoryDelete 删除分类
func (m *CategoryApi) CategoryDelete(c *gin.Context) {
	// 1、校验请求参数
	idStr := c.Query("ids")
	fmt.Println(idStr)
	// 将字符串转换为整数
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)

	// 2、调用service层 删除分类
	resultCode := m.categoryService.CategoryDelete(int64(id))

	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}
	response.Fail(resultCode, c)

}

// CategoryUpdate 更新分类
func (m *CategoryApi) CategoryUpdate(c *gin.Context) {
	// 1、校验请求参数
	var requestMap map[string]interface{}

	if err := c.BindJSON(&requestMap); err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、获取当前登录用户的ID
	session := sessions.Default(c)
	employeeID := session.Get("employee")

	// 3、调用service层 更新分类信息
	resultCode := m.categoryService.CategoryUpdate(requestMap, employeeID.(int64))

	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}
	response.Fail(resultCode, c)

}
