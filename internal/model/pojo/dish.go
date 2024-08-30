package pojo

type Dish struct {
	BasePojo
	Name        string  `gorm:"size:64;not null;unique;column:name" json:"name"`
	CategoryID  int64   `gorm:"column:category_id;not null" json:"categoryId"`
	Price       float64 `gorm:"type:decimal(10,2);column:price" json:"price,omitempty"`
	Code        string  `gorm:"size:64;not null;column:code" json:"code"`
	Image       string  `gorm:"size:200;not null;column:image" json:"image"`
	Description string  `gorm:"size:400;column:description" json:"description,omitempty"`
	Status      int     `gorm:"default:1;not null;column:status" json:"status"`
	Sort        int     `gorm:"default:0;not null;column:sort" json:"sort"`
	CreateUser  int64   `gorm:"column:create_user;not null" json:"createUser"`
	UpdateUser  int64   `gorm:"column:update_user;not null" json:"updateUser"`
	IsDeleted   int     `gorm:"default:0;not null;column:is_deleted" json:"isDeleted"`
}
