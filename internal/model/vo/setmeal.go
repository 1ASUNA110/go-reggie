package vo

import "time"

type SetmealVo struct {
	ID          int64     `json:"id,string"`
	CategoryID  int64     `json:"categoryId,string"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Status      *int      `json:"status"`
	Code        *string   `json:"code"`
	Description *string   `json:"description"`
	Image       *string   `json:"image"`
	CreateUser  int64     `json:"createUser,string"`
	UpdateUser  int64     `json:"updateUser,string"`
	IsDeleted   int       `json:"isDeleted"`
	UpdateTime  time.Time `json:"updateTime"`

	// 分类名
	CategoryName string `json:"categoryName,omitempty"`

	// 套餐关联的菜品列表
	SetmealDishes []SetmealDishVo `json:"setmealDishes,omitempty"`
}
