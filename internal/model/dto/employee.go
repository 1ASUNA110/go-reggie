package dto

type EmployeeDto struct {
	UserName string `json:"username"`  // 用户名
	Password string `json:"password"`  // 密码
	Name     string `json:"name"`      // 用户名
	Phone    string `json:"phone"`     // 手机号
	Sex      string `json:"sex"`       // 性别
	IDNumber string `json:"id_number"` // 身份证号
}
