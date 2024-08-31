package dto

type DishFlavorDto struct {
	ID     int64  `json:"id"`
	DishID int64  `json:"dishId"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}
