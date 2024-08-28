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
	var employeeLoginDto dto.EmployeeLoginDto

	// 验证码绑定失败 抛出参数错误异常
	if err := c.ShouldBind(&employeeLoginDto); err != nil {
		response.Fail(response.PARAM_ERROR(), c)
		return
	}

	// 2、调用service层 校验账号密码是否正确
	employee, errorCode := m.employeeService.Login(employeeLoginDto)

	// 3、返回响应
	if errorCode.Code == response.SUCCESS().Code {

		session := sessions.Default(c)
		session.Set("employee", employee.ID)
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
