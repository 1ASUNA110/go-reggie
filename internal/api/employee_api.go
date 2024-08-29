package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/service"
	"go-reggie/internal/utils/response"
	"strconv"
)

type EmployeeApi struct {
	employeeService *service.EmployeeService
}

var employeeApi *EmployeeApi

func NewEmployeeApi() *EmployeeApi {
	if employeeApi == nil {
		employeeApi = &EmployeeApi{
			employeeService: service.NewEmployeeService(),
		}
	}

	return employeeApi
}

// EmployeeLogin 员工登录
func (m *EmployeeApi) EmployeeLogin(c *gin.Context) {

	// 1、校验请求参数
	var EmployeeDto dto.EmployeeDto

	// 绑定失败 抛出参数错误异常
	if err := c.ShouldBind(&EmployeeDto); err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、调用service层 校验账号密码是否正确
	employee, errorCode := m.employeeService.Login(EmployeeDto)

	// 3、返回响应
	if errorCode.Code == response.SUCCESS().Code {

		session := sessions.Default(c)
		session.Set("employee", employee.ID)

		// 保存session
		session.Save()

		response.Ok(employee, c)
	} else {
		response.Fail(errorCode, c)
	}
}

// 员工退出
func (m *EmployeeApi) EmployeeLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	response.Ok(nil, c)
}

// EmployeeSave 新增员工
func (m *EmployeeApi) EmployeeSave(c *gin.Context) {
	// 1、校验请求参数
	var employeeDto dto.EmployeeDto

	// 绑定失败 抛出参数错误异常
	if err := c.ShouldBind(&employeeDto); err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、获取当前登录用户的ID
	session := sessions.Default(c)
	employeeID := session.Get("employee")

	// 3、调用service层 保存员工信息
	errorCode := m.employeeService.EmployeeSave(employeeDto, employeeID.(int64))

	if errorCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(errorCode, c)
}

// EmployeePage 员工分页查询
func (m *EmployeeApi) EmployeePage(c *gin.Context) {
	// 1、校验请求参数

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

	name := c.DefaultQuery("name", "")

	// 3、调用service层 分页查询员工信息
	employeePage, errorCode := m.employeeService.EmployeePage(page, pageSize, name)

	if errorCode.Code == response.SUCCESS().Code {
		response.Ok(employeePage, c)
		return
	}

	response.Fail(errorCode, c)
}

// EmployeeUpdate 更新员工信息
func (m *EmployeeApi) EmployeeUpdate(c *gin.Context) {
	// 1、校验请求参数
	var requestMap map[string]interface{}

	if err := c.BindJSON(&requestMap); err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、获取当前登录用户的ID
	session := sessions.Default(c)
	employeeID := session.Get("employee")

	// 3、调用service层 更新员工信息
	errorCode := m.employeeService.EmployeeUpdate(requestMap, employeeID.(int64))

	if errorCode.Code == response.SUCCESS().Code {
		response.Ok(nil, c)
		return
	}

	response.Fail(errorCode, c)

}

// EmployeeGetById 根据ID查询员工信息
func (m *EmployeeApi) EmployeeGetById(c *gin.Context) {
	// 1、校验请求参数
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil || id <= 0 {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、调用service层 根据ID查询员工信息
	employee, errorCode := m.employeeService.EmployeeGetById(int64(id))

	if errorCode.Code == response.SUCCESS().Code {
		response.Ok(employee, c)
		return
	}

	response.Fail(errorCode, c)
}
