package pojo

type Setmeal struct {
	BasePojo
	CategoryID  int64   `gorm:"not null;comment:菜品分类id" json:"categoryId,string"`
	Name        string  `gorm:"size:64;uniqueIndex:idx_setmeal_name;not null;comment:套餐名称" json:"name"`
	Price       float64 `gorm:"type:decimal(10,2);not null;comment:套餐价格" json:"price"`
	Status      *int    `gorm:"default:NULL;comment:状态 0:停用 1:启用" json:"status"`
	Code        *string `gorm:"size:32;default:NULL;comment:编码" json:"code"`
	Description *string `gorm:"size:512;default:NULL;comment:描述信息" json:"description"`
	Image       *string `gorm:"size:255;default:NULL;comment:图片" json:"image"`
	CreateUser  int64   `gorm:"not null;comment:创建人" json:"createUser"`
	UpdateUser  int64   `gorm:"not null;comment:修改人" json:"updateUser"`
	IsDeleted   int     `gorm:"not null;default:0;comment:是否删除" json:"isDeleted"`
}
