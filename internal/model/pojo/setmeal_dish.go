package pojo

type SetmealDish struct {
	BasePojo
	SetmealID  int64   `gorm:"size:32;not null;column:setmeal_id" json:"setmealId,string"`
	DishID     int64   `gorm:"size:32;not null;column:dish_id" json:"dishId,string"`
	Name       string  `gorm:"size:32;column:name" json:"name"`
	Price      float64 `gorm:"type:decimal(10,2);column:price" json:"price"`
	Copies     int     `gorm:"not null;column:copies" json:"copies"`
	Sort       int     `gorm:"default:0;column:sort" json:"sort"`
	CreateUser int64   `gorm:"not null;column:create_user" json:"createUser,string"`
	UpdateUser int64   `gorm:"not null;column:update_user" json:"updateUser,string"`
	IsDeleted  int     `gorm:"default:0;not null;column:is_deleted" json:"isDeleted"`
}
