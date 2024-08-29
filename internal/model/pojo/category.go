package pojo

// Category 菜品分类表
type Category struct {
	BasePojo
	Type       int    `gorm:"column:type;type:int(11)" json:"type"`                          // 类型: 1 菜品分类, 2 套餐分类
	Name       string `gorm:"column:name;type:varchar(64);not null;unique" json:"name"`      // 分类名称
	Sort       int    `gorm:"column:sort;type:int(11);default:0;not null" json:"sort"`       // 顺序
	CreateUser int64  `gorm:"column:create_user;type:bigint(20);not null" json:"createUser"` // 创建人
	UpdateUser int64  `gorm:"column:update_user;type:bigint(20);not null" json:"updateUser"` // 修改人
}
