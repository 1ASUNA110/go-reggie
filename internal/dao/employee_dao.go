package dao

import (
	"go-reggie/internal/model/pojo"
	"go-reggie/internal/utils/response"
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

func (m *EmployeeDao) FindEmployeeById(id int) (pojo.Employee, error) {
	var employee pojo.Employee
	err := m.Orm.Where("id = ?", id).First(&employee).Error
	return employee, err
}

func (m *EmployeeDao) FindEmployeeByUsername(username string) (pojo.Employee, error) {
	var employee pojo.Employee

	err := m.Orm.Where("username = ?", username).First(&employee).Error

	return employee, err
}

func (m *EmployeeDao) EmployeePage(page int, pageSize int, name string) (response.Page, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	query := m.Orm

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// 执行分页查询
	var employees []pojo.Employee
	query.Offset(offset).Limit(pageSize).Find(&employees)

	// 获取总记录数
	var total int64
	query.Model(&pojo.Employee{}).Count(&total)

	employeePage := response.Page{
		Total:    total,
		Records:  employees,
		Page:     int(page),
		PageSize: int(pageSize),
	}

	return employeePage, nil

}

func (m *EmployeeDao) EmployeeUpdate(updateMap map[string]interface{}) error {
	// 1、获取id
	id := updateMap["id"]

	// 2、移除id
	delete(updateMap, "id")

	return m.Orm.Model(&pojo.Employee{}).Where("id = ?", id).Updates(updateMap).Error

}
