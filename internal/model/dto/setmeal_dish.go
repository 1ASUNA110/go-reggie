package dto

type SetmealDishDto struct {
	Copies int64   `json:"copies"`
	DishID int64   `json:"dishId,string"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}
