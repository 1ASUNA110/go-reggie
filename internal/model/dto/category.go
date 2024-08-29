package dto

type CategoryDto struct {
	ID   int64  `json:"id,string"`   // 分类ID
	Type int    `json:"type,string"` // 类型: 1 菜品分类, 2 套餐分类
	Name string `json:"name"`        // 分类名称
	Sort int    `json:"sort,string"` // 顺序
}
