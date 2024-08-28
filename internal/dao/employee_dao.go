package dao

import (
	"go-reggie/internal/model/pojo"
)

var employeeDao *EmployeeDao

type EmployeeDao struct {
	BaseDao
}

func NewEmployeeDao() *EmployeeDao {
	if employeeDao == nil {
		employeeDao = &EmployeeDao{
			BaseDao: NewBaseDao(),
		}
	}

	return employeeDao

}

func (m *EmployeeDao) Login(employee pojo.Employee) pojo.Employee {
	var result pojo.Employee
	m.Orm.Where("username = ? and password = ?", employee.Username, employee.Password).First(&result)

	return result
}
