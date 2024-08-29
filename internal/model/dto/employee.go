package dto

type EmployeeDto struct {
	ID       int64  `json:"id,string"` // 用户ID
	UserName string `json:"username"`  // 用户名
	Password string `json:"password"`  // 密码
	Name     string `json:"name"`      // 用户名
	Phone    string `json:"phone"`     // 手机号
	Sex      string `json:"sex"`       // 性别
	IDNumber string `json:"id_number"` // 身份证号
	Status   int    `json:"status"`    // 状态 0:禁用，1:正常
}
