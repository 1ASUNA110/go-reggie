package vo

type SetmealDishVo struct {
	ID     int64   `json:"id,string"`
	Copies int64   `json:"copies"`
	DishID int64   `json:"dishId,string"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}
