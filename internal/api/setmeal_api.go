package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-reggie/internal/model/vo/response"
	"go-reggie/internal/service"
	"strconv"
	"strings"
)

type SetmealApi struct {
	setmealService *service.SetmealService
}

var setmealApi *SetmealApi

func NewSetmealApi() *SetmealApi {
	if setmealApi == nil {
		setmealApi = &SetmealApi{
			setmealService: service.NewSetmealService(),
		}
	}

	return setmealApi
}

// SetmealSave 保存分类
//func (m *SetmealApi) SetmealSave(c *gin.Context) {
//	// 1、校验请求参数
//	var setmealDto dto.SetmealDto
//
//	// 绑定失败 抛出参数错误异常
//	if err := c.ShouldBind(&setmealDto); err != nil {
//		fmt.Println(err)
//		response.Fail(response.PARAM_ERROR(), c)
//		return
//	}
//
//	// 2、从session中获取当前登录用户
//	session := sessions.Default(c)
//	employeeId := session.Get("employee").(int64)
//
//	// 3、调用service层 保存分类
//	resultCode := m.setmealService.SetmealSave(setmealDto, employeeId)
//
//	// 4、返回响应
//	if resultCode.Code == response.SUCCESS().Code {
//		response.Ok(nil, c)
//		return
//	}
//
//	response.Fail(resultCode, c)
//}

// SetmealPage 套餐分页查询
func (m *SetmealApi) SetmealPage(c *gin.Context) {

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

	// 2、获取参数
	name := c.Query("name")

	// 3、调用service层 分页查询分类信息
	setmealPage, resultCode := m.setmealService.SetmealPage(page, pageSize, name)

	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(setmealPage, c)
		return
	}

	response.Fail(resultCode, c)

}

// SetmealDelete 删除套餐
func (m *SetmealApi) SetmealDelete(c *gin.Context) {
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
	resultCode := m.setmealService.SetmealDelete(ids)

	// 3、返回响应
	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(resultCode, c)

}

// SetmealUpdate 更新套餐
func (m *SetmealApi) SetmealUpdate(c *gin.Context) {
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
	resultCode := m.setmealService.SetmealUpdate(requestMap, employeeID.(int64))

	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}
	response.Fail(resultCode, c)

}

//
//// SetmealList 分类列表
//func (m *SetmealApi) SetmealList(c *gin.Context) {
//
//	// 1、获取参数
//	setmealTypeStr := c.Query("type")
//	setmealType, _ := strconv.Atoi(setmealTypeStr)
//
//	// 2、调用service层 查询分类列表
//	setmealList, resultCode := m.setmealService.SetmealList(setmealType)
//
//	if resultCode.Code == response.SUCCESS().Code {
//		response.Ok(setmealList, c)
//		return
//	}
//	response.Fail(resultCode, c)
//
//}

// SetmealUpdateStatus 更新套餐状态
func (m *SetmealApi) SetmealUpdateStatus(c *gin.Context) {
	// 1、获取状态
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、获取套餐ids
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
	resultCode := m.setmealService.SetmealUpdateStatus(ids, status)

	// 4、返回响应
	if resultCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(resultCode, c)
}
