package dto

type DishDto struct {
	ID          int64           `json:"id,string"`
	Name        string          `json:"name"`
	CategoryID  int64           `json:"categoryId,string"`
	Price       float64         `json:"price,string"`
	Code        string          `json:"code"`
	Image       string          `json:"image"`
	Description string          `json:"description,omitempty"`
	Status      int             `json:"status"`
	Sort        int             `json:"sort"`
	Flavors     []DishFlavorDto `json:"flavors"`
}
