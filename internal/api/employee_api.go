package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/service"
	"go-reggie/internal/utils/response"
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

func (m *EmployeeApi) EmployeeLogin(c *gin.Context) {

	// 1、校验请求参数
	var EmployeeDto dto.EmployeeDto

	// 验证码绑定失败 抛出参数错误异常
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

func (m *EmployeeApi) EmployeeLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	response.Ok(nil, c)
}

func (m *EmployeeApi) EmployeeSave(c *gin.Context) {
	// 1、校验请求参数
	var employeeDto dto.EmployeeDto

	// 验证码绑定失败 抛出参数错误异常
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
