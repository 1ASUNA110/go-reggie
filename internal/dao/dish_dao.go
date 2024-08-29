package dao

type DishDao struct {
	BaseDao
}

var dishDao *DishDao

func NewDishDao() *DishDao {
	if dishDao == nil {
		dishDao = &DishDao{
			BaseDao: NewBaseDao(),
		}
	}

	return dishDao
}

func (m DishDao) DishCountByCategoryId(categoryId int64) (int64, error) {
	var count int64

	err := m.Orm.Table("dish").Where("category_id = ?", categoryId).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
