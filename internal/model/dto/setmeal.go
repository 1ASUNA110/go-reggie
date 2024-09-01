package dto

type SetmealDto struct {
	ID          int64   `json:"id,string"`
	CategoryID  int64   `json:"categoryId,string"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Status      *int    `json:"status"`
	Code        *string `json:"code"`
	Description *string `json:"description"`
	Image       *string `json:"image"`

	// 套餐菜品
	SetmealDishes []SetmealDishDto `json:"setmealDishes"`
}
