package service

import (
	"go-reggie/internal/dao"
	"go-reggie/internal/model/dto"
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/utils"
	"go-reggie/internal/utils/response"
	"time"
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

func (m *EmployeeService) Login(EmployeeDto dto.EmployeeDto) (pojo.Employee, response.ResultCode) {
	// 1、将页面提交的密码password进行md5加密处理
	password := utils.MD5Hash(EmployeeDto.Password)

	// 2、根据页面提交的用户名username查询数据库
	employee := pojo.Employee{}
	employee.Username = EmployeeDto.UserName
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

func (m *EmployeeService) EmployeeSave(employeeDto dto.EmployeeDto, createUser int64) response.ResultCode {
	// 1、创建employee对象
	employee := pojo.Employee{
		Username:   employeeDto.UserName,
		Name:       employeeDto.Name,
		Phone:      employeeDto.Phone,
		Sex:        employeeDto.Sex,
		IDNumber:   employeeDto.IDNumber,
		Status:     1,
		CreateUser: createUser,
		UpdateUser: createUser,
	}

	// 2、密码md5加密
	// 判断有没有密码，没有的话设置默认密码123456
	if employeeDto.Password == "" {
		employeeDto.Password = "123456"
	}

	employee.Password = utils.MD5Hash(employeeDto.Password)

	// 3、设置createTime 和 updateTime
	employee.CreateTime = time.Now()
	employee.UpdateTime = time.Now()

	// 3、查找用户名是否存在
	employee1, err := m.employeeDao.FindEmployeeByUsername(employee.Username)
	if err == nil && employee1.ID != 0 {
		return response.USER_IS_EXIST()
	}

	// 3、调用dao层保存
	err = m.employeeDao.EmployeeSave(&employee)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()

}

func (m *EmployeeService) EmployeePage(page int, pageSize int, name string) (response.Page, response.ResultCode) {

	// 3、查询数据库
	employeePage, err := m.employeeDao.EmployeePage(page, pageSize, name)

	if err != nil {
		return response.Page{}, response.SERVER_ERROR()
	}

	return employeePage, response.SUCCESS()

}

func (m *EmployeeService) EmployeeUpdate(updateMap map[string]interface{}, employeeId int64) response.ResultCode {
	// 1、判断是否有id
	if _, ok := updateMap["id"]; !ok {
		return response.PARAM_ERROR()
	}

	// 2、判断是否有密码，有的话进行md5加密
	if password, ok := updateMap["password"]; ok {
		updateMap["password"] = utils.MD5Hash(password.(string))
	}

	// 4、设置updateUser
	updateMap["update_user"] = employeeId

	// 5、设置updateTime
	updateMap["update_time"] = time.Now()

	// 2、调用dao层更新 员工信息
	err := m.employeeDao.EmployeeUpdate(updateMap)

	if err != nil {
		return response.SERVER_ERROR()
	}

	return response.SUCCESS()
}
