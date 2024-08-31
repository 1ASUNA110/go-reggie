package pojo

type DishFlavor struct {
	BasePojo
	DishID     int64  `gorm:"column:dish_id" json:"dishId"`
	Name       string `gorm:"column:name;size:64" json:"name"`
	Value      string `gorm:"column:value;size:500" json:"value"`
	CreateUser int64  `gorm:"column:create_user" json:"createUser"`
	UpdateUser int64  `gorm:"column:update_user" json:"updateUser"`
	IsDeleted  int    `gorm:"column:is_deleted;default:0" json:"is_deleted"`
}
