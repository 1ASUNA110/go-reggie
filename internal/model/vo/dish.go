package pojo

type DishVo struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	CategoryID  int64   `json:"categoryId"`
	Price       float64 `json:"price,omitempty"`
	Code        string  `json:"code"`
	Image       string  `json:"image"`
	Description string  `json:"description,omitempty"`
	Status      int     `json:"status"`
	Sort        int     `json:"sort"`
	CreateUser  int64   `json:"createUser"`
	UpdateUser  int64   `json:"updateUser"`
	IsDeleted   int     `json:"isDeleted"`
	CreateTime  int64   `json:"createTime"`
	UpdateTime  int64   `json:"updateTime"`

	// 分类名
	CategoryName string `json:"categoryName,omitempty"`
}
