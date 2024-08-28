package pojo

// Employee 结构体表示数据库中的员工信息表
type Employee struct {
	BasePojo
	Name       string `json:"name" gorm:"size:32;not null;column:name"`                // 姓名
	Username   string `json:"username" gorm:"size:32;unique;not null;column:username"` // 用户名
	Password   string `json:"password" gorm:"size:64;not null;column:password"`        // 密码
	Phone      string `json:"phone" gorm:"size:11;not null;column:phone"`              // 手机号
	Sex        string `json:"sex" gorm:"size:2;not null;column:sex"`                   // 性别
	IDNumber   string `json:"id_number" gorm:"size:18;not null;column:id_number"`      // 身份证号
	Status     int    `json:"status" gorm:"not null;default:1;column:status"`          // 状态 0:禁用，1:正常
	CreateUser int64  `json:"create_user" gorm:"not null;column:create_user"`          // 创建人
	UpdateUser int64  `json:"update_user" gorm:"not null;column:update_user"`          // 修改人
}
