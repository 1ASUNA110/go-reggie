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

func (m *EmployeeDao) EmployeeSave(employee *pojo.Employee) error {
	return m.Orm.Create(employee).Error
}

func (m *EmployeeDao) FindEmployeeById(id int64) (pojo.Employee, error) {
	var employee pojo.Employee
	err := m.Orm.Where("id = ?", id).First(&employee).Error
	return employee, err
}

func (m *EmployeeDao) FindEmployeeByUsername(username string) (pojo.Employee, error) {
	var employee pojo.Employee

	err := m.Orm.Where("username = ?", username).First(&employee).Error

	return employee, err
}
