package service

import (
	"go-reggie/internal/dao"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/utils"
	"go-reggie/internal/utils/response"
)

var employeeService *EmployeeService

type EmployeeService struct {
	//BaseService
	employeeDao *dao.EmployeeDao
}

func NewEmployeeService() *EmployeeService {
	if employeeService == nil {
		employeeService = &EmployeeService{
			employeeDao: dao.NewEmployeeDao(),
		}
	}
	return employeeService
}

func (m *EmployeeService) Login(employeeLoginDto dto.EmployeeLoginDto) (pojo.Employee, response.ResultCode) {
	// 1、将页面提交的密码password进行md5加密处理
	password := utils.MD5Hash(employeeLoginDto.Password)

	// 2、根据页面提交的用户名username查询数据库
	employee := pojo.Employee{}
	employee.Username = employeeLoginDto.Username
	employee.Password = password

	employee = m.employeeDao.Login(employee)

	// 3、如果没有查询到则返回登录失败结果
	if employee.ID == 0 {
		return employee, response.USER_LOGIN_ERROR()
	}

	// 4、查看员工状态，如果为已禁用状态，则返回员工已禁用结果
	if employee.Status == 0 {
		return employee, response.User_DISABLED_ERROR()
	}

	// 5、登录成功，返回成功结果
	return employee, response.SUCCESS()
}
