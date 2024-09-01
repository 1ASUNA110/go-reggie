package vo

import "time"

type DishVo struct {
	ID          int64     `json:"id,string"`
	Name        string    `json:"name"`
	CategoryID  int64     `json:"categoryId,string"`
	Price       float64   `json:"price,omitempty"`
	Code        string    `json:"code"`
	Image       string    `json:"image"`
	Description string    `json:"description,omitempty"`
	Status      int       `json:"status"`
	Sort        int       `json:"sort"`
	CreateUser  int64     `json:"createUser,string"`
	UpdateUser  int64     `json:"updateUser,string"`
	IsDeleted   int       `json:"isDeleted"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`

	// 分类名
	CategoryName string `json:"categoryName,omitempty"`

	// 口味
	Flavors []DishFlavorVo `json:"flavors"`
}
